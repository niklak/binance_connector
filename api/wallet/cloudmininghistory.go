package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get Cloud-Mining payment and refund history (USER_DATA)

// CloudMiningPaymentHistoryService cloud mining payment history
type CloudMiningPaymentHistoryService struct {
	C            *connector.Connector
	tranid       *int64
	clientTranId *string
	asset        *string
	startTime    uint64
	endTime      uint64
	current      *int
	size         *int
}

// Tranid set tranid
func (s *CloudMiningPaymentHistoryService) Tranid(tranid int64) *CloudMiningPaymentHistoryService {
	s.tranid = &tranid
	return s
}

// ClientTranId set clientTranId
func (s *CloudMiningPaymentHistoryService) ClientTranId(clientTranId string) *CloudMiningPaymentHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// Asset set asset
func (s *CloudMiningPaymentHistoryService) Asset(asset string) *CloudMiningPaymentHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *CloudMiningPaymentHistoryService) StartTime(startTime uint64) *CloudMiningPaymentHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *CloudMiningPaymentHistoryService) EndTime(endTime uint64) *CloudMiningPaymentHistoryService {
	s.endTime = endTime
	return s
}

// Current set current
func (s *CloudMiningPaymentHistoryService) Current(current int) *CloudMiningPaymentHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *CloudMiningPaymentHistoryService) Size(size int) *CloudMiningPaymentHistoryService {
	s.size = &size
	return s
}

func (s *CloudMiningPaymentHistoryService) Do(ctx context.Context) (res *CloudMiningPaymentHistoryResponse, err error) {

	r := request.New("/sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.startTime == 0 {
		err = fmt.Errorf("%w: startTime", apierrors.ErrMissingParameter)
		return
	}
	if s.endTime == 0 {
		err = fmt.Errorf("%w: endTime", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	r.SetParam("tranId", *s.tranid)
	r.SetParam("clientTranId", *s.clientTranId)
	r.SetParam("asset", *s.asset)
	r.SetParam("current", *s.current)
	r.SetParam("size", *s.size)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(CloudMiningPaymentHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// CloudMiningPaymentHistoryResponse define response of CloudMiningPaymentHistoryService
type CloudMiningPaymentHistoryResponse struct {
	Total int32 `json:"total"`
	Rows  []struct {
		CreateTime uint64 `json:"createTime"`
		TranId     int64  `json:"tranId"`
		Type       int32  `json:"type"`
		Asset      string `json:"asset"`
		Amount     string `json:"amount"`
		Status     string `json:"status"`
	} `json:"rows"`
}
