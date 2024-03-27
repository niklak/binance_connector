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

// Query Order (USER_DATA)
// Binance Query Order (USER_DATA) (GET /api/v3/order)
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
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/order",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	if (s.orderId == nil && s.origClientOrderId == nil) || (s.orderId != nil && s.origClientOrderId != nil) {
		err = fmt.Errorf("%w: either origClientOrderId or orderId", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("symbol", s.symbol)
	r.SetParam("orderId", s.orderId)
	r.SetParam("origClientOrderId", s.origClientOrderId)

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
