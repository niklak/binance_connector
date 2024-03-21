package wallet

import (
	"context"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Enable Fast Withdraw Switch (USER_DATA)

// EnableFastWithdrawSwitchService enable fast withdraw switch
type EnableFastWithdrawSwitchService struct {
	C *connector.Connector
}

func (s *EnableFastWithdrawSwitchService) Do(ctx context.Context) (res *EnableFastWithdrawSwitchResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: "/sapi/v1/account/enableFastWithdrawSwitch",
		SecType:  request.SecTypeSigned,
	}
	r.Init()
	_, err = s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(EnableFastWithdrawSwitchResponse)
	return
}

// EnableFastWithdrawSwitchResponse define response of EnableFastWithdrawSwitchService
// This endpoint has empty response
type EnableFastWithdrawSwitchResponse struct {
}
