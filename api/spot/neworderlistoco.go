package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// New Order List - OCO (TRADE) (POST /api/v3/orderList/oco)
//
//gen:new_service
type CreateOrderListOCOService struct {
	C                       *connector.Connector
	symbol                  string
	listClientOrderId       *string
	side                    string
	quantity                string
	aboveType               string
	aboveClientOrderId      *string
	aboveIcebergQty         *int64
	abovePrice              *float64
	aboveStopPrice          *float64
	aboveTrailingDelta      *int64
	aboveTimeInForce        *float64
	aboveStrategyId         *int64
	aboveStrategyType       *int64
	belowType               string
	belowClientOrderId      *string
	belowIcebergQty         *int64
	belowPrice              *float64
	belowStopPrice          *float64
	belowTrailingDelta      *int64
	belowTimeInForce        *float64
	belowStrategyId         *int64
	belowStrategyType       *int64
	newClientOrderId        *string
	newOrderRespType        *string
	selfTradePreventionMode string
}

// Symbol set symbol
func (s *CreateOrderListOCOService) Symbol(symbol string) *CreateOrderListOCOService {
	s.symbol = symbol
	return s
}

// ListClientOrderId set listClientOrderId
func (s *CreateOrderListOCOService) ListClientOrderId(listClientOrderId string) *CreateOrderListOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// Side set side
func (s *CreateOrderListOCOService) Side(side string) *CreateOrderListOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *CreateOrderListOCOService) Quantity(quantity string) *CreateOrderListOCOService {
	s.quantity = quantity
	return s
}

// AboveType set aboveType
func (s *CreateOrderListOCOService) AboveType(aboveType string) *CreateOrderListOCOService {
	s.aboveType = aboveType
	return s
}

// AboveClientOrderId set aboveClientOrderId
func (s *CreateOrderListOCOService) AboveClientOrderId(aboveClientOrderId string) *CreateOrderListOCOService {
	s.aboveClientOrderId = &aboveClientOrderId
	return s
}

// AboveIcebergQty set aboveIcebergQty
func (s *CreateOrderListOCOService) AboveIcebergQty(aboveIcebergQty int64) *CreateOrderListOCOService {
	s.aboveIcebergQty = &aboveIcebergQty
	return s
}

// AbovePrice set abovePrice
func (s *CreateOrderListOCOService) AbovePrice(abovePrice float64) *CreateOrderListOCOService {
	s.abovePrice = &abovePrice
	return s
}

// AboveStopPrice set aboveStopPrice
func (s *CreateOrderListOCOService) AboveStopPrice(aboveStopPrice float64) *CreateOrderListOCOService {
	s.aboveStopPrice = &aboveStopPrice
	return s
}

// AboveTrailingDelta set aboveTrailingDelta
func (s *CreateOrderListOCOService) AboveTrailingDelta(aboveTrailingDelta int64) *CreateOrderListOCOService {
	s.aboveTrailingDelta = &aboveTrailingDelta
	return s
}

// AboveTimeInForce set aboveTimeInForce
func (s *CreateOrderListOCOService) AboveTimeInForce(aboveTimeInForce float64) *CreateOrderListOCOService {
	s.aboveTimeInForce = &aboveTimeInForce
	return s
}

// AboveStrategyId set aboveStrategyId
func (s *CreateOrderListOCOService) AboveStrategyId(aboveStrategyId int64) *CreateOrderListOCOService {
	s.aboveStrategyId = &aboveStrategyId
	return s
}

// AboveStrategyType set aboveStrategyType

func (s *CreateOrderListOCOService) AboveStrategyType(aboveStrategyType int64) *CreateOrderListOCOService {
	s.aboveStrategyType = &aboveStrategyType
	return s
}

// BelowType set belowType
func (s *CreateOrderListOCOService) BelowType(belowType string) *CreateOrderListOCOService {
	s.belowType = belowType
	return s
}

// BelowClientOrderId set belowClientOrderId
func (s *CreateOrderListOCOService) BelowClientOrderId(belowClientOrderId string) *CreateOrderListOCOService {
	s.belowClientOrderId = &belowClientOrderId
	return s
}

