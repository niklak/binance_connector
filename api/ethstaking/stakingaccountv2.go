package ethstaking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// ETH Staking account V2(USER_DATA)
//
//gen:service
type ETHStakingAccountV2Service struct {
	C *connector.Connector
}

func (s *ETHStakingAccountV2Service) Do(ctx context.Context, opts ...request.RequestOption) (res *ETHStakingAccountV2Response, err error) {
	r := request.New(
		"/sapi/v2/eth-staking/eth/account",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ETHStakingAccountV2Response)
	err = json.Unmarshal(data, res)
	return
}

type ETHStakingAccountV2Response struct {
	HoldingInETH string `json:"holdingInETH"`
	Holdings     struct {
		WbethAmount string `json:"wbethAmount"`
		BethAmount  string `json:"bethAmount"`
	} `json:"holdings"`
	ThirtyDaysProfitInETH string `json:"thirtyDaysProfitInETH"`
	Profit                struct {
		AmountFromWBETH string `json:"amountFromWBETH"`
		AmountFromBETH  string `json:"amountFromBETH"`
	} `json:"profit"`
}
