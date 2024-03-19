package market

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Exchange Information endpoint (GET /api/v3/exchangeInfo)
type ExchangeInfo struct {
	C *connector.Connector
}

// Send the request
func (s *ExchangeInfo) Do(ctx context.Context, opts ...request.RequestOption) (res *ExchangeInfoResponse, err error) {

	r := newMarketRequest("/api/v3/exchangeInfo")

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ExchangeInfoResponse)
	err = json.Unmarshal(data, res)
	return
}

// ExchangeInfoResponse define exchange info response
type ExchangeInfoResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      uint64            `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

// RateLimit define rate limit
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int    `json:"limit"`
}

// ExchangeFilter define exchange filter
type ExchangeFilter struct {
	FilterType string `json:"filterType"`
	MaxNumAlgo int64  `json:"maxNumAlgoOrders"`
}

// Symbol define symbol
type SymbolInfo struct {
	Symbol                     string          `json:"symbol"`
	Status                     string          `json:"status"`
	BaseAsset                  string          `json:"baseAsset"`
	BaseAssetPrecision         int64           `json:"baseAssetPrecision"`
	QuoteAsset                 string          `json:"quoteAsset"`
	QuotePrecision             int64           `json:"quotePrecision"`
	OrderTypes                 []string        `json:"orderTypes"`
	IcebergAllowed             bool            `json:"icebergAllowed"`
	OcoAllowed                 bool            `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool            `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool            `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool            `json:"isMarginTradingAllowed"`
	Filters                    []*SymbolFilter `json:"filters"`
	Permissions                []string        `json:"permissions"`
}

// SymbolFilter define symbol filter
type SymbolFilter struct {
	FilterType       string `json:"filterType"`
	MinPrice         string `json:"minPrice"`
	MaxPrice         string `json:"maxPrice"`
	TickSize         string `json:"tickSize"`
	MinQty           string `json:"minQty"`
	MaxQty           string `json:"maxQty"`
	StepSize         string `json:"stepSize"`
	MinNotional      string `json:"minNotional"`
	Limit            uint   `json:"limit"`
	MaxNumAlgoOrders int64  `json:"maxNumAlgoOrders"`
}
