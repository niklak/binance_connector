package binance_connector

import (
	"time"

	"github.com/niklak/binance_connector/internal/client"
)

type Client struct {
	*client.Connector
}

func NewClient(apiKey, secretKey, baseURL string, timeout ...time.Duration) *Client {

	client := client.NewClient(apiKey, secretKey, baseURL, timeout...)

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
