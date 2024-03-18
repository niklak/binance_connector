package market

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/client"
	"github.com/niklak/binance_connector/internal/helpers"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Symbol Price Ticker (GET /api/v3/ticker/price)
type TickerPrice struct {
	C       *client.Connector
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
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/ticker/price",
		SecType:  request.SecTypeNone,
	}
	r.Init()
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		symbols := helpers.StringifyStringSlice(*s.symbols)
		r.SetParam("symbols", symbols)
	} else {
		err = apierrors.ErrMissingSymbol
		return
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return
	}

	if s.symbols != nil {
		if err = json.Unmarshal(data, &res); err != nil {
			return
		}
	} else if s.symbol != nil {
		dst := &TickerPriceResponse{}
		if err = json.Unmarshal(data, dst); err != nil {
			return
		}
		res = append(res, dst)
	}

	return
}

// Define TickerPrice response data
type TickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
