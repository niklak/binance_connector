package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get WBETH wrap history (USER_DATA) (GET /sapi/v1/eth-staking/wbeth/history/wrapHistory)
//
//gen:new_service
type WBETHWrapHistoryService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
	current   *uint64
	size      *uint64
}

func (s *WBETHWrapHistoryService) StartTime(startTime uint64) *WBETHWrapHistoryService {
	s.startTime = &startTime
	return s
}

func (s *WBETHWrapHistoryService) EndTime(endTime uint64) *WBETHWrapHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WBETHWrapHistoryService) Current(current uint64) *WBETHWrapHistoryService {
	s.current = &current
	return s
}

func (s *WBETHWrapHistoryService) Size(size uint64) *WBETHWrapHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *WBETHWrapHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *WBETHHistoryResponse, err error) {
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

type WBETHHistoryRow struct {
	Time         int64  `json:"time"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	ExchangeRate string `json:"exchangeRate"`
	Status       string `json:"status"`
}

// WBETHWrapHistoryResponse define response
type WBETHHistoryResponse struct {
	Total uint64             `json:"total"`
	Rows  []*WBETHHistoryRow `json:"rows"`
}
