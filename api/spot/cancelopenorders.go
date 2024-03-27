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

// Binance Cancel all open orders on a symbol (DELETE /api/v3/openOrders)
// CancelOpenOrdersService cancel open orders
//
//gen:new_service
type CancelOpenOrdersService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *CancelOpenOrdersService) Symbol(symbol string) *CancelOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelOpenOrdersService) Do(ctx context.Context, opts ...request.RequestOption) (res []*CancelOrderResponse, err error) {
	r := &request.Request{
		Method:   http.MethodDelete,
		Endpoint: "/api/v3/openOrders",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = make([]*CancelOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}
