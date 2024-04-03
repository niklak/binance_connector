package staking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Redeem Staking Product(USER_DATA)
//
//gen:new_service
type RedeemStakingProductService struct {
	C          *connector.Connector
	product    string
	positionId *string
	productId  string
	amount     *float64
}

// Product set product
func (s *RedeemStakingProductService) Product(product string) *RedeemStakingProductService {
	s.product = product
	return s
}

// PositionId set positionId
func (s *RedeemStakingProductService) PositionId(positionId string) *RedeemStakingProductService {
	s.positionId = &positionId
	return s
}

// ProductId set productId
func (s *RedeemStakingProductService) ProductId(productId string) *RedeemStakingProductService {
	s.productId = productId
	return s
}

// Amount set amount
func (s *RedeemStakingProductService) Amount(amount float64) *RedeemStakingProductService {
	s.amount = &amount
	return s
}

// Do send request
func (s *RedeemStakingProductService) Do(ctx context.Context, opts ...request.RequestOption) (res *RedeemStakingProductResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/redeem",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("product", "productId"),
	)

	r.SetParam("product", s.product)
	r.SetParam("productId", s.productId)

	r.SetParam("positionId", s.positionId)
	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(RedeemStakingProductResponse)
	err = json.Unmarshal(data, res)
	return
}

// RedeemStakingProductResponse define redeem staking product response
type RedeemStakingProductResponse struct {
	Success bool `json:"success"`
}
