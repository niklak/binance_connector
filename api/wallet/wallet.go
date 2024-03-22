package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Dust Transfer (USER_DATA)
const (
	dustTransferEndpoint = "/sapi/v1/asset/dust"
)

// DustTransferService dust transfer
type DustTransferService struct {
	C     *connector.Connector
	asset []string
}

// Asset set asset
func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.asset = asset
	return s
}

func (s *DustTransferService) Do(ctx context.Context) (res *DustTransferResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: dustTransferEndpoint,
		SecType:  request.SecTypeSigned,
	}
	for _, a := range s.asset {
		r.AddParam("asset", a)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(DustTransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// DustTransferResponse define response of DustTransferService
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}

// Asset Dividend Record (USER_DATA)
const (
	assetDividendRecordEndpoint = "/sapi/v1/asset/assetDividend"
)

// AssetDividendRecordService asset dividend record
type AssetDividendRecordService struct {
	C         *connector.Connector
	asset     *string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

// Asset set asset
func (s *AssetDividendRecordService) Asset(asset string) *AssetDividendRecordService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *AssetDividendRecordService) StartTime(startTime uint64) *AssetDividendRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *AssetDividendRecordService) EndTime(endTime uint64) *AssetDividendRecordService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AssetDividendRecordService) Limit(limit int) *AssetDividendRecordService {
	s.limit = &limit
	return s
}

func (s *AssetDividendRecordService) Do(ctx context.Context) (res *AssetDividendRecordResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: assetDividendRecordEndpoint,
		SecType:  request.SecTypeSigned,
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDividendRecordResponse)
	err = json.Unmarshal(data, res)
	return
}

// AssetDividendRecordResponse define response of AssetDividendRecordService
type AssetDividendRecordResponse struct {
	Rows []struct {
		Id      int64  `json:"id"`
		Amount  string `json:"amount"`
		Asset   string `json:"asset"`
		DivTime uint64 `json:"divTime"`
		EnInfo  string `json:"enInfo"`
		TranId  int64  `json:"tranId"`
	} `json:"rows"`
	Total int64 `json:"total"`
}

// Asset Detail (USER_DATA)
const (
	assetDetailV2Endpoint = "/sapi/v1/asset/assetDetail"
)

// AssetDetailV2Service asset detail v2
type AssetDetailV2Service struct {
	C     *connector.Connector
	asset *string
}

// Asset set asset
func (s *AssetDetailV2Service) Asset(asset string) *AssetDetailV2Service {
	s.asset = &asset
	return s
}

func (s *AssetDetailV2Service) Do(ctx context.Context) (res *AssetDetailV2Response, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: assetDetailV2Endpoint,
		SecType:  request.SecTypeSigned,
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetailV2Response)
	err = json.Unmarshal(data, res)
	return
}

// AssetDetailV2Response define response of AssetDetailV2Service
type AssetDetailV2Response struct {
	AssetDetail struct {
		MinWithdrawAmount string `json:"minWithdrawAmount"`
		DepositStatus     bool   `json:"depositStatus"`
		WithdrawFee       string `json:"withdrawFee"`
		WithdrawStatus    bool   `json:"withdrawStatus"`
		DepositTip        string `json:"depositTip"`
	} `json:"assetDetail"`
}

// Trade Fee (USER_DATA)
const (
	tradeFeeEndpoint = "/sapi/v1/asset/tradeFee"
)

// TradeFeeService trade fee
type TradeFeeService struct {
	C      *connector.Connector
	symbol *string
}

// Symbol set symbol
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = &symbol
	return s
}

func (s *TradeFeeService) Do(ctx context.Context) (res []*TradeFeeResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: tradeFeeEndpoint,
		SecType:  request.SecTypeSigned,
	}
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*TradeFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// TradeFeeResponse define response of TradeFeeService
type TradeFeeResponse struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// User Universal Transfer (USER_DATA)
const (
	userUniversalTransferEndpoint = "/sapi/v1/asset/transfer"
)

// UserUniversalTransferService user universal transfer
type UserUniversalTransferService struct {
	C            *connector.Connector
	transferType string
	asset        string
	amount       float64
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferService) TransferType(transferType string) *UserUniversalTransferService {
	s.transferType = transferType
	return s
}

