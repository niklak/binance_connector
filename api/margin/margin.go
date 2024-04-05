package margin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Cross Margin Account Transfer API Endpoint

// TransferService transfer between spot and margin account
//
//gen:new_service
type TransferService struct {
	C            *connector.Connector
	asset        string
	amount       float64
	transferType int
}

// Asset set asset
func (s *TransferService) Asset(asset string) *TransferService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferService) Amount(amount float64) *TransferService {
	s.amount = amount
	return s
}

// TransferType set transfer type
func (s *TransferService) TransferType(transferType int) *TransferService {
	s.transferType = transferType
	return s
}

// Do send request
func (s *TransferService) Do(ctx context.Context, opts ...request.RequestOption) (res *TransferResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/transfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset", "amount", "type"),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("type", s.transferType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(TransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// TransferResponse define transfer response
type TransferResponse struct {
	TranId int64 `json:"tranId"`
}

// Cross Margin Account Borrow API Endpoint

// BorrowService borrow from cross margin account
//
//gen:new_service
type BorrowService struct {
	C          *connector.Connector
	asset      string
	amount     float64
	isIsolated *string
	symbol     *string
}

// Asset set asset
func (s *BorrowService) Asset(asset string) *BorrowService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *BorrowService) Amount(amount float64) *BorrowService {
	s.amount = amount
	return s
}

// IsIsolated set isolated
func (s *BorrowService) IsIsolated(isIsolated string) *BorrowService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *BorrowService) Do(ctx context.Context, opts ...request.RequestOption) (res *BorrowResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/loan",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset", "amount"),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)

	r.SetParam("isolatedSymbol", *s.isIsolated)
	r.SetParam("symbol", *s.symbol)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(BorrowResponse)
	err = json.Unmarshal(data, res)
	return
}

// BorrowResponse define borrow response
type BorrowResponse struct {
	TranId int64 `json:"tranId"`
}

// Cross Margin Account Repay API Endpoint

// RepayService repay to cross margin account
//
//gen:new_service
type RepayService struct {
	C          *connector.Connector
	asset      string
	isIsolated *string
	symbol     *string
	amount     float64
}

// Asset set asset
func (s *RepayService) Asset(asset string) *RepayService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *RepayService) Amount(amount float64) *RepayService {
	s.amount = amount
	return s
}

// Symbol set symbol
func (s *RepayService) Symbol(symbol string) *RepayService {
	s.symbol = &symbol
	return s
}

// IsIsolated set isolated
func (s *RepayService) IsIsolated(isIsolated string) *RepayService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *RepayService) Do(ctx context.Context, opts ...request.RequestOption) (res *RepayResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/repay",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset", "amount"),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)

	r.SetParam("isIsolated", *s.isIsolated)
	r.SetParam("symbol", *s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(RepayResponse)
	err = json.Unmarshal(data, res)
	return
}

// RepayResponse define repay response
type RepayResponse struct {
	TranId int64 `json:"tranId"`
}

// Query Margin Asset API Endpoint

// QueryMarginAssetService query margin asset
//
//gen:new_service
type QueryMarginAssetService struct {
	C     *connector.Connector
	asset string
}

// Asset set asset
func (s *QueryMarginAssetService) Asset(asset string) *QueryMarginAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *QueryMarginAssetService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryMarginAssetResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/asset",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeAPIKey),
		request.RequiredParams("asset"),
	)

	r.SetParam("asset", s.asset)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(QueryMarginAssetResponse)
	err = json.Unmarshal(data, res)
	return
}

// QueryMarginAssetResponse define query margin asset response
type QueryMarginAssetResponse struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// Query Cross Margin Pair API Endpoint

// QueryCrossMarginPairService query cross margin pair
//
//gen:new_service
type QueryCrossMarginPairService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *QueryCrossMarginPairService) Symbol(symbol string) *QueryCrossMarginPairService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *QueryCrossMarginPairService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryCrossMarginPairResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/pair",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(QueryCrossMarginPairResponse)
	err = json.Unmarshal(data, res)

	return
}

// QueryCrossMarginPairResponse define query cross margin pair response
type QueryCrossMarginPairResponse struct {
	SymbolDetail struct {
		Symbol        string `json:"symbol"`
		IsMarginTrade bool   `json:"isMarginTrade"`
		IsBuyAllowed  bool   `json:"isBuyAllowed"`
		IsSellAllowed bool   `json:"isSellAllowed"`
	} `json:"symbolDetail"`
}

// Get all margin assets API Endpoint

// GetAllMarginAssetsService get all margin assets

//gen:new_service
type GetAllMarginAssetsService struct {
	C *connector.Connector
}

// Do send request
func (s *GetAllMarginAssetsService) Do(ctx context.Context, opts ...request.RequestOption) (res *GetAllMarginAssetsResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/allAssets",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(GetAllMarginAssetsResponse)
	err = json.Unmarshal(data, res)

	return
}

// GetAllMarginAssetsResponse define get all margin assets response
type GetAllMarginAssetsResponse struct {
	AssetDetailList []struct {
		AssetFullName  string `json:"assetFullName"`
		AssetName      string `json:"assetName"`
		IsBorrowable   bool   `json:"isBorrowable"`
		IsMortgageable bool   `json:"isMortgageable"`
		MinLoanAmt     string `json:"minLoanAmt"`
		MaxLoanAmt     string `json:"maxLoanAmt"`
		MinMortgageAmt string `json:"minMortgageAmt"`
		MaxMortgageAmt string `json:"maxMortgageAmt"`
		Asset          string `json:"asset"`
	} `json:"assetDetailList"`
}

// Get all margin pairs API Endpoint

// GetAllMarginPairsService get all margin pairs
//
//gen:new_service
type GetAllMarginPairsService struct {
	C *connector.Connector
}

// Do send request
func (s *GetAllMarginPairsService) Do(ctx context.Context, opts ...request.RequestOption) (res *GetAllMarginPairsResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/allPairs",
		request.SecType(request.SecTypeAPIKey),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(GetAllMarginPairsResponse)
	err = json.Unmarshal(data, res)
	return
}

// GetAllMarginPairsResponse define get all margin pairs response
type GetAllMarginPairsResponse struct {
	SymbolDetailList []struct {
		Base          string `json:"base"`
		Id            int    `json:"id"`
		IsBuyAllowed  bool   `json:"isBuyAllowed"`
		IsMarginTrade bool   `json:"isMarginTrade"`
		IsSellAllowed bool   `json:"isSellAllowed"`
		Quote         string `json:"quote"`
		Symbol        string `json:"symbol"`
	} `json:"symbolDetailList"`
}

// Query Margin Price Index API Endpoint

// QueryMarginPriceIndexService query margin price index
//
//gen:new_service
type QueryMarginPriceIndexService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *QueryMarginPriceIndexService) Symbol(symbol string) *QueryMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *QueryMarginPriceIndexService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryMarginPriceIndexResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/priceIndex",
		request.SecType(request.SecTypeAPIKey),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(QueryMarginPriceIndexResponse)
	err = json.Unmarshal(data, res)
	return
}

// QueryMarginPriceIndexResponse define query margin price index response
type QueryMarginPriceIndexResponse struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// Margin Account New Order (TRADE) API Endpoint

// MarginAccountNewOrderService margin account new order
//
//gen:new_service
type MarginAccountNewOrderService struct {
	C                *connector.Connector
	symbol           string
	isIsolated       *string
	side             string
	orderType        string
	quantity         *float64
	quoteOrderQty    *float64
	price            *float64
	stopPrice        *float64
	newClientOrderId *string
	icebergQty       *float64
	newOrderRespType *string
	sideEffectType   *string
	timeInForce      *string
}

// Symbol set symbol
func (s *MarginAccountNewOrderService) Symbol(symbol string) *MarginAccountNewOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOrderService) IsIsolated(isIsolated string) *MarginAccountNewOrderService {
	s.isIsolated = &isIsolated
	return s
}

