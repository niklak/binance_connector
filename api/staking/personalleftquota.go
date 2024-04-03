package staking

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get Personal Left Quota of Staking Product(USER_DATA)
//
//gen:new_service
type PersonalLeftQuotaService struct {
	C         *connector.Connector
	product   string
	productId string
}

// Product set product
func (s *PersonalLeftQuotaService) Product(product string) *PersonalLeftQuotaService {
	s.product = product
	return s
}

// ProductId set productId
func (s *PersonalLeftQuotaService) ProductId(productId string) *PersonalLeftQuotaService {
	s.productId = productId
	return s
}

// Do send request
func (s *PersonalLeftQuotaService) Do(ctx context.Context, opts ...request.RequestOption) (res []*PersonalLeftQuotaResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/personalLeftQuota",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("product", "productId"),
	)

	r.SetParam("product", s.product)
	r.SetParam("productId", s.productId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*PersonalLeftQuotaResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// PersonalLeftQuotaResponse define get staking asset response
type PersonalLeftQuotaResponse struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}
