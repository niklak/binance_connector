package binance_connector

import (
	"github.com/niklak/binance_connector/api/account"
	"github.com/niklak/binance_connector/api/market"
	"github.com/niklak/binance_connector/api/spot"
	"github.com/niklak/binance_connector/api/wallet"
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

type AvgPrice = market.AvgPrice
type AvgPriceResponse = market.AvgPriceResponse

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
type TestNewOrder = spot.TestNewOrder
type AccountOrderBookResponse = spot.AccountOrderBookResponse

type CancelReplaceService = spot.CancelReplaceService
type CancelReplaceResponse = spot.CancelReplaceResponse

type GetOpenOrdersService = spot.GetOpenOrdersService
type OrderResponse = spot.OrderResponse

type GetAllOrdersService = spot.GetAllOrdersService

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

// newoco.go
type NewOCOService = spot.NewOCOService
type OrderOCOResponse = spot.OrderOCOResponse

// canceloco.go
type CancelOCOService = spot.CancelOCOService

// queryoco.go
type QueryOCOService = spot.QueryOCOService
type OCOResponse = spot.OCOResponse

// queryalloco.go
type QueryAllOCOService = spot.QueryAllOCOService

// queryopenorders.go
type QueryOpenOCOService = spot.QueryOpenOCOService

// account
type AccountService = account.AccountService
type AccountResponse = account.AccountResponse

type AccountTradeListService = account.AccountTradeListService
type AccountTradeListResponse = account.AccountTradeListResponse

type QueryCurrentOrderCountUsageService = account.QueryCurrentOrderCountUsageService
type QueryCurrentOrderCountUsageResponse = account.QueryCurrentOrderCountUsageResponse

type QueryPreventedMatchesService = account.QueryPreventedMatchesService
type QueryPreventedMatchesResponse = account.QueryPreventedMatchesResponse

// wallet

type SystemStatusService = wallet.SystemStatusService
type SystemStatusResponse = wallet.SystemStatusResponse

type AllCoinsInfoService = wallet.AllCoinsInfoService
type CoinInfo = wallet.CoinInfo

type AccountSnapshotService = wallet.AccountSnapshotService
type AccountSnapshotResponse = wallet.AccountSnapshotResponse

type DisableFastWithdrawSwitchService = wallet.DisableFastWithdrawSwitchService
type DisableFastWithdrawSwitchResponse = wallet.DisableFastWithdrawSwitchResponse

type EnableFastWithdrawSwitchService = wallet.EnableFastWithdrawSwitchService
type EnableFastWithdrawSwitchResponse = wallet.EnableFastWithdrawSwitchResponse

type WithdrawService = wallet.WithdrawService
type WithdrawResponse = wallet.WithdrawResponse

type WithdrawHistoryService = wallet.WithdrawHistoryService
type WithdrawHistoryResponse = wallet.WithdrawHistoryResponse

type DepositHistoryService = wallet.DepositHistoryService
type DepositHistoryResponse = wallet.DepositHistoryResponse

type DepositAddressService = wallet.DepositAddressService
type DepositAddressResponse = wallet.DepositAddressResponse

type AccountStatusService = wallet.AccountStatusService
type AccountStatusResponse = wallet.AccountStatusResponse

type AccountApiTradingStatusService = wallet.AccountApiTradingStatusService
type AccountApiTradingStatusResponse = wallet.AccountApiTradingStatusResponse

type DustLogService = wallet.DustLogService
type DustLogResponse = wallet.DustLogResponse

type AssetDetailService = wallet.AssetDetailService
type AssetDetailResponse = wallet.AssetDetailResponse