// Asset set asset
func (s *UserUniversalTransferService) Asset(asset string) *UserUniversalTransferService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *UserUniversalTransferService) Amount(amount float64) *UserUniversalTransferService {
	s.amount = amount
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferService) FromSymbol(fromSymbol string) *UserUniversalTransferService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferService) ToSymbol(toSymbol string) *UserUniversalTransferService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferService) Do(ctx context.Context) (res *UserUniversalTransferResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: userUniversalTransferEndpoint,
		SecType:  request.SecTypeSigned,
	}
	r.SetParam("type", s.transferType)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	if s.fromSymbol != nil {
		r.SetParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.SetParam("toSymbol", *s.toSymbol)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// UserUniversalTransferResponse define response of UserUniversalTransferService
type UserUniversalTransferResponse struct {
	TranId int64 `json:"tranId"`
}

// Query User Universal Transfer History (USER_DATA)
const (
	userUniversalTransferHistoryEndpoint = "/sapi/v1/asset/transfer"
)

// UserUniversalTransferHistoryService user universal transfer history
type UserUniversalTransferHistoryService struct {
	C            *connector.Connector
	transferType string
	startTime    *uint64
	endTime      *uint64
	current      *int
	size         *int
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferHistoryService) TransferType(transferType string) *UserUniversalTransferHistoryService {
	s.transferType = transferType
	return s
}

// StartTime set startTime
func (s *UserUniversalTransferHistoryService) StartTime(startTime uint64) *UserUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *UserUniversalTransferHistoryService) EndTime(endTime uint64) *UserUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

// Current set current
func (s *UserUniversalTransferHistoryService) Current(current int) *UserUniversalTransferHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *UserUniversalTransferHistoryService) Size(size int) *UserUniversalTransferHistoryService {
	s.size = &size
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferHistoryService) FromSymbol(fromSymbol string) *UserUniversalTransferHistoryService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferHistoryService) ToSymbol(toSymbol string) *UserUniversalTransferHistoryService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferHistoryService) Do(ctx context.Context) (res *UserUniversalTransferHistoryResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: userUniversalTransferHistoryEndpoint,
		SecType:  request.SecTypeSigned,
	}
	r.SetParam("type", s.transferType)
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	if s.fromSymbol != nil {
		r.SetParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.SetParam("toSymbol", *s.toSymbol)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// UserUniversalTransferHistoryResponse define response of UserUniversalTransferHistoryService
type UserUniversalTransferHistoryResponse struct {
	Total int64 `json:"total"`
	Rows  []struct {
		Asset     string `json:"asset"`
		Amount    string `json:"amount"`
		Type      string `json:"type"`
		Status    string `json:"status"`
		TranId    int64  `json:"tranId"`
		Timestamp uint64 `json:"timestamp"`
	} `json:"rows"`
}

// Funding Wallet (USER_DATA)
const (
	fundingWalletEndpoint = "/sapi/v1/asset/get-funding-asset"
)

// FundingWalletService funding wallet
type FundingWalletService struct {
	C                *connector.Connector
	asset            *string
	needBtcValuation *string
}

// Asset set asset
func (s *FundingWalletService) Asset(asset string) *FundingWalletService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *FundingWalletService) NeedBtcValuation(needBtcValuation string) *FundingWalletService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *FundingWalletService) Do(ctx context.Context) (res []*FundingWalletResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: fundingWalletEndpoint,
		SecType:  request.SecTypeSigned,
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.SetParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*FundingWalletResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// FundingWalletResponse define response of FundingWalletService
type FundingWalletResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

// User Asset (USER_DATA)
const (
	userAssetEndpoint = "/sapi/v3/asset/getUserAsset"
)

// UserAssetService user asset
type UserAssetService struct {
	C                *connector.Connector
	asset            *string
	needBtcValuation *bool
}

// Asset set asset
func (s *UserAssetService) Asset(asset string) *UserAssetService {
	s.asset = &asset
	return s
}

// NeedBtcValuation set needBtcValuation
func (s *UserAssetService) NeedBtcValuation(needBtcValuation bool) *UserAssetService {
	s.needBtcValuation = &needBtcValuation
	return s
}

func (s *UserAssetService) Do(ctx context.Context) (res []*UserAssetResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: userAssetEndpoint,
		SecType:  request.SecTypeSigned,
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.SetParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*UserAssetResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// UserAssetResponse define response of UserAssetService
type UserAssetResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}

// BUSD Convert (TRADE)
const (
	bUSDConvertEndpoint = "/sapi/v1/asset/convert-transfer"
)

// BUSDConvertService BUSD convert
type BUSDConvertService struct {
	C            *connector.Connector
	clientTranId string
	asset        string
	amount       float64
	targetAsset  string
	accountType  *string
}

// ClientTranId set clientTranId
func (s *BUSDConvertService) ClientTranId(clientTranId string) *BUSDConvertService {
	s.clientTranId = clientTranId
	return s
}

// Asset set asset
func (s *BUSDConvertService) Asset(asset string) *BUSDConvertService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *BUSDConvertService) Amount(amount float64) *BUSDConvertService {
	s.amount = amount
	return s
}

