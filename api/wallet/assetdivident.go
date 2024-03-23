package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Asset Dividend Record (USER_DATA)

// AssetDividendRecordService asset dividend record
type AssetDividendRecordService struct {
	C         *connector.Connector
	asset     *string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// Asset set asset
func (s *AssetDividendRecordService) Asset(asset string) *AssetDividendRecordService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *AssetDividendRecordService) StartTime(startTime uint64) *AssetDividendRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AssetDividendRecordService) EndTime(endTime uint64) *AssetDividendRecordService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AssetDividendRecordService) Limit(limit int) *AssetDividendRecordService {
	s.limit = &limit
	return s
}

func (s *AssetDividendRecordService) Do(ctx context.Context) (res *AssetDividendRecordResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/asset/assetDividend",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	r.SetParam("asset", s.asset)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDividendRecordResponse)
	err = json.Unmarshal(data, res)
	return
}

// AssetDividendRecordResponse define response of AssetDividendRecordService
type AssetDividendRecordResponse struct {
	Rows []struct {
		Id      int64  `json:"id"`
		Amount  string `json:"amount"`
		Asset   string `json:"asset"`
		DivTime uint64 `json:"divTime"`
		EnInfo  string `json:"enInfo"`
		TranId  int64  `json:"tranId"`
	} `json:"rows"`
	Total int64 `json:"total"`
}
