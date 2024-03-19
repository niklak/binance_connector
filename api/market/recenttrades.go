package market

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Recent Trades List endpoint (GET /api/v3/trades)
type RecentTradesList struct {
	C      *connector.Connector
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *RecentTradesList) Symbol(symbol string) *RecentTradesList {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *RecentTradesList) Limit(limit int) *RecentTradesList {
	s.limit = &limit
	return s
}

// Send the request
func (s *RecentTradesList) Do(ctx context.Context, opts ...request.RequestOption) (res []*RecentTradesListResponse, err error) {
	r := newMarketRequest("/api/v3/trades")
	if s.symbol == "" {
		return nil, fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
	}
	r.SetParam("symbol", s.symbol)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*RecentTradesListResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// RecentTradesListResponse define recent trades list response
type RecentTradesListResponse struct {
	Id           uint64 `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         uint64 `json:"time"`
	QuoteQty     string `json:"quoteQty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBest       bool   `json:"isBestMatch"`
}
