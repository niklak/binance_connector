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

// Binance New Order endpoint (POST /api/v3/order)
// CreateOrderService create order
//
//gen:new_service
type CreateOrderService struct {
	C                       *connector.Connector
	symbol                  string
	side                    string
	orderType               string
	timeInForce             *string
	quantity                *float64
	quoteOrderQty           *float64
	price                   *float64
	newClientOrderId        *string
	strategyId              *int
	strategyType            *int
	stopPrice               *float64
	trailingDelta           *int
	icebergQty              *float64
	newOrderRespType        *string
	selfTradePreventionMode *string
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side string) *CreateOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType string) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce string) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity float64) *CreateOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CreateOrderService) QuoteOrderQty(quoteOrderQty float64) *CreateOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CreateOrderService) Price(price float64) *CreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CreateOrderService) NewClientOrderId(newClientOrderId string) *CreateOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *CreateOrderService) StrategyId(strategyId int) *CreateOrderService {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *CreateOrderService) StrategyType(strategyType int) *CreateOrderService {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice float64) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *CreateOrderService) TrailingDelta(trailingDelta int) *CreateOrderService {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateOrderService) IcebergQuantity(icebergQty float64) *CreateOrderService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *CreateOrderService) NewOrderRespType(newOrderRespType string) *CreateOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *CreateOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: "/api/v3/order",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	// no need to send request if there are no required parameters
	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	if s.side == "" {
		err = fmt.Errorf("%w: side", apierrors.ErrMissingParameter)
		return
	}
	if s.orderType == "" {
		err = fmt.Errorf("%w: orderType", apierrors.ErrMissingParameter)
		return
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
	r.SetParam("selfTradePreventionMode", s.selfTradePreventionMode)
	r.SetParam("newOrderRespType", s.newOrderRespType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	switch *s.newOrderRespType {
	case "ACK":
		res = new(CreateOrderResponseACK)
	case "RESULT":
		res = new(CreateOrderResponseRESULT)
	case "FULL":
		res = new(CreateOrderResponseFULL)
	default:
		switch s.orderType {
		case "MARKET", "LIMIT":
			res = new(CreateOrderResponseFULL)
		default:
			res = new(CreateOrderResponseACK)
		}
	}

	err = json.Unmarshal(data, res)
	return
}

// Create CreateOrderResponseACK
type CreateOrderResponseACK struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	OrderListId   int64  `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  uint64 `json:"transactTime"`
}

// Create CreateOrderResponseRESULT
type CreateOrderResponseRESULT struct {
	CreateOrderResponseACK  `json:",inline"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StopPrice               string `json:"stopPrice,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

// Create CreateOrderResponseFULL
type CreateOrderResponseFULL struct {
	CreateOrderResponseRESULT `json:",inline"`
	Fills                     []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int64  `json:"tradeId"`
	} `json:"fills"`
}
