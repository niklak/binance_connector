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

// Get Staking Product Position(USER_DATA)
//
//gen:new_service
type StakingProductPositionService struct {
	C         *connector.Connector
	product   string
	productId *string
	asset     *string
	current   *int64
	size      *int64
}

// Product set product
func (s *StakingProductPositionService) Product(product string) *StakingProductPositionService {
	s.product = product
	return s
}

// ProductId set productId
func (s *StakingProductPositionService) ProductId(productId string) *StakingProductPositionService {
	s.productId = &productId
	return s
}

// Asset set asset
func (s *StakingProductPositionService) Asset(asset string) *StakingProductPositionService {
	s.asset = &asset
	return s
}

// Current set current
func (s *StakingProductPositionService) Current(current int64) *StakingProductPositionService {
	s.current = &current
	return s
}

// Size set size
func (s *StakingProductPositionService) Size(size int64) *StakingProductPositionService {
	s.size = &size
	return s
}

// Do send request
func (s *StakingProductPositionService) Do(ctx context.Context, opts ...request.RequestOption) (res []*StakingProductPositionResponse, err error) {

	r := request.New(
		"/sapi/v1/staking/position",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.product == "" {
		err = fmt.Errorf("%w: product", apierrors.ErrMissingParameter)
		return

	}

	r.SetParam("product", s.product)

	r.SetParam("productId", s.productId)
	r.SetParam("asset", s.asset)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*StakingProductPositionResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// StakingProductPositionResponse define get staking product position response
type StakingProductPositionResponse struct {
	PositionId            string `json:"positionId"`
	ProjectId             string `json:"projectId"`
	Asset                 string `json:"asset"`
	Amount                string `json:"amount"`
	PurchaseTime          string `json:"purchaseTime"`
	Duration              string `json:"duration"`
	AccrualDays           string `json:"accrualDays"`
	RewardAsset           string `json:"rewardAsset"`
	APY                   string `json:"APY"`
	RewardAmt             string `json:"rewardAmt"`
	ExtraRewardAsset      string `json:"extraRewardAsset"`
	ExtraRewardAPY        string `json:"extraRewardAPY"`
	EstExtraRewardAmt     string `json:"estExtraRewardAmt"`
	NextInterestPay       string `json:"nextInterestPay"`
	NextInterestPayDate   string `json:"nextInterestPayDate"`
	PayInterestPeriod     string `json:"payInterestPeriod"`
	RedeemAmountEarly     string `json:"redeemAmountEarly"`
	InterestEndDate       string `json:"interestEndDate"`
	DeliverDate           string `json:"deliverDate"`
	RedeemPeriod          string `json:"redeemPeriod"`
	RedeemingAmt          string `json:"redeemingAmt"`
	PartialAmtDeliverDate string `json:"partialAmtDeliverDate"`
	CanRedeemEarly        bool   `json:"canRedeemEarly"`
	Renewable             bool   `json:"renewable"`
	Type                  string `json:"type"`
	Status                string `json:"status"`
}
