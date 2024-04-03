package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Subscribe ETH Staking V2(TRADE)  (POST /sapi/v2/eth-staking/eth/stake)
//
//gen:new_service
type SubscribeETHStakingV2Service struct {
	C      *connector.Connector
	amount float64
}

func (s *SubscribeETHStakingV2Service) Amount(amount float64) *SubscribeETHStakingV2Service {
	s.amount = amount
	return s
}

func (s *SubscribeETHStakingV2Service) Do(ctx context.Context, opts ...request.RequestOption) (res *SubscribeETHStakingV2Response, err error) {
	r := request.New(
		"/sapi/v2/eth-staking/eth/stake",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(SubscribeETHStakingV2Response)
	err = json.Unmarshal(data, res)

	return
}

type SubscribeETHStakingV2Response struct {
	Success         bool   `json:"success"`
	WBethAmount     string `json:"weth_amount"`
	ConversionRatio string `json:"conversion_ratio"`
}
