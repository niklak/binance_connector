package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

//Get BETH rewards distribution history(USER_DATA) (GET /sapi/v1/eth-staking/eth/history/rewardsHistory)
//
//gen:new_service

type BETHRewardsDistributionHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *BETHRewardsDistributionHistoryService) StartTime(startTime uint64) *BETHRewardsDistributionHistoryService {
	s.startTime = &startTime
	return s
}

func (s *BETHRewardsDistributionHistoryService) EndTime(endTime uint64) *BETHRewardsDistributionHistoryService {
	s.endTime = &endTime
	return s
}

func (s *BETHRewardsDistributionHistoryService) Current(current uint64) *BETHRewardsDistributionHistoryService {
	s.current = &current
	return s
}

func (s *BETHRewardsDistributionHistoryService) Size(size uint64) *BETHRewardsDistributionHistoryService {
	s.size = &size
	return s
}

func (s *BETHRewardsDistributionHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *BETHRewardsDistributionHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/history/rewardsHistory",
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
	res = new(BETHRewardsDistributionHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

type BETHRewardHistoryRow struct {
	Time                 int64  `json:"time"`
	Asset                string `json:"asset"`
	Holding              string `json:"holding"`
	Amount               string `json:"amount"`
	AnnualPercentageRate string `json:"annualPercentageRate"`
	Status               string `json:"status"`
}

type BETHRewardsDistributionHistoryResponse struct {
	Total int64                   `json:"total"`
	Rows  []*BETHRewardHistoryRow `json:"rows"`
}
