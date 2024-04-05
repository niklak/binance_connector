package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Funding Wallet (USER_DATA)

// FundingWalletService funding wallet
//
//gen:new_service
type FundingWalletService struct {
	C                *connector.Connector
	asset            *string
	needBtcValuation *string
}

// Asset sets asset
func (s *FundingWalletService) Asset(asset string) *FundingWalletService {
	s.asset = &asset
	return s
}

// NeedBtcValuation sets needBtcValuation
func (s *FundingWalletService) NeedBtcValuation(needBtcValuation string) *FundingWalletService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *FundingWalletService) Do(ctx context.Context) (res []*FundingWalletResponse, err error) {

	r := request.New(
		"/sapi/v1/asset/get-funding-asset",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("needBtcValuation", s.needBtcValuation)
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*FundingWalletResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// FundingWalletResponse define response of FundingWalletService
type FundingWalletResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}
