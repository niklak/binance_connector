package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Cancel an Existing Order and Send a New Order (TRADE)
//
//gen:new_service
type CancelReplaceService struct {
	C                       *connector.Connector
	symbol                  string
	side                    string
	orderType               string
	cancelReplaceMode       string
	timeInForce             *string
	quantity                *float64
	quoteOrderQty           *float64
	price                   *float64
	cancelNewClientOrderId  *string
	cancelOrigClientOrderId *string
	cancelOrderId           *int64
	newClientOrderId        *string
	strategyId              *int32
	strategyType            *int32
	stopPrice               *float64
	trailingDelta           *int64
	icebergQty              *float64
	newOrderRespType        *string
	selfTradePreventionMode *string
	cancelRestrictions      *string
}

// Symbol set symbol
func (s *CancelReplaceService) Symbol(symbol string) *CancelReplaceService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CancelReplaceService) Side(side string) *CancelReplaceService {
	s.side = side
	return s
}

// OrderType set orderType
func (s *CancelReplaceService) OrderType(orderType string) *CancelReplaceService {
	s.orderType = orderType
	return s
}

// CancelReplaceMode set cancelReplaceMode
func (s *CancelReplaceService) CancelReplaceMode(cancelReplaceMode string) *CancelReplaceService {
	s.cancelReplaceMode = cancelReplaceMode
	return s
}

// TimeInForce set timeInForce
func (s *CancelReplaceService) TimeInForce(timeInForce string) *CancelReplaceService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CancelReplaceService) Quantity(quantity float64) *CancelReplaceService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CancelReplaceService) QuoteOrderQty(quoteOrderQty float64) *CancelReplaceService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CancelReplaceService) Price(price float64) *CancelReplaceService {
	s.price = &price
	return s
}

// CancelNewClientOrderId set cancelNewClientOrderId
func (s *CancelReplaceService) CancelNewClientOrderId(cancelNewClientOrderId string) *CancelReplaceService {
	s.cancelNewClientOrderId = &cancelNewClientOrderId
	return s
}

// CancelOrigClientOrderId set cancelOrigClientOrderId
func (s *CancelReplaceService) CancelOrigClientOrderId(cancelOrigClientOrderId string) *CancelReplaceService {
	s.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return s
}

// CancelOrderId set cancelOrderId
func (s *CancelReplaceService) CancelOrderId(cancelOrderId int64) *CancelReplaceService {
	s.cancelOrderId = &cancelOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelReplaceService) NewClientOrderId(newClientOrderId string) *CancelReplaceService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// StrategyId set strategyId
func (s *CancelReplaceService) StrategyId(strategyId int32) *CancelReplaceService {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *CancelReplaceService) StrategyType(strategyType int32) *CancelReplaceService {
	s.strategyType = &strategyType
	return s
}

// StopPrice set stopPrice
func (s *CancelReplaceService) StopPrice(stopPrice float64) *CancelReplaceService {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *CancelReplaceService) TrailingDelta(trailingDelta int64) *CancelReplaceService {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty set icebergQty
func (s *CancelReplaceService) IcebergQty(icebergQty float64) *CancelReplaceService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *CancelReplaceService) NewOrderRespType(newOrderRespType string) *CancelReplaceService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CancelReplaceService) SelfTradePreventionMode(selfTradePreventionMode string) *CancelReplaceService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// CancelRestrictions set cancelRestrictions
func (s *CancelReplaceService) CancelRestrictions(cancelRestrictions string) *CancelReplaceService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

// Do send request
func (s *CancelReplaceService) Do(ctx context.Context, opts ...request.RequestOption) (res *CancelReplaceResponse, err error) {

	r := request.New(
		"/api/v3/order/cancelReplace",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol", "side", "type", "cancelReplaceMode"),
		request.RequiredOneOfParams([]string{"cancelOrderId", "cancelOrigClientOrderId"}),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("side", s.side)
	r.SetParam("type", s.orderType)
	r.SetParam("cancelReplaceMode", s.cancelReplaceMode)

	r.SetParam("cancelOrderId", s.cancelOrderId)
	r.SetParam("cancelOrigClientOrderId", s.cancelOrigClientOrderId)

	r.SetParam("timeInForce", s.timeInForce)
	r.SetParam("quantity", s.quantity)
	r.SetParam("quoteOrderQty", s.quoteOrderQty)
	r.SetParam("price", s.price)
	r.SetParam("cancelNewClientOrderId", s.cancelNewClientOrderId)
	r.SetParam("newClientOrderId", s.newClientOrderId)
	r.SetParam("strategyId", s.strategyId)
	r.SetParam("strategyType", s.strategyType)
	r.SetParam("stopPrice", s.stopPrice)
	r.SetParam("trailingDelta", s.trailingDelta)
	r.SetParam("icebergQty", s.icebergQty)
	r.SetParam("newOrderRespType", s.newOrderRespType)
	r.SetParam("selfTradePreventionMode", s.selfTradePreventionMode)
	r.SetParam("cancelRestrictions", s.cancelRestrictions)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(CancelReplaceResponse)
	err = json.Unmarshal(data, res)
	return
}

type OrderData struct {
	Code                    int      `json:"code,omitempty"`
	Msg                     string   `json:"msg,omitempty"`
	Symbol                  string   `json:"symbol,omitempty"`
	OrigClientOrderId       string   `json:"origClientOrderId,omitempty"`
	OrderId                 int64    `json:"orderId,omitempty"`
	OrderListId             int64    `json:"orderListId,omitempty"`
	ClientOrderId           string   `json:"clientOrderId,omitempty"`
	TransactTime            uint64   `json:"transactTime,omitempty"`
	Price                   string   `json:"price,omitempty"`
	OrigQty                 string   `json:"origQty,omitempty"`
	ExecutedQty             string   `json:"executedQty,omitempty"`
	CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
	Status                  string   `json:"status,omitempty"`
	TimeInForce             string   `json:"timeInForce,omitempty"`
	Type                    string   `json:"type,omitempty"`
	Side                    string   `json:"side,omitempty"`
	Fills                   []string `json:"fills,omitempty"`
	SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
}

type CancelReplaceResponse struct {
	Code             int64      `json:"code,omitempty"`
	Msg              string     `json:"msg,omitempty"`
	CancelResult     string     `json:"cancelResult,omitempty"`
	NewOrderResult   string     `json:"newOrderResult,omitempty"`
	CancelResponse   *OrderData `json:"cancelResponse,omitempty"`
	NewOrderResponse *OrderData `json:"newOrderResponse,omitempty"`
	Data             *struct {
		CancelResult     string     `json:"cancelResult,omitempty"`
		NewOrderResult   string     `json:"newOrderResult,omitempty"`
		CancelResponse   *OrderData `json:"cancelResponse,omitempty"`
		NewOrderResponse OrderData  `json:"newOrderResponse"`
	} `json:"data,omitempty"`
}