// BelowIcebergQty set belowIcebergQty
func (s *CreateOrderListOCOService) BelowIcebergQty(belowIcebergQty int64) *CreateOrderListOCOService {
	s.belowIcebergQty = &belowIcebergQty
	return s
}

// BelowPrice set belowPrice
func (s *CreateOrderListOCOService) BelowPrice(belowPrice float64) *CreateOrderListOCOService {
	s.belowPrice = &belowPrice
	return s
}

// BelowStopPrice set belowStopPrice
func (s *CreateOrderListOCOService) BelowStopPrice(belowStopPrice float64) *CreateOrderListOCOService {
	s.belowStopPrice = &belowStopPrice
	return s
}

// BelowTrailingDelta set belowTrailingDelta
func (s *CreateOrderListOCOService) BelowTrailingDelta(belowTrailingDelta int64) *CreateOrderListOCOService {
	s.belowTrailingDelta = &belowTrailingDelta
	return s
}

// BelowTimeInForce set belowTimeInForce
func (s *CreateOrderListOCOService) BelowTimeInForce(belowTimeInForce float64) *CreateOrderListOCOService {
	s.belowTimeInForce = &belowTimeInForce
	return s
}

// BelowStrategyId set belowStrategyId
func (s *CreateOrderListOCOService) BelowStrategyId(belowStrategyId int64) *CreateOrderListOCOService {
	s.belowStrategyId = &belowStrategyId
	return s
}

// BelowStrategyType set belowStrategyType
func (s *CreateOrderListOCOService) BelowStrategyType(belowStrategyType int64) *CreateOrderListOCOService {
	s.belowStrategyType = &belowStrategyType
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CreateOrderListOCOService) NewClientOrderId(newClientOrderId string) *CreateOrderListOCOService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// NewOrderRespType set newOrderRespType
func (s *CreateOrderListOCOService) NewOrderRespType(newOrderRespType string) *CreateOrderListOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderListOCOService) SelfTradePreventionMode(selfTradePreventionMode string) *CreateOrderListOCOService {
	s.selfTradePreventionMode = selfTradePreventionMode
	return s
}

// Do send request
func (s *CreateOrderListOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OrderOCOResponse, err error) {

	r := request.New(
		"/api/v3/orderList/oco",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol", "side", "quantity", "aboveType", "belowType"),
		request.SetParam("symbol", s.symbol),
		request.SetParam("side", s.side),
		request.SetParam("quantity", s.quantity),
		request.SetParam("aboveType", s.aboveType),
		request.SetParam("belowType", s.belowType),
		request.SetParam("selfTradePreventionMode", s.selfTradePreventionMode),
		request.SetParam("listClientOrderId", s.listClientOrderId),
		request.SetParam("aboveClientOrderId", s.aboveClientOrderId),
		request.SetParam("aboveIcebergQty", s.aboveIcebergQty),
		request.SetParam("abovePrice", s.abovePrice),
		request.SetParam("aboveStopPrice", s.aboveStopPrice),
		request.SetParam("aboveTrailingDelta", s.aboveTrailingDelta),
		request.SetParam("aboveTimeInForce", s.aboveTimeInForce),
		request.SetParam("aboveStrategyId", s.aboveStrategyId),
		request.SetParam("aboveStrategyType", s.aboveStrategyType),
		request.SetParam("belowClientOrderId", s.belowClientOrderId),
		request.SetParam("belowIcebergQty", s.belowIcebergQty),
		request.SetParam("belowPrice", s.belowPrice),
		request.SetParam("belowStopPrice", s.belowStopPrice),
		request.SetParam("belowTrailingDelta", s.belowTrailingDelta),
		request.SetParam("belowTimeInForce", s.belowTimeInForce),
		request.SetParam("belowStrategyId", s.belowStrategyId),
		request.SetParam("belowStrategyType", s.belowStrategyType),
		request.SetParam("newClientOrderId", s.newClientOrderId),
		request.SetParam("newOrderRespType", s.newOrderRespType),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderOCOResponse)
	err = json.Unmarshal(data, res)
	return
}
