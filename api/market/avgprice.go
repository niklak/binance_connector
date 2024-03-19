package market

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Current Average Price (GET /api/v3/avgPrice)
type AvgPrice struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *AvgPrice) Symbol(symbol string) *AvgPrice {
	s.symbol = symbol
	return s
}

// Send the request
func (s *AvgPrice) Do(ctx context.Context, opts ...request.RequestOption) (res *AvgPriceResponse, err error) {
	r := newMarketRequest("/api/v3/avgPrice")

	if s.symbol == "" {
		return nil, fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
	}

	r.SetParam("symbol", s.symbol)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(AvgPriceResponse)
	err = json.Unmarshal(data, res)

	return
}

// Define AvgPrice response data
type AvgPriceResponse struct {
	Mins  uint64 `json:"mins"`
	Price string `json:"price"`
}
