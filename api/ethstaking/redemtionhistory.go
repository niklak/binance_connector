package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get ETH redemption history (USER_DATA) (GET /sapi/v1/eth-staking/eth/history/redemptionHistory)
//
//gen:new_service
type ETHRedemptionHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *ETHRedemptionHistoryService) StartTime(startTime uint64) *ETHRedemptionHistoryService {
	s.startTime = &startTime
	return s
}

func (s *ETHRedemptionHistoryService) EndTime(endTime uint64) *ETHRedemptionHistoryService {
	s.endTime = &endTime
	return s
}

func (s *ETHRedemptionHistoryService) Current(current uint64) *ETHRedemptionHistoryService {
	s.current = &current
	return s
}

func (s *ETHRedemptionHistoryService) Size(size uint64) *ETHRedemptionHistoryService {
	s.size = &size
	return s
}

func (s *ETHRedemptionHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *ETHHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/history/redemptionHistory",
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
