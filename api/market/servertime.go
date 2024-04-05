package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Binance Check Server Time endpoint (GET /api/v3/time)
//
//gen:new_service
type ServerTime struct {
	C *connector.Connector
}

// Send the request
func (s *ServerTime) Do(ctx context.Context, opts ...request.RequestOption) (res *ServerTimeResponse, err error) {

	r := request.New("/api/v3/time")

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ServerTimeResponse)
	err = json.Unmarshal(data, res)
	return
}

// ServerTimeResponse define server time response
type ServerTimeResponse struct {
	ServerTime uint64 `json:"serverTime"`
}
