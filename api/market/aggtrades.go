package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Binance Compressed/Aggregate Trades List endpoint (GET /api/v3/aggTrades)
//
//gen:new_service
type AggTradesList struct {
	C         *connector.Connector
	symbol    string
	limit     *int
	fromId    *int
	startTime *uint64
	endTime   *uint64
}

// Symbol set symbol
func (s *AggTradesList) Symbol(symbol string) *AggTradesList {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *AggTradesList) Limit(limit int) *AggTradesList {
	s.limit = &limit
	return s
}

// FromId set fromId
func (s *AggTradesList) FromId(fromId int) *AggTradesList {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *AggTradesList) StartTime(startTime uint64) *AggTradesList {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AggTradesList) EndTime(endTime uint64) *AggTradesList {
	s.endTime = &endTime
	return s
}

// Send the request
func (s *AggTradesList) Do(ctx context.Context, opts ...request.RequestOption) (res []*AggTradesListResponse, err error) {

	r := request.New(
		"/api/v3/aggTrades",
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("limit", s.limit)
	r.SetParam("fromId", s.fromId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*AggTradesListResponse, 0)
	err = json.Unmarshal(data, &res)

	return
}

// AggTradesListResponse define compressed trades list response
type AggTradesListResponse struct {
	AggTradeId   uint64 `json:"a"`
	Price        string `json:"p"`
	Qty          string `json:"q"`
	FirstTradeId uint64 `json:"f"`
	LastTradeId  uint64 `json:"l"`
	Time         uint64 `json:"T"`
	IsBuyer      bool   `json:"m"`
	IsBest       bool   `json:"M"`
}
