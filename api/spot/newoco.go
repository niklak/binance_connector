package spot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance New OCO (TRADE) (POST /api/v3/order/oco)
// NewOCOService create new OCO order
type NewOCOService struct {
	C                       *connector.Connector
	symbol                  string
	side                    string
	quantity                float64
	price                   float64
	stopPrice               float64
	listClientOrderId       *string
	limitClientOrderId      *string
	limitStrategyId         *int
	limitStrategyType       *int
	limitIcebergQty         *float64
	trailingDelta           *int
	stopClientOrderId       *string
	stopStrategyId          *int
	stopStrategyType        *int
	stopLimitPrice          *float64
	stopIcebergQty          *float64
	stopLimitTimeInForce    *string
	newOrderRespType        *string
	selfTradePreventionMode *string
}

// Symbol set symbol
func (s *NewOCOService) Symbol(symbol string) *NewOCOService {
	s.symbol = symbol
	return s
}

// ListClientOrderId set listClientOrderId
func (s *NewOCOService) ListClientOrderId(listClientOrderId string) *NewOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// Side set side
func (s *NewOCOService) Side(side string) *NewOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *NewOCOService) Quantity(quantity float64) *NewOCOService {
	s.quantity = quantity
	return s
}

// LimitClientOrderId set limitClientOrderId
func (s *NewOCOService) LimitClientOrderId(limitClientOrderId string) *NewOCOService {
	s.limitClientOrderId = &limitClientOrderId
	return s
}

// LimitStrategyId set limitStrategyId
func (s *NewOCOService) LimitStrategyId(limitStrategyId int) *NewOCOService {
	s.limitStrategyId = &limitStrategyId
	return s
}

// LimitStrategyType set limitStrategyType
func (s *NewOCOService) LimitStrategyType(limitStrategyType int) *NewOCOService {
	s.limitStrategyType = &limitStrategyType
	return s
}

// Price set price
func (s *NewOCOService) Price(price float64) *NewOCOService {
	s.price = price
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *NewOCOService) LimitIcebergQty(limitIcebergQty float64) *NewOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// TrailingDelta set trailingDelta
func (s *NewOCOService) TrailingDelta(trailingDelta int) *NewOCOService {
	s.trailingDelta = &trailingDelta
	return s
}

// StopClientOrderId set stopClientOrderId
func (s *NewOCOService) StopClientOrderId(stopClientOrderId string) *NewOCOService {
	s.stopClientOrderId = &stopClientOrderId
	return s
}

// StopPrice set stopPrice
func (s *NewOCOService) StopPrice(stopPrice float64) *NewOCOService {
	s.stopPrice = stopPrice
	return s
}

// StopStrategyId set stopStrategyId
func (s *NewOCOService) StopStrategyId(stopStrategyId int) *NewOCOService {
	s.stopStrategyId = &stopStrategyId
	return s
}

// StopStrategyType set stopStrategyType
func (s *NewOCOService) StopStrategyType(stopStrategyType int) *NewOCOService {
	s.stopStrategyType = &stopStrategyType
	return s
}

// StopLimitPrice set stopLimitPrice
func (s *NewOCOService) StopLimitPrice(stopLimitPrice float64) *NewOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stopIcebergQty
func (s *NewOCOService) StopIcebergQty(stopIcebergQty float64) *NewOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *NewOCOService) StopLimitTimeInForce(stopLimitTimeInForce string) *NewOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set newOrderRespType
func (s *NewOCOService) NewOrderRespType(newOrderRespType string) *NewOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// selfTradePreventionMode set selfTradePreventionMode
func (s *NewOCOService) SelfTradePreventionMode(selfTradePreventionMode string) *NewOCOService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// Do send request
func (s *NewOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OrderOCOResponse, err error) {
	r := newSpotRequest("/api/v3/order/oco")

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}
	if s.side == "" {
		err = fmt.Errorf("%w: side", apierrors.ErrMissingParameter)
		return
	}
	if s.quantity == 0 {
		err = fmt.Errorf("%w: quantity", apierrors.ErrMissingParameter)
		return
	}
	if s.price == 0 {
		err = fmt.Errorf("%w: price", apierrors.ErrMissingParameter)
		return
	}
	if s.stopPrice == 0 {
		err = fmt.Errorf("%w: stopPrice", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("symbol", s.symbol)
	r.SetParam("side", s.side)
	r.SetParam("quantity", s.quantity)
	r.SetParam("price", s.price)
	r.SetParam("stopPrice", s.stopPrice)

	r.SetParam("listClientOrderId", s.listClientOrderId)
	r.SetParam("limitClientOrderId", s.limitClientOrderId)
	r.SetParam("limitStrategyId", s.limitStrategyId)
	r.SetParam("limitStrategyType", s.limitStrategyType)
	r.SetParam("limitIcebergQty", s.limitIcebergQty)
	r.SetParam("trailingDelta", s.trailingDelta)
	r.SetParam("stopClientOrderId", s.stopClientOrderId)
	r.SetParam("stopStrategyId", s.stopStrategyId)
	r.SetParam("stopStrategyType", s.stopStrategyType)
	r.SetParam("stopLimitPrice", s.stopLimitPrice)
	r.SetParam("stopIcebergQty", s.stopIcebergQty)
	r.SetParam("stopLimitTimeInForce", s.stopLimitTimeInForce)
	r.SetParam("newOrderRespType", s.newOrderRespType)
	r.SetParam("selfTradePreventionMode", s.selfTradePreventionMode)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

type OrderOCOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []OrderResponse `json:"orderReports"`
}
