package spot

type OrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumulativeQuoteQty      string `json:"cumulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyId              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
	TransactTime            uint64 `json:"transactTime,omitempty"`
}
