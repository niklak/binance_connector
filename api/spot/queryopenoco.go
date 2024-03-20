package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Query open OCO (USER_DATA) (GET /api/v3/openOrderList)
// QueryOpenOCOService query open OCO order
type QueryOpenOCOService struct {
	C *connector.Connector
}

// Do send request
func (s *QueryOpenOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OCOResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/openOrderList",
		SecType:  request.SecTypeSigned,
	}
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	return
}
