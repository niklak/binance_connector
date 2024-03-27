package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Get current open orders (GET /api/v3/openOrders)
// GetOpenOrdersService get open orders
//
//gen:new_service
type GetOpenOrdersService struct {
	C      *connector.Connector
	symbol *string
}

// Symbol set symbol
func (s *GetOpenOrdersService) Symbol(symbol string) *GetOpenOrdersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetOpenOrdersService) Do(ctx context.Context, opts ...request.RequestOption) (res []*OrderResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/openOrders",
		SecType:  request.SecTypeSigned,
	}
	r.Init()
	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*OrderResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}