// Side set side
func (s *MarginAccountNewOrderService) Side(side string) *MarginAccountNewOrderService {
	s.side = side
	return s
}

// OrderType set orderType
func (s *MarginAccountNewOrderService) OrderType(orderType string) *MarginAccountNewOrderService {
	s.orderType = orderType
	return s
}

// Quantity set quantity
func (s *MarginAccountNewOrderService) Quantity(quantity float64) *MarginAccountNewOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *MarginAccountNewOrderService) QuoteOrderQty(quoteOrderQty float64) *MarginAccountNewOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *MarginAccountNewOrderService) Price(price float64) *MarginAccountNewOrderService {
	s.price = &price
	return s
}

// StopPrice set stopPrice
func (s *MarginAccountNewOrderService) StopPrice(stopPrice float64) *MarginAccountNewOrderService {
	s.stopPrice = &stopPrice
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountNewOrderService) NewClientOrderId(newClientOrderId string) *MarginAccountNewOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// IcebergQty set icebergQty
func (s *MarginAccountNewOrderService) IcebergQty(icebergQty float64) *MarginAccountNewOrderService {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOrderService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOrderService) SideEffectType(sideEffectType string) *MarginAccountNewOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

// TimeInForce set timeInForce
func (s *MarginAccountNewOrderService) TimeInForce(timeInForce string) *MarginAccountNewOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Do send request
func (s *MarginAccountNewOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {

	r := request.New(
		"/sapi/v1/margin/order",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol", "side", "type"),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("side", s.side)
	r.SetParam("type", s.orderType)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("quantity", s.quantity)
	r.SetParam("price", s.price)
	r.SetParam("stopPrice", s.stopPrice)
	r.SetParam("newClientOrderId", s.newClientOrderId)
	r.SetParam("icebergQty", s.icebergQty)
	r.SetParam("newOrderRespType", s.newOrderRespType)
	r.SetParam("sideEffectType", s.sideEffectType)
	r.SetParam("timeInForce", s.timeInForce)
	r.SetParam("quoteOrderQty", s.quoteOrderQty)
	r.SetParam("sideEffectType", s.sideEffectType)
	r.SetParam("timeInForce", s.timeInForce)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}

	switch *s.newOrderRespType {
	case "ACK":
		res = new(MarginAccountNewOrderResponseACK)
	case "RESULT":
		res = new(MarginAccountNewOrderResponseRESULT)
	case "FULL":
		res = new(MarginAccountNewOrderResponseFULL)
	default:
		switch s.orderType {
		case "MARKET", "LIMIT":
			res = new(MarginAccountNewOrderResponseFULL)
		default:
			res = new(MarginAccountNewOrderResponseACK)
		}
	}
	err = json.Unmarshal(data, res)
	return
}

// Create MarginAccountNewOrderResponseACK
type MarginAccountNewOrderResponseACK struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId int64  `json:"clientOrderId"`
	IsIsolated    bool   `json:"isIsolated"`
	TransactTime  uint64 `json:"transactTime"`
}

// Create MarginAccountNewOrderResponseRESULT
type MarginAccountNewOrderResponseRESULT struct {
	Symbol             string `json:"symbol"`
	OrderId            int64  `json:"orderId"`
	ClientOrderId      string `json:"clientOrderId"`
	TransactTime       uint64 `json:"transactTime"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	IsIsolated         bool   `json:"isIsolated"`
	Side               string `json:"side"`
}

// Create MarginAccountNewOrderResponseFULL
type MarginAccountNewOrderResponseFULL struct {
	Symbol                string  `json:"symbol"`
	OrderId               int64   `json:"orderId"`
	ClientOrderId         string  `json:"clientOrderId"`
	TransactTime          uint64  `json:"transactTime"`
	Price                 string  `json:"price"`
	OrigQty               string  `json:"origQty"`
	ExecutedQty           string  `json:"executedQty"`
	CumulativeQuoteQty    string  `json:"cummulativeQuoteQty"`
	Status                string  `json:"status"`
	TimeInForce           string  `json:"timeInForce"`
	Type                  string  `json:"type"`
	Side                  string  `json:"side"`
	MarginBuyBorrowAmount float64 `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset"`
	IsIsolated            bool    `json:"isIsolated"`
	Fills                 []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
	} `json:"fills"`
}

// Margin Account Cancel Order (TRADE) API Endpoint

// MarginAccountCancelOrderService margin account cancel order
//
//gen:new_service
type MarginAccountCancelOrderService struct {
	C                 *connector.Connector
	symbol            string
	isIsolated        *string
	orderId           *int
	origClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *MarginAccountCancelOrderService) Symbol(symbol string) *MarginAccountCancelOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelOrderService) IsIsolated(isIsolated string) *MarginAccountCancelOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountCancelOrderService) OrderId(orderId int) *MarginAccountCancelOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountCancelOrderService) OrigClientOrderId(origClientOrderId string) *MarginAccountCancelOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountCancelOrderService) NewClientOrderId(newClientOrderId string) *MarginAccountCancelOrderService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *MarginAccountCancelOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountCancelOrderResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/order",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.RequiredOneOfParams([]string{"orderId", "origClientOrderId"}),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("orderId", s.orderId)
	r.SetParam("origClientOrderId", s.origClientOrderId)
	r.SetParam("newClientOrderId", s.newClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountCancelOrderResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountCancelOrderResponse define margin account cancel order response
type MarginAccountCancelOrderResponse struct {
	Symbol             string `json:"symbol"`
	IsIsolated         bool   `json:"isIsolated"`
	OrderId            int    `json:"orderId"`
	OrigClientOrderId  string `json:"origClientOrderId"`
	ClientOrderId      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
}

// Margin Account Cancel All Orders (TRADE) API Endpoint

// MarginAccountCancelAllOrdersService margin account cancel all orders
//
//gen:new_service
type MarginAccountCancelAllOrdersService struct {
	C          *connector.Connector
	symbol     string
	isIsolated *string
}

// Symbol set symbol
func (s *MarginAccountCancelAllOrdersService) Symbol(symbol string) *MarginAccountCancelAllOrdersService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelAllOrdersService) IsIsolated(isIsolated string) *MarginAccountCancelAllOrdersService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *MarginAccountCancelAllOrdersService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountCancelAllOrdersResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/openOrders",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("isIsolated", s.isIsolated)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountCancelAllOrdersResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountCancelAllOrdersResponse define margin account cancel all orders response
type MarginAccountCancelAllOrdersResponse struct {
	Symbol             string `json:"symbol"`
	IsIsolated         bool   `json:"isIsolated"`
	OrigClientOrderId  string `json:"origClientOrderId"`
	OrderId            int    `json:"orderId"`
	OrderListId        int    `json:"orderListId"`
	ClientOrderId      string `json:"clientOrderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
}

// CrossMarginTransferHistoryResponse define cross margin transfer history response
type CrossMarginTransferHistoryResponse struct {
	Rows []struct {
		Amount    string `json:"amount"`
		Asset     string `json:"asset"`
		Status    string `json:"status"`
		Timestamp uint64 `json:"timestamp"`
		TxId      int64  `json:"txId"`
		Type      string `json:"type"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Margin account borrow/repay(MARGIN)

// QueryMarginBorrowRepayService query
//
//gen:new_service
type QueryMarginBorrowRepayService struct {
	C              *connector.Connector
	asset          string
	ty             string
	isolatedSymbol *string
	txid           *int64
	startTime      *uint64
	endTime        *uint64
	current        *int
	size           *int
	archived       *string
}

// Asset set asset
func (s *QueryMarginBorrowRepayService) Asset(asset string) *QueryMarginBorrowRepayService {
	s.asset = asset
	return s
}

func (s *QueryMarginBorrowRepayService) Type(ty string) *QueryMarginBorrowRepayService {
	s.ty = ty
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *QueryMarginBorrowRepayService) IsolatedSymbol(isolatedSymbol string) *QueryMarginBorrowRepayService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// TxId set txid
func (s *QueryMarginBorrowRepayService) TxId(txid int64) *QueryMarginBorrowRepayService {
	s.txid = &txid
	return s
}

// StartTime set startTime
func (s *QueryMarginBorrowRepayService) StartTime(startTime uint64) *QueryMarginBorrowRepayService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *QueryMarginBorrowRepayService) EndTime(endTime uint64) *QueryMarginBorrowRepayService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *QueryMarginBorrowRepayService) Current(current int) *QueryMarginBorrowRepayService {
	s.current = &current
	return s
}

