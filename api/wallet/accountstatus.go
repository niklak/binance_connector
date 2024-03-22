package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Account Status (USER_DATA)

// AccountStatusService account status
type AccountStatusService struct {
	C *connector.Connector
}

func (s *AccountStatusService) Do(ctx context.Context) (res *AccountStatusResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/account/status",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

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
