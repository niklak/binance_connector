package account

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Query Prevented Matches (USER_DATA)
// QueryPreventedMatchesService query prevented matches
//
//gen:new_service
type QueryPreventedMatchesService struct {
	C                    *connector.Connector
	symbol               string
	preventMatchId       *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                *int
}

// Symbol set symbol
func (s *QueryPreventedMatchesService) Symbol(symbol string) *QueryPreventedMatchesService {
	s.symbol = symbol
	return s
}

// PreventMatchId set preventMatchId
func (s *QueryPreventedMatchesService) PreventMatchId(preventMatchId int64) *QueryPreventedMatchesService {
	s.preventMatchId = &preventMatchId
	return s
}

// OrderId set orderId
func (s *QueryPreventedMatchesService) OrderId(orderId int64) *QueryPreventedMatchesService {
	s.orderId = &orderId
	return s
}

// FromPreventedMatchId set fromPreventedMatchId
func (s *QueryPreventedMatchesService) FromPreventedMatchId(fromPreventedMatchId int64) *QueryPreventedMatchesService {
	s.fromPreventedMatchId = &fromPreventedMatchId
	return s
}

// Limit set limit
func (s *QueryPreventedMatchesService) Limit(limit int) *QueryPreventedMatchesService {
	s.limit = &limit
	return s
}

// Do send request
func (s *QueryPreventedMatchesService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryPreventedMatchesResponse, err error) {

	r := newAccountRequest("/api/v3/myPreventedMatches")

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("symbol", s.symbol)
	r.SetParam("preventedMatchId", s.preventMatchId)
	r.SetParam("orderId", s.orderId)
	r.SetParam("fromPreventedMatchId", s.fromPreventedMatchId)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryPreventedMatchesResponse)
	err = json.Unmarshal(data, res)
	return
}

// Create QueryPreventedMatchesResponse
type QueryPreventedMatchesResponse struct {
	PreventedMatches []struct {
		Symbol                  string `json:"symbol"`
		PreventedMatchId        int64  `json:"preventedMatchId"`
		TakerOrderId            int64  `json:"takerOrderId"`
		MakerOrderId            int64  `json:"makerOrderId"`
		TradeGroupId            int64  `json:"tradeGroupId"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode"`
		Price                   string `json:"price"`
		MakerPreventedQuantity  string `json:"makerPreventedQuantity"`
		TransactTime            uint64 `json:"transactTime"`
	} `json:"preventedMatches"`
}
