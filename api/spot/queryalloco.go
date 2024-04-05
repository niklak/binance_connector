package spot

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Binance Query all OCO (USER_DATA) (GET /api/v3/allOrderList)
// QueryAllOCOService query all OCO order
//
//gen:new_service
type QueryAllOCOService struct {
	C         *connector.Connector
	fromId    *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// FromId set fromId
func (s *QueryAllOCOService) FromId(fromId int64) *QueryAllOCOService {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *QueryAllOCOService) StartTime(startTime uint64) *QueryAllOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *QueryAllOCOService) EndTime(endTime uint64) *QueryAllOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *QueryAllOCOService) Limit(limit int) *QueryAllOCOService {
	s.limit = &limit
	return s
}

// Do send request
func (s *QueryAllOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res []*OCOResponse, err error) {

	r := request.New(
		"/api/v3/allOrderList",
		request.SecType(request.SecTypeSigned),
		request.SetParam("fromId", s.fromId),
		request.SetParam("startTime", s.startTime),
		request.SetParam("endTime", s.endTime),
		request.SetParam("limit", s.limit),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*OCOResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}