// Size set size
func (s *QueryMarginBorrowRepayService) Size(size int) *QueryMarginBorrowRepayService {
	s.size = &size
	return s
}

// Archived set archived
func (s *QueryMarginBorrowRepayService) Archived(archived string) *QueryMarginBorrowRepayService {
	s.archived = &archived
	return s
}

// Do send request
func (s *QueryMarginBorrowRepayService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryMarginBorrowRepay, err error) {

	r := request.New(
		"/sapi/v1/margin/borrow-repay",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset", "type"),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("type", s.ty)

	r.SetParam("isolatedSymbol", s.isolatedSymbol)
	r.SetParam("txId", s.txid)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)
	r.SetParam("archived", s.archived)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(QueryMarginBorrowRepay)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return
}

// QueryMarginBorrowRepay define loan record response
type QueryMarginBorrowRepay struct {
	Rows []struct {
		IsolatedSymbol string `json:"isolatedSymbol"`
		Amount         string `json:"amount"`
		Interest       string `json:"interest"`
		TxId           int64  `json:"txId"`
		Asset          string `json:"asset"`
		Principal      string `json:"principal"`
		Timestamp      uint64 `json:"timestamp"`
		Status         string `json:"status"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Interest History (USER_DATA) API Endpoint

// InterestHistoryService query interest history
//
//gen:new_service
type InterestHistoryService struct {
	C              *connector.Connector
	asset          *string
	isolatedSymbol *string
	startTime      *uint64
	endTime        *uint64
	current        *int
	size           *int
}

// Asset set asset
func (s *InterestHistoryService) Asset(asset string) *InterestHistoryService {
	s.asset = &asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *InterestHistoryService) IsolatedSymbol(isolatedSymbol string) *InterestHistoryService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// StartTime set startTime
func (s *InterestHistoryService) StartTime(startTime uint64) *InterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *InterestHistoryService) EndTime(endTime uint64) *InterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *InterestHistoryService) Current(current int) *InterestHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *InterestHistoryService) Size(size int) *InterestHistoryService {
	s.size = &size
	return s
}

// Do send request
func (s *InterestHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *InterestHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/interestHistory",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("isolatedSymbol", s.isolatedSymbol)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(InterestHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// InterestHistoryResponse define interest history response
type InterestHistoryResponse struct {
	Rows []struct {
		TxId                int64  `json:"txId"`
		InterestAccruedTime uint64 `json:"interestAccruedTime"`
		Asset               string `json:"asset"`
		RawAsset            string `json:"rawAsset"`
		Principal           string `json:"principal"`
		Interest            string `json:"interest"`
		InterestRate        string `json:"interestRate"`
		Type                string `json:"type"`
		IsolatedSymbol      string `json:"isolatedSymbol"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Force Liquidation Record (USER_DATA) API Endpoint

// ForceLiquidationRecordService query force liquidation record
//
//gen:new_service
type ForceLiquidationRecordService struct {
	C              *connector.Connector
	startTime      *uint64
	endTime        *uint64
	isolatedSymbol *string
	current        *int
	size           *int
}

// IsolatedSymbol set isolatedSymbol
func (s *ForceLiquidationRecordService) IsolatedSymbol(isolatedSymbol string) *ForceLiquidationRecordService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// StartTime set startTime
func (s *ForceLiquidationRecordService) StartTime(startTime uint64) *ForceLiquidationRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ForceLiquidationRecordService) EndTime(endTime uint64) *ForceLiquidationRecordService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *ForceLiquidationRecordService) Current(current int) *ForceLiquidationRecordService {
	s.current = &current
	return s
}

// Size set size
func (s *ForceLiquidationRecordService) Size(size int) *ForceLiquidationRecordService {
	s.size = &size
	return s
}

// Do send request
func (s *ForceLiquidationRecordService) Do(ctx context.Context, opts ...request.RequestOption) (res *ForceLiquidationRecordResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/forceLiquidationRec",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("isolatedSymbol", s.isolatedSymbol)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ForceLiquidationRecordResponse)
	err = json.Unmarshal(data, res)
	return
}

