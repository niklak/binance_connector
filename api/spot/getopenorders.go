package spot

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
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

	r := request.New(
		"/api/v3/openOrders",
		request.SecType(request.SecTypeSigned),
		request.SetParam("symbol", s.symbol),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*OrderResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}
