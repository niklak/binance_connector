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

	r := request.New("/sapi/v1/account/disableFastWithdrawSwitch",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

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
