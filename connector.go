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

// Market Endpoints:

func (c *Client) NewPingService() *Ping {
	return &Ping{C: c.Connector}
}

func (c *Client) NewTimeService() *ServerTime {
	return &ServerTime{C: c.Connector}
}

func (c *Client) NewExchangeInfoService() *ExchangeInfo {
	return &ExchangeInfo{C: c.Connector}
}

func (c *Client) NewOrderBookService() *OrderBook {
	return &OrderBook{C: c.Connector}
}

func (c *Client) NewRecentTradesListService() *RecentTradesList {
	return &RecentTradesList{C: c.Connector}
}

func (c *Client) NewHistoricalTradeLookupService() *HistoricalTradeLookup {
	return &HistoricalTradeLookup{C: c.Connector}
}

func (c *Client) NewAggTradesListService() *AggTradesList {
	return &AggTradesList{C: c.Connector}
}

func (c *Client) NewKlinesService() *Klines {
	return &Klines{C: c.Connector}
}

func (c *Client) NewUIKlinesService() *UiKlines {
	return &UiKlines{C: c.Connector}
}

func (c *Client) NewAvgPriceService() *AvgPrice {
	return &AvgPrice{C: c.Connector}
}

func (c *Client) NewTicker24hrService() *Ticker24hr {
	return &Ticker24hr{C: c.Connector}
}

func (c *Client) NewTickerPriceService() *TickerPrice {
	return &TickerPrice{C: c.Connector}
}

func (c *Client) NewTickerBookTickerService() *TickerBookTicker {
	return &TickerBookTicker{C: c.Connector}
}

func (c *Client) NewTickerService() *Ticker {
	return &Ticker{C: c.Connector}
}

// Spot Endpoints:

func (c *Client) NewTestNewOrderService() *TestNewOrder {
	return &TestNewOrder{C: c.Connector}
}

func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{C: c.Connector}
}

func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{C: c.Connector}
}

func (c *Client) NewGetOpenOrdersService() *GetOpenOrdersService {
	return &GetOpenOrdersService{C: c.Connector}
}

func (c *Client) NewGetAllOrdersService() *GetAllOrdersService {
	return &GetAllOrdersService{C: c.Connector}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{C: c.Connector}
}

func (c *Client) NewCancelOpenOrdersService() *CancelOpenOrdersService {
	return &CancelOpenOrdersService{C: c.Connector}
}
