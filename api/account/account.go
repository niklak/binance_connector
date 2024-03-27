package account

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Account Information (USER_DATA) (GET /api/v3/account)
// AccountService get account information
//
//gen:new_service
type AccountService struct {
	C *connector.Connector
}

// Do send request
func (s *AccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *AccountResponse, err error) {
	r := newAccountRequest("/api/v3/account")
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(AccountResponse)
	err = json.Unmarshal(data, res)
	return
}

// Create AccountResponse
type AccountResponse struct {
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       uint64    `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}

// Balance define user balance of your account
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
