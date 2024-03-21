package wallet

import (
	"context"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Disable Fast Withdraw Switch (USER_DATA)

// DisableFastWithdrawSwitchService disable fast withdraw switch
type DisableFastWithdrawSwitchService struct {
	C *connector.Connector
}

func (s *DisableFastWithdrawSwitchService) Do(ctx context.Context) (res *DisableFastWithdrawSwitchResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: "/sapi/v1/account/disableFastWithdrawSwitch",
		SecType:  request.SecTypeSigned,
	}
	r.Init()
	if _, err = s.C.CallAPI(ctx, r); err != nil {
		return
	}
	res = new(DisableFastWithdrawSwitchResponse)
	return
}

// DisableFastWithdrawSwitchResponse define response of DisableFastWithdrawSwitchService
// This endpoint has empty response
type DisableFastWithdrawSwitchResponse struct {
}
