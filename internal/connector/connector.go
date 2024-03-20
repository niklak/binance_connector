package connector

import (
	"bytes"
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
	"github.com/niklak/binance_connector/internal/request"
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
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}

	if err = r.Validate(); err != nil {
		return
	}

	u, err := c.apiBaseURL.Parse(r.Endpoint)
	if err != nil {
		return
	}

	fullURL := u.String()
	if r.RecvWindow > 0 {
		r.SetParam(recvWindowKey, r.RecvWindow)
	}
	if r.SecType == request.SecTypeSigned {
		r.SetParam(timestampKey, helpers.CurrentTimestamp()-c.TimeOffset)
	}
	queryString := r.Query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.Form.Encode()
	header := http.Header{}
	if r.Header != nil {
		header = r.Header.Clone()
	}
	//header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.SecType == request.SecTypeAPIKey || r.SecType == request.SecTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.SecType == request.SecTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.logger.Debug().Str("full_url", fullURL).Str("body", bodyString).Msg("")
	r.FullURL = fullURL
	r.Header = header
	r.Body = body
	return nil
}

func (c *Connector) CallAPI(ctx context.Context, r *request.Request, opts ...request.RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if r.Endpoint != "/api/v3/order/cancelReplace" {
		if err != nil {
			return
		}
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
		if r.Endpoint != "/api/v3/order/cancelReplace" {
			return nil, apiErr
		}
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
