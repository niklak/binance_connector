package market

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Old Trade Lookup endpoint (GET /api/v3/historicalTrades)
type HistoricalTradeLookup struct {
	C      *connector.Connector
	symbol string
	limit  *uint
	fromId *int64
}

// Symbol set symbol
func (s *HistoricalTradeLookup) Symbol(symbol string) *HistoricalTradeLookup {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *HistoricalTradeLookup) Limit(limit uint) *HistoricalTradeLookup {
	s.limit = &limit
	return s
}

// FromId set fromId
func (s *HistoricalTradeLookup) FromId(fromId int64) *HistoricalTradeLookup {
	s.fromId = &fromId
	return s
}

// Send the request
func (s *HistoricalTradeLookup) Do(ctx context.Context, opts ...request.RequestOption) (res []*RecentTradesListResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/historicalTrades",
		SecType:  request.SecTypeAPIKey,
	}
	r.Init()

	if s.symbol == "" {
		return nil, fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
	}

	r.SetParam("symbol", s.symbol)
	r.SetParam("limit", s.limit)
	r.SetParam("fromId", s.fromId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}
