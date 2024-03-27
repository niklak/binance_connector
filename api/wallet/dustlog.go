package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// DustLog(USER_DATA)

// DustLogService dust log
//
//gen:new_service
type DustLogService struct {
	C           *connector.Connector
	accountType *string
	startTime   *uint64
	endTime     *uint64
}

// AccountType set accountType
func (s *DustLogService) AccountType(accountType string) *DustLogService {
	s.accountType = &accountType
	return s
}

// StartTime set startTime
func (s *DustLogService) StartTime(startTime uint64) *DustLogService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *DustLogService) EndTime(endTime uint64) *DustLogService {
	s.endTime = &endTime
	return s
}

func (s *DustLogService) Do(ctx context.Context) (res *DustLogResponse, err error) {

	r := request.New("/sapi/v1/asset/dribblet",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("accountType", s.accountType)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DustLogResponse)
	err = json.Unmarshal(data, res)
	return
}

// DustLogResponse define response of DustLogService
type DustLogResponse struct {
	Total              int `json:"total"`
	UserAssetDribblets []struct {
		OperateTime              uint64 `json:"operateTime"`
		TotalTransferedAmount    string `json:"totalTransferedAmount"`
		TotalServiceChargeAmount string `json:"totalServiceChargeAmount"`
		TransId                  int64  `json:"transId"`
		UserAssetDribbletDetails []struct {
			TransId             int64  `json:"transId"`
			ServiceChargeAmount string `json:"serviceChargeAmount"`
			Amount              string `json:"amount"`
			OperateTime         uint64 `json:"operateTime"`
			TransferedAmount    string `json:"transferedAmount"`
			FromAsset           string `json:"fromAsset"`
		} `json:"userAssetDribbletDetails"`
	} `json:"userAssetDribblets"`
}
