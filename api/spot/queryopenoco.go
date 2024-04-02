package spot

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Query open OCO (USER_DATA) (GET /api/v3/openOrderList)
// QueryOpenOCOService query open OCO order
//
//gen:new_service
type QueryOpenOCOService struct {
	C *connector.Connector
}

// Do send request
func (s *QueryOpenOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OCOResponse, err error) {

	r := request.New(
		"/api/v3/openOrderList",
		request.SecType(request.SecTypeSigned),
	)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	return
}
