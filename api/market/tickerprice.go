package market

import (
	"context"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/helpers"
	"github.com/niklak/binance_connector/request"
)

// Binance Symbol Price Ticker (GET /api/v3/ticker/price)
//
//gen:new_service
type TickerPrice struct {
	C       *connector.Connector
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *TickerPrice) Symbol(symbol string) *TickerPrice {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *TickerPrice) Symbols(symbols []string) *TickerPrice {
	s.symbols = &symbols
	return s
}

// Send the request
func (s *TickerPrice) Do(ctx context.Context, opts ...request.RequestOption) (res []*TickerPriceResponse, err error) {

	r := request.New("/api/v3/ticker/price")

	if s.symbol != nil {
		r.SetParam("symbol", s.symbol)
	} else if s.symbols != nil {
		symbols := helpers.StringifyStringSlice(*s.symbols)
		r.SetParam("symbols", symbols)
	}

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return
	}

	expectArray := s.symbols != nil || (s.symbol == nil && s.symbols == nil)
	res, err = helpers.DeserializeIntoSlice[TickerPriceResponse](data, expectArray)
	return
}

// Define TickerPrice response data
type TickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
