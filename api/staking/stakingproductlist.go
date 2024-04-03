package staking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get Staking Product List (USER_DATA)
//
//gen:new_service
type StakingProductListService struct {
	C       *connector.Connector
	product string
	asset   *string
	current *int64
	size    *int64
}

// Product set product
func (s *StakingProductListService) Product(product string) *StakingProductListService {
	s.product = product
	return s
}

// Asset set asset
func (s *StakingProductListService) Asset(asset string) *StakingProductListService {
	s.asset = &asset
	return s
}

// Current set current
func (s *StakingProductListService) Current(current int64) *StakingProductListService {
	s.current = &current
	return s
}

// Size set size
func (s *StakingProductListService) Size(size int64) *StakingProductListService {
	s.size = &size
	return s
}

// Do send request
func (s *StakingProductListService) Do(ctx context.Context, opts ...request.RequestOption) (res []*StakingProductListResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/productList",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("product"),
	)

	r.SetParam("product", s.product)

	r.SetParam("asset", s.asset)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &res)
	return
}

// StakingProductListResponse define staking product list response
type StakingProductListResponse struct {
	ProjectId string `json:"projectId"`
	Detail    struct {
		Asset       string `json:"asset"`
		RewardAsset string `json:"rewardAsset"`
		Duration    int64  `json:"duration"`
		Renewable   bool   `json:"renewable"`
		Apy         string `json:"apy"`
	} `json:"detail"`
	Quota struct {
		TotalPersonalQuota string `json:"totalPersonalQuota"`
		Minimum            string `json:"minimum"`
	} `json:"quota"`
}
