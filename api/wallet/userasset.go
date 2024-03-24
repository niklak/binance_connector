package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// User Asset (USER_DATA)

// UserAssetService user asset
type UserAssetService struct {
	C                *connector.Connector
	asset            *string
	needBtcValuation *bool
}

// Asset set asset
func (s *UserAssetService) Asset(asset string) *UserAssetService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *UserAssetService) NeedBtcValuation(needBtcValuation bool) *UserAssetService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *UserAssetService) Do(ctx context.Context) (res []*UserAssetResponse, err error) {

	r := request.New("/sapi/v3/asset/getUserAsset",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("needBtcValuation", s.needBtcValuation)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*UserAssetResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// UserAssetResponse define response of UserAssetService
type UserAssetResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}
