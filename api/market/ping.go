package market

import (
	"context"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Test Connectivity endpoint (GET /api/v3/ping)
//
//gen:new_service
type Ping struct {
	C *connector.Connector
}

// Send the request
func (s *Ping) Do(ctx context.Context, opts ...request.RequestOption) (err error) {
	r := request.New("/api/v3/ping")
	_, err = s.C.CallAPI(ctx, r, opts...)
	return
}
