package binance_connector

import (
	"github.com/niklak/binance_connector/api/market"
	"github.com/niklak/binance_connector/api/spot"
)

// market
type Ping = market.Ping

type ServerTime = market.ServerTime
type ServerTimeResponse = market.ServerTimeResponse

type ExchangeInfo = market.ExchangeInfo
type ExchangeInfoResponse = market.ExchangeInfoResponse

type OrderBook = market.OrderBook
type OrderBookResponse = market.OrderBookResponse

type RecentTradesList = market.RecentTradesList
type RecentTradesListResponse = market.RecentTradesListResponse

type HistoricalTradeLookup = market.HistoricalTradeLookup
type HistoricalTradeLookupResponse = market.RecentTradesListResponse

type AggTradesList = market.AggTradesList
type AggTradesListResponse = market.AggTradesListResponse

type Klines = market.Klines
type KlinesResponse = market.KlinesResponse

type UiKlines = market.UiKlines
type UiKlinesResponse = market.UiKlinesResponse

type Ticker24hr = market.Ticker24hr
type Ticker24hrResponse = market.Ticker24hrResponse

type TickerPrice = market.TickerPrice
type TickerPriceResponse = market.TickerPriceResponse

type TickerBookTicker = market.TickerBookTicker
type TickerBookTickerResponse = market.TickerBookTickerResponse

type Ticker = market.Ticker
type TickerResponse = market.TickerResponse

// spot
// createorder.go
type CreateOrderService = spot.CreateOrderService
type CreateOrderResponseACK = spot.CreateOrderResponseACK
type CreateOrderResponseRESULT = spot.CreateOrderResponseRESULT
type CreateOrderResponseFULL = spot.CreateOrderResponseFULL

// getorder.go
type GetOrderService = spot.GetOrderService
type GetOrderResponse = spot.GetOrderResponse

// cancelorder.go
type CancelOrderService = spot.CancelOrderService
type CancelOrderResponse = spot.CancelOrderResponse

// cancelopenorders.go
type CancelOpenOrdersService = spot.CancelOpenOrdersService
