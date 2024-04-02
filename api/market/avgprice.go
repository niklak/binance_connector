package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Current Average Price (GET /api/v3/avgPrice)
//
//gen:new_service
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

	r := request.New(
		"/api/v3/avgPrice",
		request.RequiredParams("symbol"),
	)

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
