package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get WBETH Rate History (USER_DATA) (GET /sapi/v1/eth-staking/eth/history/rateHistory)
//
//gen:new_service
type WBETHRateHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *WBETHRateHistoryService) StartTime(startTime uint64) *WBETHRateHistoryService {
	s.startTime = &startTime
	return s
}

func (s *WBETHRateHistoryService) EndTime(endTime uint64) *WBETHRateHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WBETHRateHistoryService) Current(current uint64) *WBETHRateHistoryService {
	s.current = &current
	return s
}

func (s *WBETHRateHistoryService) Size(size uint64) *WBETHRateHistoryService {
	s.size = &size
	return s
}

func (s *WBETHRateHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *WBETHRateHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/history/rateHistory",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(WBETHRateHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

type WBETHRateHistoryRow struct {
	AnnualPercentageRate string `json:"annualPercentageRate"`
	ExchangeRate         string `json:"exchangeRate"`
	Time                 int64  `json:"time"`
}

type WBETHRateHistoryResponse struct {
	Total uint64                 `json:"total"`
	Rows  []*WBETHRateHistoryRow `json:"rows"`
}