// ForceLiquidationRecordResponse define force liquidation record response
type ForceLiquidationRecordResponse struct {
	Rows []struct {
		AvgPrice    string `json:"avgPrice"`
		ExecutedQty string `json:"executedQty"`
		OrderId     int    `json:"orderId"`
		Price       string `json:"price"`
		Qty         string `json:"qty"`
		Side        string `json:"side"`
		Symbol      string `json:"symbol"`
		TimeInForce string `json:"timeInForce"`
		IsIsolated  bool   `json:"isIsolated"`
		UpdatedTime uint64 `json:"updatedTime"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Query Query Cross Margin Account Details (USER_DATA) API Endpoint

// CrossMarginAccountDetailService query cross margin account details
//
//gen:new_service
type CrossMarginAccountDetailService struct {
	C *connector.Connector
}

// Do send request
func (s *CrossMarginAccountDetailService) Do(ctx context.Context, opts ...request.RequestOption) (res *CrossMarginAccountDetailResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/account",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(CrossMarginAccountDetailResponse)
	err = json.Unmarshal(data, res)
	return
}

// CrossMarginAccountDetailResponse define cross margin account detail response
type CrossMarginAccountDetailResponse struct {
	BorrowEnabled       bool   `json:"borrowEnabled"`
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool   `json:"tradeEnabled"`
	TransferEnabled     bool   `json:"transferEnabled"`
	UserAssets          []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	} `json:"userAssets"`
}

// Query Margin Account's Order (USER_DATA) API Endpoint

// MarginAccountOrderService query margin account's order
//
//gen:new_service
type MarginAccountOrderService struct {
	C                 *connector.Connector
	symbol            string
	isIsolated        *string
	orderId           *int
	origClientOrderId *string
}

// Symbol set symbol
func (s *MarginAccountOrderService) Symbol(symbol string) *MarginAccountOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountOrderService) IsIsolated(isIsolated string) *MarginAccountOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountOrderService) OrderId(orderId int) *MarginAccountOrderService {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountOrderService) OrigClientOrderId(origClientOrderId string) *MarginAccountOrderService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *MarginAccountOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountOrderResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/order",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("orderId", s.orderId)
	r.SetParam("origClientOrderId", s.origClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountOrderResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountOrderResponse define margin account order response
type MarginAccountOrderResponse struct {
	ClientOrderId      string `json:"clientOrderId"`
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
	ExecutedQty        string `json:"executedQty"`
	IcebergQty         string `json:"icebergQty"`
	IsWorking          bool   `json:"isWorking"`
	OrderId            int    `json:"orderId"`
	OrigQty            string `json:"origQty"`
	Price              string `json:"price"`
	Side               string `json:"side"`
	Status             string `json:"status"`
	StopPrice          string `json:"stopPrice"`
	Symbol             string `json:"symbol"`
	IsIsolated         bool   `json:"isIsolated"`
	Time               uint64 `json:"time"`
	TimeInForce        string `json:"timeInForce"`
	OrderType          string `json:"type"`
	UpdateTime         uint64 `json:"updateTime"`
}

// Query Margin Account's Open Order (USER_DATA) API Endpoint

// MarginAccountOpenOrderService query margin account's open order
//
//gen:new_service
type MarginAccountOpenOrderService struct {
	C          *connector.Connector
	symbol     *string
	isIsolated *string
}

// Symbol set symbol
func (s *MarginAccountOpenOrderService) Symbol(symbol string) *MarginAccountOpenOrderService {
	s.symbol = &symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountOpenOrderService) IsIsolated(isIsolated string) *MarginAccountOpenOrderService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *MarginAccountOpenOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountOpenOrderResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/openOrders",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("isIsolated", s.isIsolated)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountOpenOrderResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountOpenOrderResponse define margin account open order response
type MarginAccountOpenOrderResponse struct {
	Orders []struct {
		ClientOrderId      string `json:"clientOrderId"`
		CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
		ExecutedQty        string `json:"executedQty"`
		IcebergQty         string `json:"icebergQty"`
		IsWorking          bool   `json:"isWorking"`
		OrderId            int    `json:"orderId"`
		OrigQty            string `json:"origQty"`
		Price              string `json:"price"`
		Side               string `json:"side"`
		Status             string `json:"status"`
		StopPrice          string `json:"stopPrice"`
		Symbol             string `json:"symbol"`
		IsIsolated         bool   `json:"isIsolated"`
		Time               uint64 `json:"time"`
		TimeInForce        string `json:"timeInForce"`
		OrderType          string `json:"type"`
		UpdateTime         uint64 `json:"updateTime"`
	} `json:"orders"`
}

// Query Margin Account's All Orders (USER_DATA) API Endpoint

// MarginAccountAllOrderService query margin account's all order
//
//gen:new_service
type MarginAccountAllOrderService struct {
	C          *connector.Connector
	symbol     string
	isIsolated *string
	orderId    *int
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// Symbol set symbol
func (s *MarginAccountAllOrderService) Symbol(symbol string) *MarginAccountAllOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountAllOrderService) IsIsolated(isIsolated string) *MarginAccountAllOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountAllOrderService) OrderId(orderId int) *MarginAccountAllOrderService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *MarginAccountAllOrderService) StartTime(startTime uint64) *MarginAccountAllOrderService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountAllOrderService) EndTime(endTime uint64) *MarginAccountAllOrderService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *MarginAccountAllOrderService) Limit(limit int) *MarginAccountAllOrderService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountAllOrderService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountAllOrderResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/allOrders",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("orderId", s.orderId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountAllOrderResponse)
	err = json.Unmarshal(data, res)

	return
}

// MarginAccountAllOrderResponse define margin account all order response
type MarginAccountAllOrderResponse struct {
	Orders []struct {
		ClientOrderId      string `json:"clientOrderId"`
		CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
		ExecutedQty        string `json:"executedQty"`
		IcebergQty         string `json:"icebergQty"`
		IsWorking          bool   `json:"isWorking"`
		OrderId            int    `json:"orderId"`
		OrigQty            string `json:"origQty"`
		Price              string `json:"price"`
		Side               string `json:"side"`
		Status             string `json:"status"`
		StopPrice          string `json:"stopPrice"`
		Symbol             string `json:"symbol"`
		IsIsolated         bool   `json:"isIsolated"`
		Time               uint64 `json:"time"`
		TimeInForce        string `json:"timeInForce"`
		OrderType          string `json:"type"`
		UpdateTime         uint64 `json:"updateTime"`
	} `json:"orders"`
}

// Margin Account New OCO (TRADE) API Endpoint

// MarginAccountNewOCOService create new oco order
//
//gen:new_service
type MarginAccountNewOCOService struct {
	C                    *connector.Connector
	symbol               string
	isIsolated           *string
	listClientOrderId    *string
	side                 string
	quantity             float64
	limitClientOrderId   *string
	price                float64
	limitIcebergQty      *float64
	stopClientOrderId    *string
	stopPrice            float64
	stopLimitPrice       *float64
	stopIcebergQty       *float64
	stopLimitTimeInForce *string
	newOrderRespType     *string
	sideEffectType       *string
}

// Symbol set symbol
func (s *MarginAccountNewOCOService) Symbol(symbol string) *MarginAccountNewOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountNewOCOService) IsIsolated(isIsolated string) *MarginAccountNewOCOService {
	s.isIsolated = &isIsolated
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountNewOCOService) ListClientOrderId(listClientOrderId string) *MarginAccountNewOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// Side set side
func (s *MarginAccountNewOCOService) Side(side string) *MarginAccountNewOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *MarginAccountNewOCOService) Quantity(quantity float64) *MarginAccountNewOCOService {
	s.quantity = quantity
	return s
}

// LimitClientOrderId set limitClientOrderId
func (s *MarginAccountNewOCOService) LimitClientOrderId(limitClientOrderId string) *MarginAccountNewOCOService {
	s.limitClientOrderId = &limitClientOrderId
	return s
}

// Price set price
func (s *MarginAccountNewOCOService) Price(price float64) *MarginAccountNewOCOService {
	s.price = price
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *MarginAccountNewOCOService) LimitIcebergQty(limitIcebergQty float64) *MarginAccountNewOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderId set stopClientOrderId
func (s *MarginAccountNewOCOService) StopClientOrderId(stopClientOrderId string) *MarginAccountNewOCOService {
	s.stopClientOrderId = &stopClientOrderId
	return s
}

// StopPrice set stopPrice
func (s *MarginAccountNewOCOService) StopPrice(stopPrice float64) *MarginAccountNewOCOService {
	s.stopPrice = stopPrice
	return s
}

// StopLimitPrice set stopLimitPrice
func (s *MarginAccountNewOCOService) StopLimitPrice(stopLimitPrice float64) *MarginAccountNewOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stopIcebergQty
func (s *MarginAccountNewOCOService) StopIcebergQty(stopIcebergQty float64) *MarginAccountNewOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *MarginAccountNewOCOService) StopLimitTimeInForce(stopLimitTimeInForce string) *MarginAccountNewOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set newOrderRespType
func (s *MarginAccountNewOCOService) NewOrderRespType(newOrderRespType string) *MarginAccountNewOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *MarginAccountNewOCOService) SideEffectType(sideEffectType string) *MarginAccountNewOCOService {
	s.sideEffectType = &sideEffectType
	return s
}

// Do send request
func (s *MarginAccountNewOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountNewOCOResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/orderList",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol", "side", "quantity", "price", "stopPrice"),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("side", s.side)
	r.SetParam("quantity", s.quantity)
	r.SetParam("price", s.price)
	r.SetParam("stopPrice", s.stopPrice)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("listClientOrderId", s.listClientOrderId)
	r.SetParam("limitClientOrderId", s.limitClientOrderId)
	r.SetParam("limitIcebergQty", s.limitIcebergQty)
	r.SetParam("stopClientOrderId", s.stopClientOrderId)
	r.SetParam("stopLimitPrice", s.stopLimitPrice)
	r.SetParam("stopIcebergQty", s.stopIcebergQty)
	r.SetParam("stopLimitTimeInForce", s.stopLimitTimeInForce)
	r.SetParam("newOrderRespType", s.newOrderRespType)
	r.SetParam("sideEffectType", s.sideEffectType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountNewOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountNewOCOService response
type MarginAccountNewOCOResponse struct {
	OrderListId           int    `json:"orderListId"`
	ContingencyType       string `json:"contingencyType"`
	ListStatusType        string `json:"listStatusType"`
	ListOrderStatus       string `json:"listOrderStatus"`
	ListClientOrderId     string `json:"listClientOrderId"`
	TransactionTime       uint64 `json:"transactionTime"`
	Symbol                string `json:"symbol"`
	MarginBuyBorrowAmount string `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string `json:"marginBuyBorrowAsset"`
	Orders                []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol             string `json:"symbol"`
		OrderId            int    `json:"orderId"`
		OrderListId        int    `json:"orderListId"`
		ClientOrderId      string `json:"clientOrderId"`
		TransactTime       uint64 `json:"transactTime"`
		Price              string `json:"price"`
		OrigQty            string `json:"origQty"`
		ExecutedQty        string `json:"executedQty"`
		CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
		Status             string `json:"status"`
		TimeInForce        string `json:"timeInForce"`
		OrderType          string `json:"type"`
		Side               string `json:"side"`
		StopPrice          string `json:"stopPrice"`
	} `json:"orderReports"`
}

