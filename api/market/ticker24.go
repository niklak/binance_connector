package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/helpers"
	"github.com/niklak/binance_connector/request"
)

// Binance 24hr Ticker Price Change Statistics (GET /api/v3/ticker/24hr)
//
//gen:new_service
type Ticker24hr struct {
	C       *connector.Connector
	symbol  *string
	symbols *[]string
}

// Symbol set symbol
func (s *Ticker24hr) Symbol(symbol string) *Ticker24hr {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *Ticker24hr) Symbols(symbols []string) *Ticker24hr {
	s.symbols = &symbols
	return s
}

// Send the request
func (s *Ticker24hr) Do(ctx context.Context, opts ...request.RequestOption) (res []*Ticker24hrResponse, err error) {

	r := request.New("/api/v3/ticker/24hr")

	if s.symbol != nil {
		r.SetParam("symbol", s.symbol)
	} else if s.symbols != nil {
		symbols := helpers.StringifyStringSlice(*s.symbols)
		r.SetParam("symbols", symbols)
	}

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	if s.symbols != nil || (s.symbol == nil && s.symbols == nil) {
		if err = json.Unmarshal(data, &res); err != nil {
			return
		}
	} else if s.symbol != nil {
		dst := &Ticker24hrResponse{}
		if err = json.Unmarshal(data, dst); err != nil {
			return
		}
		res = append(res, dst)
	}
	return
}

// Define Ticker24hr response data
type Ticker24hrResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint64 `json:"openTime"`
	CloseTime          uint64 `json:"closeTime"`
	FirstId            uint64 `json:"firstId"`
	LastId             uint64 `json:"lastId"`
	Count              uint64 `json:"count"`
}
