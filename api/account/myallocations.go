package account

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

// Query Allocations (USER_DATA) (GET /api/v3/myAllocations)
// AccountAllocationsService retrieves allocations resulting from SOR order placement.
//
//gen:new_service
type AccountAllocationsService struct {
	C                *connector.Connector
	symbol           string
	startTime        *uint64
	endTime          *uint64
	fromAllocationId *int
	limit            *int
	orderId          *int64
}

// Symbol sets the symbol parameter (required)
func (s *AccountAllocationsService) Symbol(symbol string) *AccountAllocationsService {
	s.symbol = symbol
	return s
}

// StartTime sets the startTime parameter (optional)
func (s *AccountAllocationsService) StartTime(startTime uint64) *AccountAllocationsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter (optional)
func (s *AccountAllocationsService) EndTime(endTime uint64) *AccountAllocationsService {
	s.endTime = &endTime
	return s
}

// FromAllocationId sets the fromAllocationId parameter (optional)
func (s *AccountAllocationsService) FromAllocationId(fromAllocationId int) *AccountAllocationsService {
	s.fromAllocationId = &fromAllocationId
	return s
}

// Limit sets the limit parameter (optional)
func (s *AccountAllocationsService) Limit(limit int) *AccountAllocationsService {
	s.limit = &limit
	return s
}

// OrderId sets the orderId parameter (optional)
func (s *AccountAllocationsService) OrderId(orderId int64) *AccountAllocationsService {
	s.orderId = &orderId
	return s
}

// Do executes the request
func (s *AccountAllocationsService) Do(ctx context.Context, opts ...request.RequestOption) (res []*AccountAllocationsResponse, err error) {

	r := request.New(
		"/api/v3/myAllocations",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.SetParam("symbol", s.symbol),
		request.SetParam("startTime", s.startTime),
		request.SetParam("endTime", s.endTime),
		request.SetParam("fromAllocationId", s.fromAllocationId),
		request.SetParam("limit", s.limit),
		request.SetParam("orderId", s.orderId),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*AccountAllocationsResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

type AccountAllocationsResponse struct {
	Symbol          string `json:"symbol"`
	AllocationId    int    `json:"allocationId"`
	AllocationType  string `json:"allocationType"`
	OrderId         int    `json:"orderId"`
	OrderListId     int    `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsAllocator     bool   `json:"isAllocator"`
}
