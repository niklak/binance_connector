package spot

import (
	"net/http"

	"github.com/niklak/binance_connector/internal/request"
)

func newSpotRequest(endpoint string) *request.Request {
	r := request.Request{Method: http.MethodPost, SecType: request.SecTypeSigned}
	r.Endpoint = endpoint
	return r.Init()
}
