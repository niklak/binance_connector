package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Redeem ETH (TRADE) (POST /sapi/v1/eth-staking/eth/redeem )
//
//gen:new_service
type RedeemETHService struct {
	C      *connector.Connector
	asset  *string
	amount float64
}

func (s *RedeemETHService) Asset(asset string) *RedeemETHService {
	s.asset = &asset
	return s
}

func (s *RedeemETHService) Amount(amount float64) *RedeemETHService {
	s.amount = amount
	return s
}

func (s *RedeemETHService) Do(ctx context.Context, opts ...request.RequestOption) (res *RedeemETHResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/redeem",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(RedeemETHResponse)
	err = json.Unmarshal(data, res)

	return
}

type RedeemETHResponse struct {
	Success         bool   `json:"success"`
	ArrivalTime     int64  `json:"arrival_time"`
	EthAmount       string `json:"eth_amount"`
	ConversionRatio string `json:"conversion_ratio"`
}
