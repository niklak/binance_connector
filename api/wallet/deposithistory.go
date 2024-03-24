package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Deposit History (supporting network) (USER_DATA)

// DepositHistoryService deposit history
type DepositHistoryService struct {
	C         *connector.Connector
	coin      *string
	status    *int
	startTime *uint64
	endTime   *uint64
	offset    *int
	limit     *int
	txid      *string
}

// Coin set coin
func (s *DepositHistoryService) Coin(coin string) *DepositHistoryService {
	s.coin = &coin
	return s
}

// Status set status
func (s *DepositHistoryService) Status(status int) *DepositHistoryService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *DepositHistoryService) StartTime(startTime uint64) *DepositHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *DepositHistoryService) EndTime(endTime uint64) *DepositHistoryService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *DepositHistoryService) Offset(offset int) *DepositHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *DepositHistoryService) Limit(limit int) *DepositHistoryService {
	s.limit = &limit
	return s
}

// TxId set txid
func (s *DepositHistoryService) TxId(txid string) *DepositHistoryService {
	s.txid = &txid
	return s
}

func (s *DepositHistoryService) Do(ctx context.Context) (res []*DepositHistoryResponse, err error) {

	r := request.New("/sapi/v1/capital/deposit/hisrec",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("coin", s.coin)
	r.SetParam("status", s.status)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("offset", s.offset)
	r.SetParam("limit", s.limit)
	r.SetParam("txId", s.txid)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*DepositHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// DepositHistoryResponse define response of DepositHistoryService
type DepositHistoryResponse struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    uint64 `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}
