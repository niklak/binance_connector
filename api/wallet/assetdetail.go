package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Asset Detail (USER_DATA)

// AssetDetailV2Service asset detail v2
//
//gen:new_service
type AssetDetailV2Service struct {
	C     *connector.Connector
	asset *string
}

// Asset set asset
func (s *AssetDetailV2Service) Asset(asset string) *AssetDetailV2Service {
	s.asset = &asset
	return s
}

func (s *AssetDetailV2Service) Do(ctx context.Context) (res *AssetDetailV2Response, err error) {

	r := request.New("/sapi/v1/asset/assetDetail",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetailV2Response)
	err = json.Unmarshal(data, res)
	return
}

// AssetDetailV2Response define response of AssetDetailV2Service
type AssetDetailV2Response struct {
	AssetDetail struct {
		MinWithdrawAmount string `json:"minWithdrawAmount"`
		DepositStatus     bool   `json:"depositStatus"`
		WithdrawFee       string `json:"withdrawFee"`
		WithdrawStatus    bool   `json:"withdrawStatus"`
		DepositTip        string `json:"depositTip"`
	} `json:"assetDetail"`
}
