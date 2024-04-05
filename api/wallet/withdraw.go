package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Withdraw(USER_DATA)

// WithdrawService withdraw
//
//gen:new_service
type WithdrawService struct {
	C                  *connector.Connector
	coin               string
	withdrawOrderId    *string
	network            *string
	address            string
	addressTag         *string
	amount             float64
	transactionFeeFlag *bool
	name               *string
	walletType         *int
}

// Coin set coin
func (s *WithdrawService) Coin(coin string) *WithdrawService {
	s.coin = coin
	return s
}

// WithdrawOrderId set withdrawOrderId
func (s *WithdrawService) WithdrawOrderId(withdrawOrderId string) *WithdrawService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Network set network
func (s *WithdrawService) Network(network string) *WithdrawService {
	s.network = &network
	return s
}

// Address set address
func (s *WithdrawService) Address(address string) *WithdrawService {
	s.address = address
	return s
}

// AddressTag set addressTag
func (s *WithdrawService) AddressTag(addressTag string) *WithdrawService {
	s.addressTag = &addressTag
	return s
}

// Amount set amount
func (s *WithdrawService) Amount(amount float64) *WithdrawService {
	s.amount = amount
	return s
}

// TransactionFeeFlag set transactionFeeFlag
func (s *WithdrawService) TransactionFeeFlag(transactionFeeFlag bool) *WithdrawService {
	s.transactionFeeFlag = &transactionFeeFlag
	return s
}

// Name set name
func (s *WithdrawService) Name(name string) *WithdrawService {
	s.name = &name
	return s
}

// WalletType set walletType
func (s *WithdrawService) WalletType(walletType int) *WithdrawService {
	s.walletType = &walletType
	return s
}

func (s *WithdrawService) Do(ctx context.Context) (res *WithdrawResponse, err error) {

	r := request.New(
		"/sapi/v1/capital/withdraw/apply",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("coin", "address", "amount"),
	)

	r.SetParam("coin", s.coin)
	r.SetParam("address", s.address)
	r.SetParam("amount", s.amount)

	r.SetParam("withdrawOrderId", s.withdrawOrderId)
	r.SetParam("network", s.network)
	r.SetParam("addressTag", s.addressTag)
	r.SetParam("transactionFeeFlag", s.transactionFeeFlag)
	r.SetParam("name", s.name)
	r.SetParam("walletType", s.walletType)
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(WithdrawResponse)
	err = json.Unmarshal(data, res)
	return
}

// WithdrawResponse define response of WithdrawService
type WithdrawResponse struct {
	Id string `json:"id"`
}
