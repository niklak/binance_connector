package wallet

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Query User Universal Transfer History (USER_DATA)

// UserUniversalTransferHistoryService user universal transfer history
//
//gen:new_service
type UserUniversalTransferHistoryService struct {
	C            *connector.Connector
	transferType string
	startTime    *uint64
	endTime      *uint64
	current      *int
	size         *int
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferHistoryService) TransferType(transferType string) *UserUniversalTransferHistoryService {
	s.transferType = transferType
	return s
}

// StartTime set startTime
func (s *UserUniversalTransferHistoryService) StartTime(startTime uint64) *UserUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UserUniversalTransferHistoryService) EndTime(endTime uint64) *UserUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *UserUniversalTransferHistoryService) Current(current int) *UserUniversalTransferHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *UserUniversalTransferHistoryService) Size(size int) *UserUniversalTransferHistoryService {
	s.size = &size
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferHistoryService) FromSymbol(fromSymbol string) *UserUniversalTransferHistoryService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferHistoryService) ToSymbol(toSymbol string) *UserUniversalTransferHistoryService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferHistoryService) Do(ctx context.Context) (res *UserUniversalTransferHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/asset/transfer",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("type"),
	)

	r.SetParam("type", s.transferType)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)
	r.SetParam("fromSymbol", s.fromSymbol)
	r.SetParam("toSymbol", s.toSymbol)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// UserUniversalTransferHistoryResponse define response of UserUniversalTransferHistoryService
type UserUniversalTransferHistoryResponse struct {
	Total int64 `json:"total"`
	Rows  []struct {
		Asset     string `json:"asset"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		TranId    int64  `json:"tranId"`
		Timestamp uint64 `json:"timestamp"`
	} `json:"rows"`
}
