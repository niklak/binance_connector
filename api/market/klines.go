package market

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Kline/Candlestick Data endpoint (GET /api/v3/klines)
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
	r := newMarketRequest("/api/v3/klines")

	if s.symbol == "" {
		return nil, fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
	}

	if s.interval == "" {
		return nil, fmt.Errorf("%w: interval", apierrors.ErrMissingParameter)
	}

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
		openTime := kline[0].(float64)
		open := kline[1].(string)
		high := kline[2].(string)
		low := kline[3].(string)
		close := kline[4].(string)
		volume := kline[5].(string)
		closeTime := kline[6].(float64)
		quoteAssetVolume := kline[7].(string)
		numberOfTrades := kline[8].(float64)
		takerBuyBaseAssetVolume := kline[9].(string)
		takerBuyQuoteAssetVolume := kline[10].(string)

		// create a KlinesResponse struct using the parsed fields
		klinesResponse := &KlinesResponse{
			OpenTime:                 uint64(openTime),
			Open:                     open,
			High:                     high,
			Low:                      low,
			Close:                    close,
			Volume:                   volume,
			CloseTime:                uint64(closeTime),
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           uint64(numberOfTrades),
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		}
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
