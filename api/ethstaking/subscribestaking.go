package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Subscribe ETH Staking(TRADE)  (POST /sapi/v1/eth-staking/eth/stake)
//
//gen:new_service
type SubscribeETHStakingService struct {
	C      *connector.Connector
	amount float64
}

func (s *SubscribeETHStakingService) Amount(amount float64) *SubscribeETHStakingService {
	s.amount = amount
	return s
}

func (s *SubscribeETHStakingService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubscribeETHStakingResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/stake",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(SubscribeETHStakingResponse)
	err = json.Unmarshal(data, res)

	return
}

type SubscribeETHStakingResponse struct {
	Success bool `json:"success"`
}
