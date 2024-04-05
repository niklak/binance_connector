package spot

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Query Order (USER_DATA) (GET /api/v3/order)
// GetOrderService get order
//
//gen:new_service
type GetOrderService struct {
	C                 *connector.Connector
	symbol            string
	orderId           *int64
	origClientOrderId *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetOrderService) OrderId(orderId int64) *GetOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *GetOrderService) OrigClientOrderId(origClientOrderId string) *GetOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *GetOrderResponse, err error) {

	r := request.New(
		"/api/v3/order",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.RequiredOneOfParams([]string{"orderId", "origClientOrderId"}),
		request.SetParam("symbol", s.symbol),
		request.SetParam("orderId", s.orderId),
		request.SetParam("origClientOrderId", s.origClientOrderId),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(GetOrderResponse)
	err = json.Unmarshal(data, &res)
	return
}

// Create GetOrderResponse
type GetOrderResponse = OrderResponse
