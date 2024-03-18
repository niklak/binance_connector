package binance_connector

import (
	"time"

	"github.com/niklak/binance_connector/api/market"
	"github.com/niklak/binance_connector/api/spot"
	"github.com/niklak/binance_connector/internal/client"
)

type Client struct {
	*client.Connector
}

func NewClient(apiKey, secretKey, baseURL string, timeout ...time.Duration) *Client {

	client := client.NewClient(apiKey, secretKey, baseURL, timeout...)

	return &Client{Connector: client.Init()}
}

func (c *Client) NewTickerPriceService() *market.TickerPrice {
	return &market.TickerPrice{C: c.Connector}
}

func (c *Client) NewCreateOrderService() *spot.CreateOrderService {
	return &spot.CreateOrderService{C: c.Connector}
}

func (c *Client) NewGetOrderService() *spot.GetOrderService {
	return &spot.GetOrderService{C: c.Connector}
}

func (c *Client) NewCancelOrderService() *spot.CancelOrderService {
	return &spot.CancelOrderService{C: c.Connector}
}

func (c *Client) NewCancelOpenOrdersService() *spot.CancelOpenOrdersService {
	return &spot.CancelOpenOrdersService{C: c.Connector}
}
