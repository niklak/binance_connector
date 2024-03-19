package binance_connector

import (
	"net/http"
	"time"

	"github.com/niklak/binance_connector/internal/connector"
)

type Client struct {
	*connector.Connector
}

func NewClient(apiKey, secretKey string, baseURL ...string) *Client {

	var bURL string
	if len(baseURL) > 0 {
		bURL = baseURL[0]
	}
	client := connector.NewConnector(apiKey, secretKey, bURL)

	return &Client{Connector: client.Init()}
}

func NewClientWithTimeout(apiKey, secretKey, baseURL string, timeout time.Duration) *Client {

	client := connector.NewConnector(apiKey, secretKey, baseURL, timeout)

	return &Client{Connector: client.Init()}
}

func NewWithHttpClient(apiKey, secretKey, baseURL string, httpClient *http.Client) *Client {

	client := connector.NewConnectorWithClient(apiKey, secretKey, baseURL, httpClient)

	return &Client{Connector: client.Init()}
}

func (c *Client) NewTickerPriceService() *TickerPrice {
	return &TickerPrice{C: c.Connector}
}

func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{C: c.Connector}
}

func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{C: c.Connector}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{C: c.Connector}
}

func (c *Client) NewCancelOpenOrdersService() *CancelOpenOrdersService {
	return &CancelOpenOrdersService{C: c.Connector}
}
