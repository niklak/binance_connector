package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Binance Rolling window price change statistics (GET /api/v3/ticker)
//
//gen:new_service
type Ticker struct {
	C          *connector.Connector
	symbol     string
	windowSize *string
	tickerType *string
}

// Symbol set symbol
func (s *Ticker) Symbol(symbol string) *Ticker {
	s.symbol = symbol
	return s
}

// WindowSize set windowSize
func (s *Ticker) WindowSize(windowSize string) *Ticker {
	s.windowSize = &windowSize
	return s
}

// Type set type
func (s *Ticker) Type(tickerType string) *Ticker {
	s.tickerType = &tickerType
	return s
}

// Send the request
func (s *Ticker) Do(ctx context.Context, opts ...request.RequestOption) (res *TickerResponse, err error) {

	r := request.New(
		"/api/v3/ticker",
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("windowSize", s.windowSize)
	r.SetParam("type", s.tickerType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(TickerResponse)
	err = json.Unmarshal(data, res)
	return
}

// Define Ticker response data
type TickerResponse struct {
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
