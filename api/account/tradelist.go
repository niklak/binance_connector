package account

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Get trades for a specific account and symbol (USER_DATA) (GET /api/v3/myTrades)
// AccountTradeListService get trades for a specific account and symbol
//
//gen:new_service
type AccountTradeListService struct {
	C         *connector.Connector
	symbol    string
	orderId   *int64
	startTime *uint64
	endTime   *uint64
	fromId    *int64
	limit     *int
}

// Symbol set symbol
func (s *AccountTradeListService) Symbol(symbol string) *AccountTradeListService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *AccountTradeListService) OrderId(orderId int64) *AccountTradeListService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *AccountTradeListService) StartTime(startTime uint64) *AccountTradeListService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AccountTradeListService) EndTime(endTime uint64) *AccountTradeListService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *AccountTradeListService) FromId(fromId int64) *AccountTradeListService {
	s.fromId = &fromId
	return s
}

// Limit set limit
func (s *AccountTradeListService) Limit(limit int) *AccountTradeListService {
	s.limit = &limit
	return s
}

// Do send request
func (s *AccountTradeListService) Do(ctx context.Context, opts ...request.RequestOption) (res []*AccountTradeListResponse, err error) {

	r := request.New(
		"/api/v3/myTrades",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.SetParam("symbol", s.symbol),
		request.SetParam("orderId", s.orderId),
		request.SetParam("startTime", s.startTime),
		request.SetParam("endTime", s.endTime),
		request.SetParam("fromId", s.fromId),
		request.SetParam("limit", s.limit),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*AccountTradeListResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

//type AccountTradeListResponse []AccountTrade

type AccountTradeListResponse struct {
	Id              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	OrderId         int64  `json:"orderId"`
	OrderListId     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            uint64 `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}
