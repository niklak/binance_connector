package market

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance UI Klines GET /api/v3/uiKlines
type UiKlines struct {
	C         *connector.Connector
	symbol    string
	interval  string
	limit     *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *UiKlines) Symbol(symbol string) *UiKlines {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *UiKlines) Interval(interval string) *UiKlines {
	s.interval = interval
	return s
}

// Limit set limit
func (s *UiKlines) Limit(limit int) *UiKlines {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *UiKlines) StartTime(startTime uint64) *UiKlines {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UiKlines) EndTime(endTime uint64) *UiKlines {
	s.endTime = &endTime
	return s
}

// Send the request
func (s *UiKlines) Do(ctx context.Context, opts ...request.RequestOption) (res []*UiKlinesResponse, err error) {
	r := newMarketRequest("/api/v3/uiKlines")

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
		return nil, err
	}
	var uiklinesResponseArray UiKlinesResponseArray
	err = json.Unmarshal(data, &uiklinesResponseArray)
	if err != nil {
		return
	}
	res = make([]*UiKlinesResponse, 0)
	for _, uikline := range uiklinesResponseArray {
		openTime := uikline[0].(float64)
		open := uikline[1].(string)
		high := uikline[2].(string)
		low := uikline[3].(string)
		close := uikline[4].(string)
		volume := uikline[5].(string)
		closeTime := uikline[6].(float64)
		quoteAssetVolume := uikline[7].(string)
		numberOfTrades := uikline[8].(float64)
		takerBuyBaseAssetVolume := uikline[9].(string)
		takerBuyQuoteAssetVolume := uikline[10].(string)

		// create a KlinesResponse struct using the parsed fields
		uiklinesResponse := &UiKlinesResponse{
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
		res = append(res, uiklinesResponse)
	}
	return
}

type UiKlinesResponseArray [][]interface{}

// Define UiKlines response data
type UiKlinesResponse struct {
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
