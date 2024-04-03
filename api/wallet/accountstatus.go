package wallet

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Account Status (USER_DATA)

// AccountStatusService account status
//
//gen:new_service
type AccountStatusService struct {
	C *connector.Connector
}

func (s *AccountStatusService) Do(ctx context.Context) (res *AccountStatusResponse, err error) {

	r := request.New(
		"/sapi/v1/account/status",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AccountStatusResponse)
	err = json.Unmarshal(data, res)
	return
}

// AccountStatusResponse define response of AccountStatusService
type AccountStatusResponse struct {
	Data string `json:"data"`
}
