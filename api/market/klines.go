package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Binance Kline/Candlestick Data endpoint (GET /api/v3/klines)
//
//gen:new_service
type Klines struct {
	C         *connector.Connector
	symbol    string
	interval  string
	limit     *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *Klines) Symbol(symbol string) *Klines {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *Klines) Interval(interval string) *Klines {
	s.interval = interval
	return s
}

// Limit set limit
func (s *Klines) Limit(limit int) *Klines {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *Klines) StartTime(startTime uint64) *Klines {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *Klines) EndTime(endTime uint64) *Klines {
	s.endTime = &endTime
	return s
}

func (s *Klines) Do(ctx context.Context, opts ...request.RequestOption) (res []*KlinesResponse, err error) {

	r := request.New(
		"/api/v3/klines",
		request.RequiredParams("symbol", "interval"),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("interval", s.interval)

	r.SetParam("limit", s.limit)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	var klinesResponseArray KlinesResponseArray

	if err = json.Unmarshal(data, &klinesResponseArray); err != nil {
		return
	}

	res = make([]*KlinesResponse, 0)
	for _, kline := range klinesResponseArray {

		// create a KlinesResponse struct using the parsed fields
		klinesResponse := (&KlinesResponse{}).fromRawKline(kline)
		res = append(res, klinesResponse)
	}
	return
}

type KlinesResponseArray [][]interface{}

// Define Klines response data
type KlinesResponse struct {
	OpenTime                 uint64 `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                uint64 `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           uint64 `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

func (s *KlinesResponse) fromRawKline(kline []interface{}) *KlinesResponse {
	s.OpenTime = uint64(kline[0].(float64))
	s.Open = kline[1].(string)
	s.High = kline[2].(string)
	s.Low = kline[3].(string)
	s.Close = kline[4].(string)
	s.Volume = kline[5].(string)
	s.CloseTime = uint64(kline[6].(float64))
	s.QuoteAssetVolume = kline[7].(string)
	s.NumberOfTrades = uint64(kline[8].(float64))
	s.TakerBuyBaseAssetVolume = kline[9].(string)
	s.TakerBuyQuoteAssetVolume = kline[10].(string)
	return s
}
