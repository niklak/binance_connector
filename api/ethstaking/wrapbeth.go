package ethstaking

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Wrap BETH(TRADE) (POST /sapi/v1/eth-staking/wbeth/wrap)
//
//gen:new_service
type WrapBETHService struct {
	C      *connector.Connector
	amount float64
}

func (s *WrapBETHService) Amount(amount float64) *WrapBETHService {
	s.amount = amount
	return s
}

func (s *WrapBETHService) Do(ctx context.Context, opts ...request.RequestOption) (res *WrapBETHResponse, err error) {
	r := request.New(
		"/sapi/v1/eth-staking/wbeth/wrap",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("amount"),
	)

	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(WrapBETHResponse)
	err = json.Unmarshal(data, res)

	return
}

type WrapBETHResponse struct {
	Success      bool   `json:"success"`
	WbethAmount  string `json:"wbethAmount"`
	ExchangeRate string `json:"exchangeRate"`
}
