package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get Assets That Can Be Converted Into BNB (USER_DATA)

// AssetDetailService asset detail
//
//gen:new_service
type AssetDetailService struct {
	C           *connector.Connector
	accountType *string
}

// AccountType set accountType
func (s *AssetDetailService) AccountType(accountType string) *AssetDetailService {
	s.accountType = &accountType
	return s
}

func (s *AssetDetailService) Do(ctx context.Context) (res *AssetDetailResponse, err error) {

	r := request.New("/sapi/v1/asset/dust-btc",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("accountType", s.accountType)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetailResponse)
	err = json.Unmarshal(data, res)
	return
}

// AssetDetailResponse define response of AssetDetailService
type AssetDetailResponse struct {
	Details []struct {
		Asset            string `json:"asset"`
		AssetFullName    string `json:"assetFullName"`
		AmountFree       string `json:"amountFree"`
		ToBTC            string `json:"toBTC"`
		ToBNB            string `json:"toBNB"`
		ToBNBOffExchange string `json:"toBNBOffExchange"`
		Exchange         string `json:"exchange"`
	} `json:"details"`
	TotalTransferBtc   string `json:"totalTransferBtc"`
	TotalTransferBnb   string `json:"totalTransferBnb"`
	DribbletPercentage string `json:"dribbletPercentage"`
}
