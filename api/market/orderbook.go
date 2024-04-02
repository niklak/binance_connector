package market

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Order Book endpoint (GET /api/v3/depth)
//
//gen:new_service
type OrderBook struct {
	C      *connector.Connector
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *OrderBook) Symbol(symbol string) *OrderBook {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *OrderBook) Limit(limit int) *OrderBook {
	s.limit = &limit
	return s
}

// Send the request
func (s *OrderBook) Do(ctx context.Context, opts ...request.RequestOption) (res *OrderBookResponse, err error) {

	r := request.New(
		"/api/v3/depth",
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(OrderBookResponse)
	err = json.Unmarshal(data, res)
	return
}

// OrderBookResponse define order book response
type OrderBookResponse struct {
	LastUpdateId uint64         `json:"lastUpdateId"`
	Bids         [][]*big.Float `json:"bids"`
	Asks         [][]*big.Float `json:"asks"`
}
