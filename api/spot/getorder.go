package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Query Order (USER_DATA)
// Binance Query Order (USER_DATA) (GET /api/v3/order)
// GetOrderService get order
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
		err = apierrors.ErrMissingSymbol
		return
	}

	if (s.orderId == nil && s.origClientOrderId == nil) || (s.orderId != nil && s.origClientOrderId != nil) {
		err = apierrors.ErrEitherOrderIdOrOrigClientOrderId
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
type GetOrderResponse struct {
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
