package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Cancel Order endpoint (DELETE /api/v3/order)
// CancelOrderService cancel order
//
//gen:new_service
type CancelOrderService struct {
	C                  *connector.Connector
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

	r := request.New(
		"/api/v3/order",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.RequiredOneOfParams([]string{"orderId", "origClientOrderId"}),
		request.SetParam("symbol", s.symbol),
		request.SetParam("orderId", s.orderId),
		request.SetParam("origClientOrderId", s.origClientOrderId),
		request.SetParam("newClientOrderId", s.newClientOrderId),
		request.SetParam("cancelRestrictions", s.cancelRestrictions),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)

	err = json.Unmarshal(data, res)
	return
}

// Create CancelOrderResponse
type CancelOrderResponse = OrderResponse
