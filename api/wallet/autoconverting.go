package wallet

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Query auto-converting stable coins (USER_DATA)

// AutoConvertStableCoinService auto convert stable coin
//
//gen:new_service
type AutoConvertStableCoinService struct {
	C *connector.Connector
}

func (s *AutoConvertStableCoinService) Do(ctx context.Context) (res *AutoConvertStableCoinResponse, err error) {

	r := request.New(
		"/sapi/v1/capital/contract/convertible-coins",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AutoConvertStableCoinResponse)
	err = json.Unmarshal(data, res)
	return
}

// AutoConvertStableCoinResponse define response of AutoConvertStableCoinService
type AutoConvertStableCoinResponse struct {
	ConvertEnabled bool `json:"convertEnabled"`
	Coins          []struct {
		Asset string `json:"coin"`
	} `json:"coins"`
	ExchangeRates []struct {
		Asset string `json:"coin"`
	} `json:"exchangeRates"`
}
