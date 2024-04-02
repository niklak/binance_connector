package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Get all account orders; active, canceled, or filled (GET /api/v3/allOrders)
// GetAllOrdersService get all orders
//
//gen:new_service
type GetAllOrdersService struct {
	C         *connector.Connector
	symbol    string
	orderId   *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// Symbol set symbol
func (s *GetAllOrdersService) Symbol(symbol string) *GetAllOrdersService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetAllOrdersService) OrderId(orderId int64) *GetAllOrdersService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *GetAllOrdersService) StartTime(startTime uint64) *GetAllOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetAllOrdersService) EndTime(endTime uint64) *GetAllOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAllOrdersService) Limit(limit int) *GetAllOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAllOrdersService) Do(ctx context.Context, opts ...request.RequestOption) (res []*OrderResponse, err error) {

	r := request.New(
		"/api/v3/allOrders",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("orderId", s.orderId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*OrderResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}
