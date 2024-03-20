package account

import (
	"net/http"

	"github.com/niklak/binance_connector/internal/request"
)

func newAccountRequest(endpoint string) *request.Request {
	r := request.Request{Method: http.MethodGet, SecType: request.SecTypeSigned}
	r.Endpoint = endpoint
	return r.Init()
}
