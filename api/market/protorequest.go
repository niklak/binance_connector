package market

import (
	"net/http"

	"github.com/niklak/binance_connector/internal/request"
)

func newMarketRequest(endpoint string) *request.Request {
	r := request.Request{Method: http.MethodGet, SecType: request.SecTypeNone}
	r.Endpoint = endpoint
	return r.Init()
}