// Margin Account Cancel OCO (TRADE)
//
//gen:new_service
type MarginAccountCancelOCOService struct {
	C                 *connector.Connector
	symbol            string
	isIsolated        *string
	orderListId       *int
	listClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *MarginAccountCancelOCOService) Symbol(symbol string) *MarginAccountCancelOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountCancelOCOService) IsIsolated(isIsolated string) *MarginAccountCancelOCOService {
	s.isIsolated = &isIsolated
	return s
}

// OrderListId set orderListId
func (s *MarginAccountCancelOCOService) OrderListId(orderListId int) *MarginAccountCancelOCOService {
	s.orderListId = &orderListId
	return s
}

// ListClientOrderId set listClientOrderId
func (s *MarginAccountCancelOCOService) ListClientOrderId(listClientOrderId string) *MarginAccountCancelOCOService {
	s.listClientOrderId = &listClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *MarginAccountCancelOCOService) NewClientOrderId(newClientOrderId string) *MarginAccountCancelOCOService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *MarginAccountCancelOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountCancelOCOResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/orderList",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("orderListId", s.orderListId)
	r.SetParam("listClientOrderId", s.listClientOrderId)
	r.SetParam("newClientOrderId", s.newClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountCancelOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountCancelOCOService response
type MarginAccountCancelOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol             string `json:"symbol"`
		OrigClientOrderId  string `json:"origClientOrderId"`
		OrderId            int    `json:"orderId"`
		OrderListId        int    `json:"orderListId"`
		ClientOrderId      string `json:"clientOrderId"`
		Price              string `json:"price"`
		OrigQty            string `json:"origQty"`
		ExecutedQty        string `json:"executedQty"`
		CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
		Status             string `json:"status"`
		TimeInForce        string `json:"timeInForce"`
		OrderType          string `json:"type"`
		Side               string `json:"side"`
		StopPrice          string `json:"stopPrice"`
	} `json:"orderReports"`
}

// Query Margin Account's OCO (USER_DATA) (HMAC SHA256)

//gen:new_service
type MarginAccountQueryOCOService struct {
	C                 *connector.Connector
	isIsolated        *string
	symbol            *string
	orderListId       *int
	origClientOrderId *string
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryOCOService) IsIsolated(isIsolated string) *MarginAccountQueryOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryOCOService) Symbol(symbol string) *MarginAccountQueryOCOService {
	s.symbol = &symbol
	return s
}

// OrderListId set orderListId
func (s *MarginAccountQueryOCOService) OrderListId(orderListId int) *MarginAccountQueryOCOService {
	s.orderListId = &orderListId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *MarginAccountQueryOCOService) OrigClientOrderId(origClientOrderId string) *MarginAccountQueryOCOService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *MarginAccountQueryOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryOCOResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/orderList",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("symbol", s.symbol)
	r.SetParam("orderListId", s.orderListId)
	r.SetParam("origClientOrderId", s.origClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryOCOService response
type MarginAccountQueryOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's all OCO (USER_DATA)

//gen:new_service
type MarginAccountQueryAllOCOService struct {
	C          *connector.Connector
	isIsolated *string
	symbol     *string
	fromId     *int
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryAllOCOService) IsIsolated(isIsolated string) *MarginAccountQueryAllOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryAllOCOService) Symbol(symbol string) *MarginAccountQueryAllOCOService {
	s.symbol = &symbol
	return s
}

// FromId set fromId
func (s *MarginAccountQueryAllOCOService) FromId(fromId int) *MarginAccountQueryAllOCOService {
	s.fromId = &fromId
	return s
}

