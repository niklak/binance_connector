package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Purchase Staking Product(USER_DATA)
//
//gen:new_service
type PurchaseStakingProductService struct {
	C         *connector.Connector
	product   string
	productId string
	amount    float64
	renewable *string
}

// Product set product
func (s *PurchaseStakingProductService) Product(product string) *PurchaseStakingProductService {
	s.product = product
	return s
}

// ProductId set productId
func (s *PurchaseStakingProductService) ProductId(productId string) *PurchaseStakingProductService {
	s.productId = productId
	return s
}

// Amount set amount
func (s *PurchaseStakingProductService) Amount(amount float64) *PurchaseStakingProductService {
	s.amount = amount
	return s
}

// Renewable set renewable
func (s *PurchaseStakingProductService) Renewable(renewable string) *PurchaseStakingProductService {
	s.renewable = &renewable
	return s
}

// Do send request
func (s *PurchaseStakingProductService) Do(ctx context.Context, opts ...request.RequestOption) (res *PurchaseStakingProductResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/purchase",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.product == "" {
		err = fmt.Errorf("%w: product", apierrors.ErrMissingParameter)
		return
	}
	if s.productId == "" {
		err = fmt.Errorf("%w: productId", apierrors.ErrMissingParameter)
		return
	}
	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("product", s.product)
	r.SetParam("productId", s.productId)
	r.SetParam("amount", s.amount)

	r.SetParam("renewable", s.renewable)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PurchaseStakingProductResponse)
	err = json.Unmarshal(data, res)
	return
}

// PurchaseStakingProductResponse define purchase staking product response
type PurchaseStakingProductResponse struct {
	PositionId string `json:"positionId"`
	Success    bool   `json:"success"`
}
