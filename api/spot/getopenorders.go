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

// Create OpenOrdersResponse
type OrderResponse struct {
	Symbol                  string `json:"symbol"`
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
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}
