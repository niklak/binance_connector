package wallet

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Dust Transfer (USER_DATA)

// DustTransferService dust transfer
//
//gen:new_service
type DustTransferService struct {
	C           *connector.Connector
	accountType *string
	asset       []string
}

// AccountType set accountType
func (s *DustTransferService) AccountType(accountType string) *DustTransferService {
	s.accountType = &accountType
	return s
}

// Asset set asset
func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.asset = asset
	return s
}

func (s *DustTransferService) Do(ctx context.Context) (res *DustTransferResponse, err error) {

	r := request.New(
		"/sapi/v1/asset/dust",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset"),
	)

	r.SetParam("accountType", s.accountType)

	asset := strings.Join(s.asset, ",")
	//TODO: check if this is correct -> asset=BTC,ETH
	r.SetParam("asset", asset)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DustTransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// DustTransferResponse define response of DustTransferService
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}
