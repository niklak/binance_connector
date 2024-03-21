package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// System Status (System)

// SystemStatusService get account info
type SystemStatusService struct {
	C *connector.Connector
}

func (s *SystemStatusService) Do(ctx context.Context, opts ...request.RequestOption) (res []*SystemStatusResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/system/status",
		SecType:  request.SecTypeNone,
	}
	r.Init()
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SystemStatusResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// SystemStatusResponse define response of GetSystemStatusService
type SystemStatusResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}
