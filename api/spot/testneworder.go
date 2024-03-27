package spot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

//gen:new_service
type TestNewOrder struct {
	C                   *connector.Connector
	symbol              string
	side                string
	orderType           string
	timeInForce         *string
	quantity            *float64
	quoteOrderQty       *float64
	price               *float64
	newClientOrderId    *string
	strategyId          *int
	strategyType        *int
	stopPrice           *float64
	trailingDelta       *int
	icebergQty          *float64
	newOrderRespType    *string
	selfTradePrevention *string
}

// Symbol set symbol
func (s *TestNewOrder) Symbol(symbol string) *TestNewOrder {
	s.symbol = symbol
	return s
}

// Side set side
func (s *TestNewOrder) Side(side string) *TestNewOrder {
	s.side = side
	return s
}

// OrderType set orderType
func (s *TestNewOrder) OrderType(orderType string) *TestNewOrder {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *TestNewOrder) TimeInForce(timeInForce string) *TestNewOrder {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *TestNewOrder) Quantity(quantity float64) *TestNewOrder {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *TestNewOrder) QuoteOrderQty(quoteOrderQty float64) *TestNewOrder {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *TestNewOrder) Price(price float64) *TestNewOrder {
	s.price = &price
	return s
}

// NewClientOrderId set newClientOrderId
func (s *TestNewOrder) NewClientOrderId(newClientOrderId string) *TestNewOrder {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *TestNewOrder) StrategyId(strategyId int) *TestNewOrder {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *TestNewOrder) StrategyType(strategyType int) *TestNewOrder {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *TestNewOrder) StopPrice(stopPrice float64) *TestNewOrder {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *TestNewOrder) TrailingDelta(trailingDelta int) *TestNewOrder {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty set icebergQty
func (s *TestNewOrder) IcebergQty(icebergQty float64) *TestNewOrder {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *TestNewOrder) NewOrderRespType(newOrderRespType string) *TestNewOrder {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePrevention set selfTradePrevention
func (s *TestNewOrder) SelfTradePrevention(selfTradePrevention string) *TestNewOrder {
	s.selfTradePrevention = &selfTradePrevention
	return s
}

// Send the request
func (s *TestNewOrder) Do(ctx context.Context, opts ...request.RequestOption) (res *AccountOrderBookResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v3/order/test",
		SecType:  request.SecTypeNone,
	}
	r.Init()

	if s.symbol == "" {
		return nil, fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
	}
	if s.side == "" {
		return nil, fmt.Errorf("%w: side", apierrors.ErrMissingParameter)
	}
	if s.orderType == "" {
		return nil, fmt.Errorf("%w: orderType", apierrors.ErrMissingParameter)
	}
	r.SetParam("symbol", s.symbol)
	r.SetParam("side", s.side)
	r.SetParam("type", s.orderType)
	r.SetParam("timeInForce", s.timeInForce)
	r.SetParam("quantity", s.quantity)
	r.SetParam("quoteOrderQty", s.quoteOrderQty)
	r.SetParam("price", s.price)
	r.SetParam("newClientOrderId", s.newClientOrderId)
	r.SetParam("strategyId", s.strategyId)
	r.SetParam("strategyType", s.strategyType)
	r.SetParam("stopPrice", s.stopPrice)
	r.SetParam("trailingDelta", s.trailingDelta)
	r.SetParam("icebergQty", s.icebergQty)
	r.SetParam("newOrderRespType", s.newOrderRespType)
	r.SetParam("selfTradePreventionMode", s.selfTradePrevention)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AccountOrderBookResponse)
	err = json.Unmarshal(data, res)
	return
}

// Create AccountOrderBookResponse
type AccountOrderBookResponse struct {
}
