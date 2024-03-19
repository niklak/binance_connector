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

// Binance Get all account orders; active, canceled, or filled (GET /api/v3/allOrders)
// GetAllOrdersService get all orders
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
func (s *GetAllOrdersService) Do(ctx context.Context, opts ...request.RequestOption) (res []*AllOrdersResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/allOrders",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("symbol", s.symbol)
	r.SetParam("orderId", s.orderId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*AllOrdersResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// Create NewAllOrdersResponse
type AllOrdersResponse struct {
	Symbol                  string `json:"symbol"`
	ListClientOrderId       string `json:"listClientOrderId"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cumulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}