// TargetAsset set targetAsset
func (s *BUSDConvertService) TargetAsset(targetAsset string) *BUSDConvertService {
	s.targetAsset = targetAsset
	return s
}

// AccountType set accountType
func (s *BUSDConvertService) AccountType(accountType string) *BUSDConvertService {
	s.accountType = &accountType
	return s
}

func (s *BUSDConvertService) Do(ctx context.Context) (res *BUSDConvertResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: bUSDConvertEndpoint,
		SecType:  request.SecTypeSigned,
	}
	r.SetParam("clientTranId", s.clientTranId)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("targetAsset", s.targetAsset)
	if s.accountType != nil {
		r.SetParam("accountType", *s.accountType)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(BUSDConvertResponse)
	err = json.Unmarshal(data, res)
	return
}

// BUSDConvertResponse define response of BUSDConvertService
type BUSDConvertResponse struct {
	TranId int64  `json:"tranId"`
	Status string `json:"status"`
}

// BUSD Convert History (USER_DATA)
const (
	bUSDConvertHistoryEndpoint = "/sapi/v1/asset/convert-transfer/queryByPage"
)

// BUSDConvertHistoryService BUSD convert history
type BUSDConvertHistoryService struct {
	C            *connector.Connector
	tranId       *int64
	clientTranId *string
	asset        *string
	startTime    uint64
	endTime      uint64
	accountType  *string
	current      *int
	size         *int
}

// TranId set tranId
func (s *BUSDConvertHistoryService) TranId(tranId int64) *BUSDConvertHistoryService {
	s.tranId = &tranId
	return s
}

// ClientTranId set clientTranId
func (s *BUSDConvertHistoryService) ClientTranId(clientTranId string) *BUSDConvertHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// Asset set asset
func (s *BUSDConvertHistoryService) Asset(asset string) *BUSDConvertHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *BUSDConvertHistoryService) StartTime(startTime uint64) *BUSDConvertHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *BUSDConvertHistoryService) EndTime(endTime uint64) *BUSDConvertHistoryService {
	s.endTime = endTime
	return s
}

// AccountType set accountType
func (s *BUSDConvertHistoryService) AccountType(accountType string) *BUSDConvertHistoryService {
	s.accountType = &accountType
	return s
}

// Current set current
func (s *BUSDConvertHistoryService) Current(current int) *BUSDConvertHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *BUSDConvertHistoryService) Size(size int) *BUSDConvertHistoryService {
	s.size = &size
	return s
}