// StartTime set startTime
func (s *MarginAccountQueryAllOCOService) StartTime(startTime uint64) *MarginAccountQueryAllOCOService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountQueryAllOCOService) EndTime(endTime uint64) *MarginAccountQueryAllOCOService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *MarginAccountQueryAllOCOService) Limit(limit int) *MarginAccountQueryAllOCOService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountQueryAllOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryAllOCOResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/allOrderList",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("symbol", s.symbol)
	r.SetParam("fromId", s.fromId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryAllOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryAllOCOService response
type MarginAccountQueryAllOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's Open OCO (USER_DATA)
//
//gen:new_service
type MarginAccountQueryOpenOCOService struct {
	C          *connector.Connector
	isIsolated *string
	symbol     *string
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryOpenOCOService) IsIsolated(isIsolated string) *MarginAccountQueryOpenOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginAccountQueryOpenOCOService) Symbol(symbol string) *MarginAccountQueryOpenOCOService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginAccountQueryOpenOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryOpenOCOResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/openOrderList",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryOpenOCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryOpenOCOService response
type MarginAccountQueryOpenOCOResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	IsIsolated        bool   `json:"isIsolated"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

// Query Margin Account's Trade List (USER_DATA)
//
//gen:new_service
type MarginAccountQueryTradeListService struct {
	C          *connector.Connector
	symbol     string
	isIsolated *string
	orderId    *int
	startTime  *uint64
	endTime    *uint64
	fromId     *int
	limit      *int
}

// Symbol set symbol
func (s *MarginAccountQueryTradeListService) Symbol(symbol string) *MarginAccountQueryTradeListService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *MarginAccountQueryTradeListService) IsIsolated(isIsolated string) *MarginAccountQueryTradeListService {
	s.isIsolated = &isIsolated
	return s
}

// OrderId set orderId
func (s *MarginAccountQueryTradeListService) OrderId(orderId int) *MarginAccountQueryTradeListService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *MarginAccountQueryTradeListService) StartTime(startTime uint64) *MarginAccountQueryTradeListService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginAccountQueryTradeListService) EndTime(endTime uint64) *MarginAccountQueryTradeListService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *MarginAccountQueryTradeListService) FromId(fromId int) *MarginAccountQueryTradeListService {
	s.fromId = &fromId
	return s
}

// Limit set limit
func (s *MarginAccountQueryTradeListService) Limit(limit int) *MarginAccountQueryTradeListService {
	s.limit = &limit
	return s
}

// Do send request
func (s *MarginAccountQueryTradeListService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryTradeListResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/myTrades",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("orderId", s.orderId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("fromId", s.fromId)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryTradeListResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryTradeListService response
type MarginAccountQueryTradeListResponse struct {
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              int    `json:"id"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	OrderId         int    `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Symbol          string `json:"symbol"`
	IsIsolated      bool   `json:"isIsolated"`
	Time            uint64 `json:"time"`
}

// Query Margin Account's Max Borrow (USER_DATA)
//
//gen:new_service
type MarginAccountQueryMaxBorrowService struct {
	C              *connector.Connector
	asset          string
	isolatedSymbol *string
}

// Asset set asset
func (s *MarginAccountQueryMaxBorrowService) Asset(asset string) *MarginAccountQueryMaxBorrowService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *MarginAccountQueryMaxBorrowService) IsolatedSymbol(isolatedSymbol string) *MarginAccountQueryMaxBorrowService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// Do send request
func (s *MarginAccountQueryMaxBorrowService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryMaxBorrowResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/maxBorrowable",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset"),
	)

	r.SetParam("asset", s.asset)

	r.SetParam("isolatedSymbol", s.isolatedSymbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryMaxBorrowResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryMaxBorrowService response
type MarginAccountQueryMaxBorrowResponse struct {
	Amount      string `json:"amount"`
	BorrowLimit string `json:"borrowLimit"`
}

// Query Margin Account's Max Transfer-Out Amount (USER_DATA)
//
//gen:new_service
type MarginAccountQueryMaxTransferOutAmountService struct {
	C              *connector.Connector
	asset          string
	isolatedSymbol *string
}

// Asset set asset
func (s *MarginAccountQueryMaxTransferOutAmountService) Asset(asset string) *MarginAccountQueryMaxTransferOutAmountService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *MarginAccountQueryMaxTransferOutAmountService) IsolatedSymbol(isolatedSymbol string) *MarginAccountQueryMaxTransferOutAmountService {
	s.isolatedSymbol = &isolatedSymbol
	return s
}

// Do send request
func (s *MarginAccountQueryMaxTransferOutAmountService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountQueryMaxTransferOutAmountResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/maxTransferable",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset"),
	)

	r.SetParam("asset", s.asset)

	r.SetParam("isolatedSymbol", s.isolatedSymbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountQueryMaxTransferOutAmountResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountQueryMaxTransferOutAmountService response
type MarginAccountQueryMaxTransferOutAmountResponse struct {
	Amount string `json:"amount"`
}

// Get Summary of Margin account (USER_DATA) - GET /sapi/v1/margin/tradeCoeff (HMAC SHA256)
//
//gen:new_service
type MarginAccountSummaryService struct {
	C *connector.Connector
}

// Do send request
func (s *MarginAccountSummaryService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginAccountSummaryResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/tradeCoeff",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginAccountSummaryResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginAccountSummaryService response
type MarginAccountSummaryResponse struct {
	NormalBar           string `json:"normalBar"`
	MarginCallBar       string `json:"marginCallBar"`
	ForceLiquidationBar string `json:"forceLiquidationBar"`
}

// Isolated Margin Account Transfer (MARGIN)
//
//gen:new_service
type MarginIsolatedAccountTransferService struct {
	C         *connector.Connector
	asset     string
	symbol    string
	transFrom string
	transTo   string
	amount    float64
}

// Asset set asset
func (s *MarginIsolatedAccountTransferService) Asset(asset string) *MarginIsolatedAccountTransferService {
	s.asset = asset
	return s
}

// Symbol set symbol
func (s *MarginIsolatedAccountTransferService) Symbol(symbol string) *MarginIsolatedAccountTransferService {
	s.symbol = symbol
	return s
}

// TransFrom set transFrom
func (s *MarginIsolatedAccountTransferService) TransFrom(transFrom string) *MarginIsolatedAccountTransferService {
	s.transFrom = transFrom
	return s
}

// TransTo set transTo
func (s *MarginIsolatedAccountTransferService) TransTo(transTo string) *MarginIsolatedAccountTransferService {
	s.transTo = transTo
	return s
}

// Amount set amount
func (s *MarginIsolatedAccountTransferService) Amount(amount float64) *MarginIsolatedAccountTransferService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginIsolatedAccountTransferService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedAccountTransferResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/transfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset", "symbol", "transFrom", "transTo", "amount"),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("symbol", s.symbol)
	r.SetParam("transFrom", s.transFrom)
	r.SetParam("transTo", s.transTo)
	r.SetParam("amount", s.amount)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedAccountTransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedAccountTransferService response
type MarginIsolatedAccountTransferResponse struct {
	TranId string `json:"tranId"`
}

// Isolated Margin Account Transfer History (MARGIN)
//
//gen:new_service
type MarginIsolatedAccountTransferHistoryService struct {
	C         *connector.Connector
	asset     *string
	symbol    string
	transFrom *string
	transTo   *string
	startTime *uint64
	endTime   *uint64
	current   *int
	size      *int
	archived  *string
}

// Asset set asset
func (s *MarginIsolatedAccountTransferHistoryService) Asset(asset string) *MarginIsolatedAccountTransferHistoryService {
	s.asset = &asset
	return s
}

// Symbol set symbol
func (s *MarginIsolatedAccountTransferHistoryService) Symbol(symbol string) *MarginIsolatedAccountTransferHistoryService {
	s.symbol = symbol
	return s
}

// TransFrom set transFrom
func (s *MarginIsolatedAccountTransferHistoryService) TransFrom(transFrom string) *MarginIsolatedAccountTransferHistoryService {
	s.transFrom = &transFrom
	return s
}

// TransTo set transTo
func (s *MarginIsolatedAccountTransferHistoryService) TransTo(transTo string) *MarginIsolatedAccountTransferHistoryService {
	s.transTo = &transTo
	return s
}

// StartTime set startTime
func (s *MarginIsolatedAccountTransferHistoryService) StartTime(startTime uint64) *MarginIsolatedAccountTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginIsolatedAccountTransferHistoryService) EndTime(endTime uint64) *MarginIsolatedAccountTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *MarginIsolatedAccountTransferHistoryService) Current(current int) *MarginIsolatedAccountTransferHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *MarginIsolatedAccountTransferHistoryService) Size(size int) *MarginIsolatedAccountTransferHistoryService {
	s.size = &size
	return s
}

// Archived set archived
func (s *MarginIsolatedAccountTransferHistoryService) Archived(archived string) *MarginIsolatedAccountTransferHistoryService {
	s.archived = &archived
	return s
}

// Do send request
func (s *MarginIsolatedAccountTransferHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedAccountTransferHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/transfer",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("asset", s.asset)
	r.SetParam("transFrom", s.transFrom)
	r.SetParam("transTo", s.transTo)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("current", s.current)
	r.SetParam("size", s.size)
	r.SetParam("archived", s.archived)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedAccountTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedAccountTransferHistoryService response
type MarginIsolatedAccountTransferHistoryResponse struct {
	Rows []struct {
		Amount    string `json:"amount"`
		Asset     string `json:"asset"`
		Status    string `json:"status"`
		TimeStamp uint64 `json:"timeStamp"`
		TxId      int64  `json:"txId"`
		TransFrom string `json:"transFrom"`
		TransTo   string `json:"transTo"`
	} `json:"rows"`
	Total int64 `json:"total"`
}

// Query Isolated Margin Account Info (USER_DATA)

//gen:new_service
type MarginIsolatedAccountInfoService struct {
	C       *connector.Connector
	symbols *string
}

// Symbols set symbols
func (s *MarginIsolatedAccountInfoService) Symbols(symbols string) *MarginIsolatedAccountInfoService {
	s.symbols = &symbols
	return s
}

// Do send request
func (s *MarginIsolatedAccountInfoService) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/account",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("symbols", s.symbols)

	if s.symbols != nil {
		res = new(MarginIsolatedAccountInfoResponseSymbols)
	} else {
		res = new(MarginIsolatedAccountInfoResponse)
	}

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, res)
	return
}

type MarginIsolatedAccountInfoAssets struct {
	BaseAsset struct {
		Asset         string `json:"asset"`
		BorrowEnabled bool   `json:"borrowEnabled"`
		Free          string `json:"free"`
		Interest      string `json:"interest"`
		Locked        string `json:"locked"`
		NetAsset      string `json:"netAsset"`
		NetAssetOfBtc string `json:"netAssetOfBtc"`
		RepayEnabled  bool   `json:"repayEnabled"`
		TotalAsset    string `json:"totalAsset"`
	} `json:"baseAsset"`
	QuoteAsset struct {
		Asset         string `json:"asset"`
		BorrowEnabled bool   `json:"borrowEnabled"`
		Free          string `json:"free"`
		Interest      string `json:"interest"`
		Locked        string `json:"locked"`
		NetAsset      string `json:"netAsset"`
		NetAssetOfBtc string `json:"netAssetOfBtc"`
		RepayEnabled  bool   `json:"repayEnabled"`
		TotalAsset    string `json:"totalAsset"`
	} `json:"quoteAsset"`
	Symbol            string `json:"symbol"`
	IsolatedCreated   bool   `json:"isolatedCreated"`
	Enabled           bool   `json:"enabled"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// MarginIsolatedAccountInfoService response if symbols parameter is not sent
type MarginIsolatedAccountInfoResponseSymbols struct {
	Assets []*MarginIsolatedAccountInfoAssets `json:"assets"`
}

