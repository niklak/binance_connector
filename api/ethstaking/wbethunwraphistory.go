package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get WBETH unwrap history (USER_DATA) (GET /sapi/v1/eth-staking/wbeth/history/unwrapHistory)
//
//gen:new_service
type WBETHUnwrapHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *WBETHUnwrapHistoryService) StartTime(startTime uint64) *WBETHUnwrapHistoryService {
	s.startTime = &startTime
	return s
}

func (s *WBETHUnwrapHistoryService) EndTime(endTime uint64) *WBETHUnwrapHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WBETHUnwrapHistoryService) Current(current uint64) *WBETHUnwrapHistoryService {
	s.current = &current
	return s
}

func (s *WBETHUnwrapHistoryService) Size(size uint64) *WBETHUnwrapHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *WBETHUnwrapHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *WBETHHistoryResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/wbeth/history/wrapHistory",
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
	res = new(WBETHHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}
