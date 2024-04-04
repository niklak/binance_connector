package ethstaking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// ETH Staking account (USER_DATA) (GET /sapi/v1/eth-staking/account)
//
//gen:new_service
type ETHStakingAccountService struct {
	C *connector.Connector
}

func (s *ETHStakingAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *ETHStakingAccountResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/account",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ETHStakingAccountResponse)
	err = json.Unmarshal(data, res)
	return
}

type ETHStakingAccountResponse struct {
	CumulativeProfitInBETH string `json:"cumulativeProfitInBETH"`
	LastDayProfitInBETH    string `json:"lastDayProfitInBETH"`
}