func (s *BUSDConvertHistoryService) Do(ctx context.Context) (res *BUSDConvertHistoryResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: bUSDConvertHistoryEndpoint,
		SecType:  request.SecTypeSigned,
	}
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	if s.tranId != nil {
		r.SetParam("tranId", *s.tranId)
	}
	if s.clientTranId != nil {
		r.SetParam("clientTranId", *s.clientTranId)
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.accountType != nil {
		r.SetParam("accountType", *s.accountType)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(BUSDConvertHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// BUSDConvertHistoryResponse define response of BUSDConvertHistoryService
type BUSDConvertHistoryResponse struct {
	Total int32 `json:"total"`
	Rows  []struct {
		TranId         int64  `json:"tranId"`
		Type           int32  `json:"type"`
		Time           uint64 `json:"time"`
		DeductedAsset  string `json:"deductedAsset"`
		DeductedAmount string `json:"deductedAmount"`
		TargetAsset    string `json:"targetAsset"`
		TargetAmount   string `json:"targetAmount"`
		Status         string `json:"status"`
		AccountType    string `json:"accountType"`
	} `json:"rows"`
}

// Get Cloud-Mining payment and refund history (USER_DATA)
const (
	cloudMiningPaymentHistoryEndpoint = "/sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage"
)

// CloudMiningPaymentHistoryService cloud mining payment history
type CloudMiningPaymentHistoryService struct {
	C            *connector.Connector
	tranid       *int64
	clientTranId *string
	asset        *string
	startTime    uint64
	endTime      uint64
	current      *int
	size         *int
}

// Tranid set tranid
func (s *CloudMiningPaymentHistoryService) Tranid(tranid int64) *CloudMiningPaymentHistoryService {
	s.tranid = &tranid
	return s
}

// ClientTranId set clientTranId
func (s *CloudMiningPaymentHistoryService) ClientTranId(clientTranId string) *CloudMiningPaymentHistoryService {
	s.clientTranId = &clientTranId
	return s
}

// Asset set asset
func (s *CloudMiningPaymentHistoryService) Asset(asset string) *CloudMiningPaymentHistoryService {
	s.asset = &asset
	return s
}

// StartTime set startTime
func (s *CloudMiningPaymentHistoryService) StartTime(startTime uint64) *CloudMiningPaymentHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *CloudMiningPaymentHistoryService) EndTime(endTime uint64) *CloudMiningPaymentHistoryService {
	s.endTime = endTime
	return s
}

// Current set current
func (s *CloudMiningPaymentHistoryService) Current(current int) *CloudMiningPaymentHistoryService {
	s.current = &current
	return s
}

// Size set size
func (s *CloudMiningPaymentHistoryService) Size(size int) *CloudMiningPaymentHistoryService {
	s.size = &size
	return s
}

func (s *CloudMiningPaymentHistoryService) Do(ctx context.Context) (res *CloudMiningPaymentHistoryResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: cloudMiningPaymentHistoryEndpoint,
		SecType:  request.SecTypeSigned,
	}
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	if s.tranid != nil {
		r.SetParam("tranId", *s.tranid)
	}
	if s.clientTranId != nil {
		r.SetParam("clientTranId", *s.clientTranId)
	}
	if s.asset != nil {
		r.SetParam("asset", *s.asset)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(CloudMiningPaymentHistoryResponse)
	err = json.Unmarshal(data, res)
	return
}

// CloudMiningPaymentHistoryResponse define response of CloudMiningPaymentHistoryService
type CloudMiningPaymentHistoryResponse struct {
	Total int32 `json:"total"`
	Rows  []struct {
		CreateTime uint64 `json:"createTime"`
		TranId     int64  `json:"tranId"`
		Type       int32  `json:"type"`
		Asset      string `json:"asset"`
		Amount     string `json:"amount"`
		Status     string `json:"status"`
	} `json:"rows"`
}

// Get API Key Permission (USER_DATA)
const (
	apiKeyPermissionEndpoint = "/sapi/v1/account/apiRestrictions"
)

// APIKeyPermissionService get api key permission
type APIKeyPermissionService struct {
	C *connector.Connector
}

func (s *APIKeyPermissionService) Do(ctx context.Context) (res *APIKeyPermissionResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: apiKeyPermissionEndpoint,
		SecType:  request.SecTypeSigned,
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(APIKeyPermissionResponse)
	err = json.Unmarshal(data, res)
	return
}

// APIKeyPermissionResponse define response of APIKeyPermissionService
type APIKeyPermissionResponse struct {
	IPRestrict                     bool   `json:"ipRestrict"`
	CreateTime                     uint64 `json:"createTime"`
	EnableWithdrawals              bool   `json:"enableWithdrawals"`
	EnableInternalTransfer         bool   `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool   `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool   `json:"enableVanillaOptions"`
	EnableReading                  bool   `json:"enableReading"`
	EnableFutures                  bool   `json:"enableFutures"`
	EnableMargin                   bool   `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool   `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime uint64 `json:"tradingAuthorityExpirationTime"`
}

// Query auto-converting stable coins (USER_DATA)
const (
	autoConvertStableCoinEndpoint = "/sapi/v1/capital/contract/convertible-coins"
)

// AutoConvertStableCoinService auto convert stable coin
type AutoConvertStableCoinService struct {
	C *connector.Connector
}

func (s *AutoConvertStableCoinService) Do(ctx context.Context) (res *AutoConvertStableCoinResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: autoConvertStableCoinEndpoint,
		SecType:  request.SecTypeSigned,
	}
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AutoConvertStableCoinResponse)
	err = json.Unmarshal(data, res)
	return
}

// AutoConvertStableCoinResponse define response of AutoConvertStableCoinService
type AutoConvertStableCoinResponse struct {
	ConvertEnabled bool `json:"convertEnabled"`
	Coins          []struct {
		Asset string `json:"coin"`
	} `json:"coins"`
	ExchangeRates []struct {
		Asset string `json:"coin"`
	} `json:"exchangeRates"`
}