// MarginIsolatedAccountInfoService response if symbols parameter is sent
type MarginIsolatedAccountInfoResponse struct {
	Assets              []*MarginIsolatedAccountInfoAssets `json:"assets"`
	TotalAssetOfBtc     string                             `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string                             `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string                             `json:"totalNetAssetOfBtc"`
}

// Disable Isolated Margin Account (TRADE)
//
//gen:new_service
type MarginIsolatedAccountDisableService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *MarginIsolatedAccountDisableService) Symbol(symbol string) *MarginIsolatedAccountDisableService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *MarginIsolatedAccountDisableService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedAccountDisableResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/account",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedAccountDisableResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedAccountDisableService response
type MarginIsolatedAccountDisableResponse struct {
	Success bool `json:"success"`
}

// Enable Isolated Margin Account (TRADE)
//
//gen:new_service
type MarginIsolatedAccountEnableService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *MarginIsolatedAccountEnableService) Symbol(symbol string) *MarginIsolatedAccountEnableService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *MarginIsolatedAccountEnableService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedAccountEnableResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/account",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedAccountEnableResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedAccountEnableService response
type MarginIsolatedAccountEnableResponse struct {
	Success bool `json:"success"`
}

// Query Enabled Isolated Margin Account Limit (USER_DATA)
//
//gen:new_service
type MarginIsolatedAccountLimitService struct {
	C *connector.Connector
}

