package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get ETH staking history (USER_DATA) (GET /sapi/v1/eth-staking/eth/history/stakingHistory)
//
//gen:new_service
type ETHStakingHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *ETHStakingHistoryService) StartTime(startTime uint64) *ETHStakingHistoryService {
	s.startTime = &startTime
	return s
}

func (s *ETHStakingHistoryService) EndTime(endTime uint64) *ETHStakingHistoryService {
	s.endTime = &endTime
	return s
}

func (s *ETHStakingHistoryService) Current(current uint64) *ETHStakingHistoryService {
	s.current = &current
	return s
}

func (s *ETHStakingHistoryService) Size(size uint64) *ETHStakingHistoryService {
	s.size = &size
	return s
}

func (s *ETHStakingHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *ETHHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/history/stakingHistory",
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
	res = new(ETHHistoryResponse)
	err = json.Unmarshal(data, res)

	return
}

type ETHHistoryRow struct {
	Time             int64  `json:"time"`
	ArrivalTime      int64  `json:"arrivalTime"`
	Asset            string `json:"asset"`
	Amount           string `json:"amount"`
	Status           string `json:"status"`
	DistributeAsset  string `json:"distributeAsset"`
	DistributeAmount string `json:"distributeAmount"`
	ConversionRatio  string `json:"conversionRatio"`
}

type ETHHistoryResponse struct {
	Total uint64           `json:"total"`
	Rows  []*ETHHistoryRow `json:"rows"`
}
