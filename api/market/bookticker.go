package market

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/helpers"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Symbol Order Book Ticker (GET /api/v3/ticker/bookTicker)
type TickerBookTicker struct {
	C       *connector.Connector
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *TickerBookTicker) Symbol(symbol string) *TickerBookTicker {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *TickerBookTicker) Symbols(symbols []string) *TickerBookTicker {
	s.symbols = &symbols
	return s
}

func (s *TickerBookTicker) Do(ctx context.Context, opts ...request.RequestOption) (res []*TickerBookTickerResponse, err error) {

	r := newMarketRequest("/api/v3/ticker/bookTicker")

	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		symbols := helpers.StringifyStringSlice(*s.symbols)
		r.SetParam("symbols", symbols)
	} else {
		err = fmt.Errorf("%w: either symbol or symbols", apierrors.ErrMissingParameter)
		return
	}
	data, err := s.C.CallAPI(ctx, r, opts...)

	if err != nil {
		return []*TickerBookTickerResponse{}, err
	}
	var raw json.RawMessage
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return
	}

	if s.symbols != nil {
		if err = json.Unmarshal(data, &res); err != nil {
			return
		}
	} else if s.symbol != nil {
		dst := &TickerBookTickerResponse{}
		if err = json.Unmarshal(data, dst); err != nil {
			return
		}
		res = append(res, dst)
	}

	return res, nil
}

// Define TickerBookTicker response data
type TickerBookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}
