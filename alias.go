package binance_connector

import (
	"github.com/niklak/binance_connector/api/market"
	"github.com/niklak/binance_connector/api/spot"
)

// market
type TickerPrice = market.TickerPrice
type TickerPriceResponse = market.TickerPriceResponse

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
