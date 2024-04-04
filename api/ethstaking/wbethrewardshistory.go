package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get WBETH rewards history(USER_DATA) (GET /sapi/v1/eth-staking/eth/history/wbethRewardsHistory)
//
//gen:new_service
type WBETHRewardsHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *WBETHRewardsHistoryService) StartTime(startTime uint64) *WBETHRewardsHistoryService {
	s.startTime = &startTime
	return s
}

func (s *WBETHRewardsHistoryService) EndTime(endTime uint64) *WBETHRewardsHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WBETHRewardsHistoryService) Current(current uint64) *WBETHRewardsHistoryService {
	s.current = &current
	return s
}

func (s *WBETHRewardsHistoryService) Size(size uint64) *WBETHRewardsHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *WBETHRewardsHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *WBETHRewardsHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/history/wbethRewardsHistory",
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
	res = new(WBETHRewardsHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

type WBETHRewardsHistoryRow struct {
	Time                 int64  `json:"time"`
	AmountInETH          string `json:"amountInETH"`
	Holding              string `json:"holding"`
	HoldingInETH         string `json:"holdingInETH"`
	AnnualPercentageRate string `json:"annualPercentageRate"`
}

// WBETHRewardsHistoryResponse define response of WBETHRewardsHistoryService.Do
type WBETHRewardsHistoryResponse struct {
	Total           uint64                    `json:"total"`
	EstRewardsInETH string                    `json:"estRewardsInETH"`
	Rows            []*WBETHRewardsHistoryRow `json:"rows"`
}
