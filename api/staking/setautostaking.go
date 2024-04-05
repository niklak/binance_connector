package staking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Set Auto Staking(USER_DATA)
//
//gen:new_service
type SetAutoStakingService struct {
	C          *connector.Connector
	product    string
	positionId string
	renewable  string
}

// Product set product
func (s *SetAutoStakingService) Product(product string) *SetAutoStakingService {
	s.product = product
	return s
}

// PositionId set positionId
func (s *SetAutoStakingService) PositionId(positionId string) *SetAutoStakingService {
	s.positionId = positionId
	return s
}

// Renewable set renewable
func (s *SetAutoStakingService) Renewable(renewable string) *SetAutoStakingService {
	s.renewable = renewable
	return s
}

// Do send request
func (s *SetAutoStakingService) Do(ctx context.Context, opts ...request.RequestOption) (res *SetAutoStakingResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/setAutoStaking",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("product", "positionId", "renewable"),
	)

	r.SetParam("product", s.product)
	r.SetParam("positionId", s.positionId)
	r.SetParam("renewable", s.renewable)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SetAutoStakingResponse)
	err = json.Unmarshal(data, res)
	return
}

// SetAutoStakingResponse define set auto staking response
type SetAutoStakingResponse struct {
	Success bool `json:"success"`
}
