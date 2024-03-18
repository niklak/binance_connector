package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/client"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Cancel Order endpoint (DELETE /api/v3/order)
// CancelOrderService cancel order
type CancelOrderService struct {
	C                  *client.Connector
	symbol             string
	orderId            *int64
	origClientOrderId  *string
	newClientOrderId   *string
	cancelRestrictions *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *CancelOrderService) OrderId(orderId int64) *CancelOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *CancelOrderService) OrigClientOrderId(origClientOrderId string) *CancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelOrderService) NewClientOrderId(newClientOrderId string) *CancelOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// CancelRestrictions set cancelRestrictions
func (s *CancelOrderService) CancelRestrictions(cancelRestrictions string) *CancelOrderService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *CancelOrderResponse, err error) {
	r := &request.Request{
		Method:   http.MethodDelete,
		Endpoint: "/api/v3/order",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	if s.symbol == "" {
		err = apierrors.ErrMissingSymbol
		return
	}

	r.SetParam("symbol", s.symbol)
	r.SetParam("orderId", *s.orderId)
	r.SetParam("origClientOrderId", *s.origClientOrderId)
	r.SetParam("newClientOrderId", *s.newClientOrderId)
	r.SetParam("cancelRestrictions", *s.cancelRestrictions)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)

	err = json.Unmarshal(data, res)
	return
}

// Create CancelOrderResponse
type CancelOrderResponse struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CumulativeQuoteQty  string `json:"cumulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	SelfTradePrevention string `json:"selfTradePrevention"`
	IcebergQty          string `json:"icebergQty,omitempty"`
	PreventedMatchId    int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity   string `json:"preventedQuantity,omitempty"`
	StopPrice           string `json:"stopPrice,omitempty"`
	StrategyId          int64  `json:"strategyId,omitempty"`
	StrategyType        int64  `json:"strategyType,omitempty"`
	TrailingDelta       string `json:"trailingDelta,omitempty"`
	TrailingTime        int64  `json:"trailingTime,omitempty"`
}
