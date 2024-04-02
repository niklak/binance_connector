package account

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Account Commission (USER_DATA) (GET /api/v3/account/commission)
// AccountCommissionService gets current account commission rates.
//
//gen:new_service
type AccountCommissionService struct {
	C      *connector.Connector
	symbol string
}

// Symbol sets the symbol parameter (required)
func (s *AccountCommissionService) Symbol(symbol string) *AccountCommissionService {
	s.symbol = symbol
	return s
}

// Do executes the request
func (s *AccountCommissionService) Do(ctx context.Context, opts ...request.RequestOption) (res *AccountCommissionResponse, err error) {

	r := request.New("/api/v3/account/commission", request.SecType(request.SecTypeSigned))

	if s.symbol == "" {
		err = fmt.Errorf("%w: symbol", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(AccountCommissionResponse)
	err = json.Unmarshal(data, res)
	return
}

type Commission struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}
type CommissionDiscount struct {
	EnabledForAccount bool   `json:"enabledForAccount"`
	EnabledForSymbol  bool   `json:"enabledForSymbol"`
	DiscountAsset     string `json:"discountAsset"`
	Discount          string `json:"discount"`
}
type AccountCommissionResponse struct {
	Symbol             string             `json:"symbol"`
	StandardCommission Commission         `json:"standardCommission"`
	TaxCommission      Commission         `json:"taxCommission"`
	Discount           CommissionDiscount `json:"discount"`
}
