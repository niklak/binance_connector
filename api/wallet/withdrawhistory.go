package wallet

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Withdraw History (supporting network) (USER_DATA)

// WithdrawHistoryService withdraw history
//
//gen:new_service
type WithdrawHistoryService struct {
	C               *connector.Connector
	coin            *string
	withdrawOrderId *string
	status          *int
	offset          *int
	limit           *int
	startTime       *uint64
	endTime         *uint64
}

// Coin set coin
func (s *WithdrawHistoryService) Coin(coin string) *WithdrawHistoryService {
	s.coin = &coin
	return s
}

// WithdrawOrderId set withdrawOrderId
func (s *WithdrawHistoryService) WithdrawOrderId(withdrawOrderId string) *WithdrawHistoryService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Status set status
func (s *WithdrawHistoryService) Status(status int) *WithdrawHistoryService {
	s.status = &status
	return s
}

// Offset set offset
func (s *WithdrawHistoryService) Offset(offset int) *WithdrawHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *WithdrawHistoryService) Limit(limit int) *WithdrawHistoryService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *WithdrawHistoryService) StartTime(startTime uint64) *WithdrawHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *WithdrawHistoryService) EndTime(endTime uint64) *WithdrawHistoryService {
	s.endTime = &endTime
	return s
}

func (s *WithdrawHistoryService) Do(ctx context.Context) (res []*WithdrawHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/capital/withdraw/history",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("coin", s.coin)
	r.SetParam("withdrawOrderId", s.withdrawOrderId)
	r.SetParam("status", s.status)
	r.SetParam("offset", s.offset)
	r.SetParam("limit", s.limit)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*WithdrawHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// WithdrawHistoryResponse define response of WithdrawHistoryService
type WithdrawHistoryResponse struct {
	Id              string `json:"id"`
	Amount          string `json:"amount"`
	TransactionFee  string `json:"transactionFee"`
	Coin            string `json:"coin"`
	Status          int    `json:"status"`
	Address         string `json:"address"`
	TxId            string `json:"txId"`
	ApplyTime       string `json:"applyTime"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	WithdrawOrderId string `json:"withdrawOrderId"`
	Info            string `json:"info"`
	ConfirmNo       int    `json:"confirmNo"`
	WalletType      int    `json:"walletType"`
	TxKey           string `json:"txKey"`
}
