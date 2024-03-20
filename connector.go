package binance_connector

import (
	"github.com/niklak/binance_connector/internal/connector"
)

type ConnectorOption = connector.ConnectorOption

var BaseURL = connector.BaseURL
var HTTPClient = connector.HTTPClient
var Timeout = connector.Timeout
var TimeOffset = connector.TimeOffset

type Client struct {
	*connector.Connector
}

func NewClient(apiKey, secretKey string, baseURL ...string) *Client {

	var c *connector.Connector
	if len(baseURL) > 0 {
		c = connector.NewConnector(apiKey, secretKey, BaseURL(baseURL[0]))
	} else {
		c = connector.NewConnector(apiKey, secretKey)
	}

	return &Client{Connector: c}
}

func NewClientWithOptions(apiKey, secretKey string, options ...ConnectorOption) *Client {

	c := connector.NewConnector(apiKey, secretKey, options...)

	return &Client{Connector: c}
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
func (c *Client) NewCancelReplaceService() *CancelReplaceService {
	return &CancelReplaceService{C: c.Connector}
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

func (c *Client) NewNewOCOService() *NewOCOService {
	return &NewOCOService{C: c.Connector}
}

func (c *Client) NewCancelOCOService() *CancelOCOService {
	return &CancelOCOService{C: c.Connector}
}

func (c *Client) NewQueryOCOService() *QueryOCOService {
	return &QueryOCOService{C: c.Connector}
}

func (c *Client) NewQueryAllOCOService() *QueryAllOCOService {
	return &QueryAllOCOService{C: c.Connector}
}

func (c *Client) NewAccountService() *AccountService {
	return &AccountService{C: c.Connector}
}

func (c *Client) NewAccountTradeListService() *AccountTradeListService {
	return &AccountTradeListService{C: c.Connector}
}

func (c *Client) NewQueryCurrentOrderCountUsageService() *QueryCurrentOrderCountUsageService {
	return &QueryCurrentOrderCountUsageService{C: c.Connector}
}

func (c *Client) NewQueryPreventedMatchesService() *QueryPreventedMatchesService {
	return &QueryPreventedMatchesService{C: c.Connector}
}