// Do send request
func (s *MarginIsolatedAccountLimitService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedAccountLimitResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/accountLimit",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedAccountLimitResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedAccountLimitService response
type MarginIsolatedAccountLimitResponse struct {
	EnabledAccount int `json:"enabledAccount"`
	MaxAccount     int `json:"maxAccount"`
}

// Query Isolated Margin Symbol (USER_DATA)
//
//gen:new_service
type MarginIsolatedSymbolService struct {
	C      *connector.Connector
	symbol string
}

// Symbol set symbol
func (s *MarginIsolatedSymbolService) Symbol(symbol string) *MarginIsolatedSymbolService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *MarginIsolatedSymbolService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginIsolatedSymbolResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/pair",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginIsolatedSymbolResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginIsolatedSymbolService response
type MarginIsolatedSymbolResponse struct {
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// Get All Isolated Margin Symbol(USER_DATA)

// AllIsolatedMarginSymbolService returns all isolated margin symbols
//
//gen:new_service
type AllIsolatedMarginSymbolService struct {
	C *connector.Connector
}

// Do send request
func (s *AllIsolatedMarginSymbolService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginIsolatedSymbolResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/allPairs",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginIsolatedSymbolResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginIsolatedSymbolAllService returns all isolated margin symbols

// Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
//
//gen:new_service
type MarginToggleBnbBurnService struct {
	C               *connector.Connector
	spotBNBBurn     *string
	interestBNBBurn *string
}

// SpotBNBBurn set spotBNBBurn
func (s *MarginToggleBnbBurnService) SpotBNBBurn(spotBNBBurn string) *MarginToggleBnbBurnService {
	s.spotBNBBurn = &spotBNBBurn
	return s
}

// InterestBNBBurn set interestBNBBurn
func (s *MarginToggleBnbBurnService) InterestBNBBurn(interestBNBBurn string) *MarginToggleBnbBurnService {
	s.interestBNBBurn = &interestBNBBurn
	return s
}

// Do send request
func (s *MarginToggleBnbBurnService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginToggleBnbBurnResponse, err error) {

	r := request.New(
		"/sapi/v1/bnbBurn",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("spotBNBBurn", s.spotBNBBurn)
	r.SetParam("interestBNBBurn", s.interestBNBBurn)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginToggleBnbBurnResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginToggleBnbBurnService response
type MarginToggleBnbBurnResponse struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// Get BNB Burn Status (USER_DATA)
//
//gen:new_service
type MarginBnbBurnStatusService struct {
	C *connector.Connector
}

// Do send request
func (s *MarginBnbBurnStatusService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginBnbBurnStatusResponse, err error) {

	r := request.New(
		"/sapi/v1/bnbBurn",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(MarginBnbBurnStatusResponse)
	err = json.Unmarshal(data, res)
	return
}

// MarginBnbBurnStatusService response
type MarginBnbBurnStatusResponse struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// Query Margin Interest Rate History (USER_DATA)
//
//gen:new_service
type MarginInterestRateHistoryService struct {
	C         *connector.Connector
	asset     string
	vipLevel  *int
	startTime *uint64
	endTime   *uint64
}

// Asset set asset
func (s *MarginInterestRateHistoryService) Asset(asset string) *MarginInterestRateHistoryService {
	s.asset = asset
	return s
}

// VipLevel set vipLevel
func (s *MarginInterestRateHistoryService) VipLevel(vipLevel int) *MarginInterestRateHistoryService {
	s.vipLevel = &vipLevel
	return s
}

// StartTime set startTime
func (s *MarginInterestRateHistoryService) StartTime(startTime uint64) *MarginInterestRateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginInterestRateHistoryService) EndTime(endTime uint64) *MarginInterestRateHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *MarginInterestRateHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginInterestRateHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/interestRateHistory",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("asset"),
	)

	r.SetParam("asset", s.asset)

	r.SetParam("vipLevel", s.vipLevel)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginInterestRateHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginInterestRateHistoryService response
type MarginInterestRateHistoryResponse struct {
	Asset             string  `json:"asset"`
	DailyInterestRate float64 `json:"dailyInterestRate"`
	Timestamp         uint64  `json:"timestamp"`
	VIPLevel          int     `json:"vipLevel"`
}

// Query Cross Margin Fee Data (USER_DATA)
//
//gen:new_service
type MarginCrossMarginFeeService struct {
	C        *connector.Connector
	vipLevel *int
	coin     *string
}

// VipLevel set vipLevel
func (s *MarginCrossMarginFeeService) VipLevel(vipLevel int) *MarginCrossMarginFeeService {
	s.vipLevel = &vipLevel
	return s
}

// Coin set coin
func (s *MarginCrossMarginFeeService) Coin(coin string) *MarginCrossMarginFeeService {
	s.coin = &coin
	return s
}

// Do send request
func (s *MarginCrossMarginFeeService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginCrossMarginFeeResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/crossMarginFee",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("vipLevel", s.vipLevel)
	r.SetParam("coin", s.coin)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginCrossMarginFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginCrossMarginFeeService response
type MarginCrossMarginFeeResponse struct {
	VIPLevel        int    `json:"vipLevel"`
	Coin            string `json:"coin"`
	TransferIn      bool   `json:"transferIn"`
	Borrowable      bool   `json:"transferOut"`
	DailyInterest   string `json:"dailyInterest"`
	YearlyInterest  string `json:"yearlyInterest"`
	BorrowLimit     string `json:"borrowLimit"`
	MarginablePairs struct {
		Pair string `json:"pair"`
	} `json:"marginablePairs"`
}

// Query Isolated Margin Fee Data (USER_DATA)

//gen:new_service
type MarginIsolatedMarginFeeService struct {
	C        *connector.Connector
	vipLevel *int
	symbol   *string
}

// VipLevel set vipLevel
func (s *MarginIsolatedMarginFeeService) VipLevel(vipLevel int) *MarginIsolatedMarginFeeService {
	s.vipLevel = &vipLevel
	return s
}

// Symbol set symbol
func (s *MarginIsolatedMarginFeeService) Symbol(symbol string) *MarginIsolatedMarginFeeService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginIsolatedMarginFeeService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginIsolatedMarginFeeResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolatedMarginFee",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	r.SetParam("vipLevel", s.vipLevel)
	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginIsolatedMarginFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginIsolatedMarginFeeService response
type MarginIsolatedMarginFeeResponse struct {
	VIPLevel int    `json:"vipLevel"`
	Symbol   string `json:"symbol"`
	Leverage string `json:"leverage"`
	Data     struct {
		Coin          string `json:"coin"`
		DailyInterest string `json:"dailyInterest"`
		BorrowLimit   string `json:"borrowLimit"`
	} `json:"data"`
}

// Query Isolated Margin Tier Data (USER_DATA)

//gen:new_service
type MarginIsolatedMarginTierService struct {
	C      *connector.Connector
	symbol string
	tier   *int
}

// Symbol set symbol
func (s *MarginIsolatedMarginTierService) Symbol(symbol string) *MarginIsolatedMarginTierService {
	s.symbol = symbol
	return s
}

// Tier set tier
func (s *MarginIsolatedMarginTierService) Tier(tier int) *MarginIsolatedMarginTierService {
	s.tier = &tier
	return s
}

// Do send request
func (s *MarginIsolatedMarginTierService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginIsolatedMarginTierResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/isolated/marginTier",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
	)

	r.SetParam("symbol", s.symbol)

	r.SetParam("tier", s.tier)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginIsolatedMarginTierResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginIsolatedMarginTierService response
type MarginIsolatedMarginTierResponse struct {
	Symbol                  string `json:"symbol"`
	Tier                    int    `json:"tier"`
	EffectiveMultiple       string `json:"effectiveMultiple"`
	InitialRiskRatio        string `json:"initialRiskRatio"`
	LiquidationRiskRatio    string `json:"liquidationRiskRatio"`
	BaseAssetMaxBorrowable  string `json:"baseAssetMaxBorrowable"`
	QuoteAssetMaxBorrowable string `json:"quoteAssetMaxBorrowable"`
}

// Query Current Margin Order Count Usage (TRADE)
//
//gen:new_service
type MarginCurrentOrderCountService struct {
	C          *connector.Connector
	isIsolated *string
	symbol     *string
}

// IsIsolated set isIsolated
func (s *MarginCurrentOrderCountService) IsIsolated(isIsolated string) *MarginCurrentOrderCountService {
	s.isIsolated = &isIsolated
	return s
}

// Symbol set symbol
func (s *MarginCurrentOrderCountService) Symbol(symbol string) *MarginCurrentOrderCountService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginCurrentOrderCountService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginCurrentOrderCountResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/rateLimit/order",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("isIsolated", s.isIsolated)
	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginCurrentOrderCountResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginCurrentOrderCountService response
type MarginCurrentOrderCountResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

// Margin Dustlog (USER_DATA)
//
//gen:new_service
type MarginDustlogService struct {
	C         *connector.Connector
	startTime *uint64
	endTime   *uint64
}

// StartTime set startTime
func (s *MarginDustlogService) StartTime(startTime uint64) *MarginDustlogService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginDustlogService) EndTime(endTime uint64) *MarginDustlogService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *MarginDustlogService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginDustlogResponse, err error) {

	r := request.New(
		"/sapi/v1/asset/dribblet",
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return
	}
	res = new(MarginDustlogResponse)
	err = json.Unmarshal(data, res)
	return
}

type MarginDustlogResponse struct {
	Total              uint8               `json:"total"` //Total counts of exchange
	UserAssetDribblets []UserAssetDribblet `json:"userAssetDribblets"`
}

// UserAssetDribblet represents one dust log row
type UserAssetDribblet struct {
	OperateTime              int64                     `json:"operateTime"`
	TotalTransferedAmount    string                    `json:"totalTransferedAmount"`    //Total transfered BNB amount for this exchange.
	TotalServiceChargeAmount string                    `json:"totalServiceChargeAmount"` //Total service charge amount for this exchange.
	TransId                  int64                     `json:"transId"`
	UserAssetDribbletDetails []UserAssetDribbletDetail `json:"userAssetDribbletDetails"` //Details of this exchange.
}

// DustLog represents one dust log informations
type UserAssetDribbletDetail struct {
	TransId             int    `json:"transId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	Amount              string `json:"amount"`
	OperateTime         int64  `json:"operateTime"` //The time of this exchange.
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}

// Cross margin collateral ratio (MARKET_DATA)
//
//gen:new_service
type MarginCrossCollateralRatioService struct {
	C *connector.Connector
}

// Do send request
func (s *MarginCrossCollateralRatioService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginCrossCollateralRatioResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/crossCollateralRatio",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginCrossCollateralRatioResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginCrossCollateralRatioService response
type MarginCrossCollateralRatioResponse struct {
	Collaterals []*struct {
		MinUsdValue  string `json:"minUsdValue"`
		MaxUsdValue  string `json:"maxUsdValue"`
		DiscountRate string `json:"discountRate"`
	} `json:"collaterals"`
	AssetNames []*struct {
		Asset string `json:"asset"`
	} `json:"assetNames"`
}

// Get Small Liability Exchange Coin List (USER_DATA)
//
//gen:new_service
type MarginSmallLiabilityExchangeCoinListService struct {
	C *connector.Connector
}

// Do send request
func (s *MarginSmallLiabilityExchangeCoinListService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginSmallLiabilityExchangeCoinListResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/smallLiability/exchangeCoinList",
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginSmallLiabilityExchangeCoinListResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginSmallLiabilityExchangeCoinListService response
type MarginSmallLiabilityExchangeCoinListResponse struct {
	Asset           string `json:"asset"`
	Interest        string `json:"interest"`
	Principal       string `json:"principal"`
	LiabilityOfBUSD string `json:"liabilityOfBUSD"`
}

// Small Liability Exchange (MARGIN)
//
//gen:new_service
type MarginSmallLiabilityExchangeService struct {
	C          *connector.Connector
	assetNames string
}

// AssetNames set assetNames
func (s *MarginSmallLiabilityExchangeService) AssetNames(assetNames string) *MarginSmallLiabilityExchangeService {
	s.assetNames = assetNames
	return s
}

// Do send request
func (s *MarginSmallLiabilityExchangeService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginSmallLiabilityExchangeResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/smallLiability/exchange",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("assetNames"),
	)

	r.SetParam("assetNames", s.assetNames)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginSmallLiabilityExchangeResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginSmallLiabilityExchangeService response
type MarginSmallLiabilityExchangeResponse struct {
}

// Get Small Liability Exchange History (USER_DATA)
//
//gen:new_service
type MarginSmallLiabilityExchangeHistoryService struct {
	C         *connector.Connector
	current   int
	size      int
	startTime *uint64
	endTime   *uint64
}

// Current set current
func (s *MarginSmallLiabilityExchangeHistoryService) Current(current int) *MarginSmallLiabilityExchangeHistoryService {
	s.current = current
	return s
}

// Size set size
func (s *MarginSmallLiabilityExchangeHistoryService) Size(size int) *MarginSmallLiabilityExchangeHistoryService {
	s.size = size
	return s
}

// StartTime set startTime
func (s *MarginSmallLiabilityExchangeHistoryService) StartTime(startTime uint64) *MarginSmallLiabilityExchangeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *MarginSmallLiabilityExchangeHistoryService) EndTime(endTime uint64) *MarginSmallLiabilityExchangeHistoryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *MarginSmallLiabilityExchangeHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res []*MarginSmallLiabilityExchangeHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/margin/exchange-small-liability-history",
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("current", "size"),
	)

	r.SetParam("current", s.current)
	r.SetParam("size", s.size)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*MarginSmallLiabilityExchangeHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// MarginSmallLiabilityExchangeHistoryService response
type MarginSmallLiabilityExchangeHistoryResponse struct {
	Total int `json:"total"`
	Rows  []*struct {
		Asset        string `json:"asset"`
		Amount       string `json:"amount"`
		TargetAsset  string `json:"targetAsset"`
		TargetAmount string `json:"targetAmount"`
		BizType      string `json:"bizType"`
		Timestamp    uint64 `json:"timestamp"`
	} `json:"rows"`
}
