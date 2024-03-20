package account

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Query Current Order Count Usage (TRADE)
// QueryCurrentOrderCountUsageService query current order count usage
type QueryCurrentOrderCountUsageService struct {
	C *connector.Connector
}

// Do send request
func (s *QueryCurrentOrderCountUsageService) Do(ctx context.Context, opts ...request.RequestOption) (res []*QueryCurrentOrderCountUsageResponse, err error) {

	r := newAccountRequest("/api/v3/rateLimit/order")

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*QueryCurrentOrderCountUsageResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// Create QueryCurrentOrderCountUsageResponse
type QueryCurrentOrderCountUsageResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}
