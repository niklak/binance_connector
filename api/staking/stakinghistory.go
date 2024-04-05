package staking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get Staking History(USER_DATA)
//
//gen:new_service
type StakingHistoryService struct {
	C         *connector.Connector
	product   string
	txnType   string
	asset     *string
	startTime *uint64
	endTime   *uint64
	current   *int64
	size      *int64
}

// Product set product
func (s *StakingHistoryService) Product(product string) *StakingHistoryService {
	s.product = product
	return s
}

// TxnType set txnType
func (s *StakingHistoryService) TxnType(txnType string) *StakingHistoryService {
	s.txnType = txnType
	return s
}

// Asset set asset
func (s *StakingHistoryService) Asset(asset string) *StakingHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *StakingHistoryService) StartTime(startTime uint64) *StakingHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *StakingHistoryService) EndTime(endTime uint64) *StakingHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *StakingHistoryService) Current(current int64) *StakingHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *StakingHistoryService) Size(size int64) *StakingHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *StakingHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res []*StakingHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/stakingRecord",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("product", "txnType"),
	)

	r.SetParam("product", s.product)
	r.SetParam("type", s.txnType)

	r.SetParam("asset", *s.asset)
	r.SetParam("startTime", *s.startTime)
	r.SetParam("endTime", *s.endTime)
	r.SetParam("current", *s.current)
	r.SetParam("size", *s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*StakingHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// StakingHistoryResponse define get staking history response
type StakingHistoryResponse struct {
	PositionId  string `json:"positionId"`
	Time        uint64 `json:"time"`
	Asset       string `json:"asset"`
	Project     string `json:"project"`
	Amount      string `json:"amount"`
	LockPeriod  string `json:"lockPeriod"`
	DeliverDate string `json:"deliverDate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}
