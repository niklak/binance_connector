package connector

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/niklak/binance_connector/internal/helpers"
	"github.com/niklak/binance_connector/request"
)

type Connector struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Timeout    time.Duration
	TimeOffset int64
	logger     zerolog.Logger
	apiBaseURL *url.URL
}

func (c *Connector) Init() *Connector {

	if c.BaseURL == "" {
		c.BaseURL = defaultApiURL
	}
	if c.Timeout == 0 {
		c.Timeout = defaultTimeout
	}

	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: c.Timeout}
	}

	if c.HTTPClient.Timeout == 0 {
		c.HTTPClient.Timeout = defaultTimeout
	}
	c.logger = log.With().Str("module", "binance_connector").Logger()

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		//panics!
		c.logger.Fatal().Err(err).Msg("")
	}
	c.apiBaseURL = u
	return c
}

func (c *Connector) parseRequest(r *request.Request, opts ...request.RequestOption) (err error) {

	// According to https://binance-docs.github.io/apidocs/spot/en/#general-api-information,
	// all methods accepts query parameters and have priority over form-data.

	// set request options from user
	for _, opt := range opts {
		opt(r)
	}

	u, err := c.apiBaseURL.Parse(r.Endpoint)
	if err != nil {
		return
	}

	if r.RecvWindow > 0 {
		r.SetParam(recvWindowKey, r.RecvWindow)
	}
	if r.SecType == request.SecTypeSigned {
		r.SetParam(timestampKey, helpers.CurrentTimestamp()-c.TimeOffset)
	}

	//header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	if r.SecType == request.SecTypeAPIKey || r.SecType == request.SecTypeSigned {
		r.Header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.SecType == request.SecTypeSigned {
		mac := hmac.New(sha256.New, []byte(c.SecretKey))

		if _, err = mac.Write([]byte(r.Query.Encode())); err != nil {
			return
		}

		r.Query.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
	}

	u.RawQuery = r.Query.Encode()
	r.FullURL = u.String()

	c.logger.Debug().Str("full_url", r.FullURL).Msg("")
	return
}

func (c *Connector) CallAPI(ctx context.Context, r *request.Request, opts ...request.RequestOption) (data []byte, err error) {

	if r == nil {
		return nil, ErrRequestCantBeNil
	}

	if err = r.Validate(); err != nil {
		return
	}

	if err = c.parseRequest(r, opts...); err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, r.FullURL, r.Body)
	if err != nil {
		return
	}
	req.Header = r.Header

	c.logger.Debug().Str("request", fmt.Sprintf("%#v", req)).Msg("")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}

	if data, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	defer func() {
		closeErr := resp.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	c.logger.Debug().
		Str("response", fmt.Sprintf("%#v", resp)).
		Str("response_body", string(data)).
		Int("response_status_code", resp.StatusCode).Msg("")

	if resp.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		err = json.Unmarshal(data, apiErr)
		if err != nil {
			c.logger.Error().Err(err).Msg("")
		}
		return nil, apiErr
	}
	return data, nil
}

type ConnectorOption func(*Connector)

func Timeout(timeout time.Duration) ConnectorOption {
	return func(c *Connector) {
		c.Timeout = timeout
	}
}

func TimeOffset(offset int64) ConnectorOption {
	return func(c *Connector) {
		c.TimeOffset = offset
	}
}

func HTTPClient(client *http.Client) ConnectorOption {
	return func(c *Connector) {
		c.HTTPClient = client
	}
}

func BaseURL(baseURL string) ConnectorOption {
	return func(c *Connector) {
		c.BaseURL = baseURL
	}
}

func NewConnector(apiKey, secretKey string, options ...ConnectorOption) *Connector {

	client := &Connector{
		APIKey:    apiKey,
		SecretKey: secretKey,
	}

	for _, option := range options {
		option(client)
	}

	return client.Init()
}
