package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Deposit Address (supporting network) (USER_DATA)

// DepositAddressService deposit address
type DepositAddressService struct {
	C       *connector.Connector
	coin    string
	amount  *string
	network *string
}

// Coin set coin
func (s *DepositAddressService) Coin(coin string) *DepositAddressService {
	s.coin = coin
	return s
}

// Amount set amount
func (s *DepositAddressService) Amount(amount string) *DepositAddressService {
	s.amount = &amount
	return s
}

// Network set network
func (s *DepositAddressService) Network(network string) *DepositAddressService {
	s.network = &network
	return s
}

func (s *DepositAddressService) Do(ctx context.Context) (res *DepositAddressResponse, err error) {

	r := request.New("/sapi/v1/capital/deposit/address",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.coin == "" {
		err = fmt.Errorf("%w: coin", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("coin", s.coin)
	r.SetParam("amount", s.amount)
	r.SetParam("network", s.network)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DepositAddressResponse)
	err = json.Unmarshal(data, res)
	return
}

// DepositAddressResponse define response of DepositAddressService
type DepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
