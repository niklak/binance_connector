package ethstaking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Get current ETH staking quota (USER_DATA) (GET /sapi/v1/eth-staking/eth/quota)
//
//gen:new_service
type CurrentETHStakingQuotaService struct {
	C *connector.Connector
}

func (s *CurrentETHStakingQuotaService) Do(ctx context.Context, opts ...request.RequestOption) (res *CurrentETHStakingQuotaResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/eth/quota",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(CurrentETHStakingQuotaResponse)
	err = json.Unmarshal(data, res)
	return
}

type CurrentETHStakingQuotaResponse struct {
	LeftStakingPersonalQuota    string `json:"leftStakingPersonalQuota"`
	LeftRedemptionPersonalQuota string `json:"leftStakingPoolQuota"`
}
