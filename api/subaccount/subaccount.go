package subaccount

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Create a Virtual Sub-account(For Master Account)
//
//gen:new_service
type CreateSubAccountService struct {
	C                *connector.Connector
	subAccountString string
}

func (s *CreateSubAccountService) SubAccountString(subAccountString string) *CreateSubAccountService {
	s.subAccountString = subAccountString
	return s
}

func (s *CreateSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *CreateSubAccountResp, err error) {
	r := request.New(
		"/sapi/v1/sub-account/virtualSubAccount",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.subAccountString == "" {
		err = fmt.Errorf("%w: subAccountString", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("subAccountString", s.subAccountString)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type CreateSubAccountResp struct {
	Email string `json:"email"`
}

// Query Sub-account List (For Master Account)
//
//gen:new_service
type QuerySubAccountListService struct {
	C        *connector.Connector
	email    *string
	isFreeze *string
	page     *int
	limit    *int
}

func (s *QuerySubAccountListService) Email(email string) *QuerySubAccountListService {
	s.email = &email
	return s
}

func (s *QuerySubAccountListService) IsFreeze(isFreeze string) *QuerySubAccountListService {
	s.isFreeze = &isFreeze
	return s
}

func (s *QuerySubAccountListService) Page(page int) *QuerySubAccountListService {
	s.page = &page
	return s
}

func (s *QuerySubAccountListService) Limit(limit int) *QuerySubAccountListService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountListService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubAccountListResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/list",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("email", s.email)
	r.SetParam("isFreeze", s.isFreeze)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountListResp)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountListResp struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

// Query Sub-account Spot Asset Transfer History (For Master Account)
//
//gen:new_service
type QuerySubAccountSpotAssetTransferHistoryService struct {
	C         *connector.Connector
	fromEmail *string
	toEmail   *string
	startTime *uint64
	endTime   *uint64
	page      *int
	limit     *int
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) FromEmail(fromEmail string) *QuerySubAccountSpotAssetTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) ToEmail(toEmail string) *QuerySubAccountSpotAssetTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) StartTime(startTime uint64) *QuerySubAccountSpotAssetTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) EndTime(endTime uint64) *QuerySubAccountSpotAssetTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Page(page int) *QuerySubAccountSpotAssetTransferHistoryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Limit(limit int) *QuerySubAccountSpotAssetTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountSpotAssetTransferHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res []*SubAccountTransferHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/sub-account/sub/transfer/history",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("fromEmail", s.fromEmail)
	r.SetParam("toEmail", s.toEmail)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountTransferHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountTransferHistoryResponse struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Qty    string `json:"qty"`
	Status string `json:"status"`
	TranId int64  `json:"tranId"`
	Time   uint64 `json:"time"`
}

// Query Sub-account Futures Asset Transfer History (For Master Account)

//gen:new_service
type QuerySubAccountFuturesAssetTransferHistoryService struct {
	C           *connector.Connector
	email       string
	futuresType int
	startTime   *uint64
	endTime     *uint64
	page        *int
	limit       *int
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Email(email string) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.email = email
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) FuturesType(futuresType int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.futuresType = futuresType
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) StartTime(startTime uint64) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) EndTime(endTime uint64) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Page(page int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Limit(limit int) *QuerySubAccountFuturesAssetTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QuerySubAccountFuturesAssetTransferHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *QuerySubAccountFuturesAssetTransferHistoryResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/internalTransfer",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.futuresType == 0 {
		err = fmt.Errorf("%w: futuresType", apierrors.ErrMissingParameter)
		return

	}

	r.SetParam("email", s.email)
	r.SetParam("futuresType", s.futuresType)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountFuturesAssetTransferHistoryResp)
	err = json.Unmarshal(data, &res)
	return
}

type QuerySubAccountFuturesAssetTransferHistoryResp struct {
	Success     bool                                `json:"success"`
	FuturesType int64                               `json:"futuresType"`
	Transfers   []SubAccountTransferHistoryResponse `json:"transfers"`
}

// Sub-account Futures Asset Transfer (For Master Account)

//gen:new_service
type SubAccountFuturesAssetTransferService struct {
	C           *connector.Connector
	fromEmail   string
	toEmail     string
	futuresType int64
	asset       string
	amount      float32
}

func (s *SubAccountFuturesAssetTransferService) FromEmail(fromEmail string) *SubAccountFuturesAssetTransferService {
	s.fromEmail = fromEmail
	return s
}

func (s *SubAccountFuturesAssetTransferService) ToEmail(toEmail string) *SubAccountFuturesAssetTransferService {
	s.toEmail = toEmail
	return s
}

func (s *SubAccountFuturesAssetTransferService) FuturesType(futuresType int64) *SubAccountFuturesAssetTransferService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFuturesAssetTransferService) Asset(asset string) *SubAccountFuturesAssetTransferService {
	s.asset = asset
	return s
}

func (s *SubAccountFuturesAssetTransferService) Amount(amount float32) *SubAccountFuturesAssetTransferService {
	s.amount = amount
	return s
}

func (s *SubAccountFuturesAssetTransferService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubAccountFuturesAssetTransferResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/internalTransfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("fromEmail", s.fromEmail)
	r.SetParam("toEmail", s.toEmail)
	r.SetParam("futuresType", s.futuresType)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAssetTransferResp)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountFuturesAssetTransferResp struct {
	Success bool   `json:"success"`
	TxnId   string `json:"txnId"`
}

// Query Sub-account Assets (For Master Account)

//gen:new_service
type QuerySubAccountAssetsService struct {
	C     *connector.Connector
	email string
}

func (s *QuerySubAccountAssetsService) Email(email string) *QuerySubAccountAssetsService {
	s.email = email
	return s
}

func (s *QuerySubAccountAssetsService) Do(ctx context.Context, opts ...request.RequestOption) (res *QuerySubAccountAssetsResp, err error) {

	r := request.New(
		"/sapi/v3/sub-account/assets",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountAssetsResp)
	err = json.Unmarshal(data, &res)
	return
}

type QuerySubAccountAssetsResp struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

//gen:new_service
type QuerySubAccountSpotAssetsSummaryService struct {
	C     *connector.Connector
	email *string
	page  *int
	size  *int
}

func (s *QuerySubAccountSpotAssetsSummaryService) Email(email string) *QuerySubAccountSpotAssetsSummaryService {
	s.email = &email
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Page(page int) *QuerySubAccountSpotAssetsSummaryService {
	s.page = &page
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Size(size int) *QuerySubAccountSpotAssetsSummaryService {
	s.size = &size
	return s
}

func (s *QuerySubAccountSpotAssetsSummaryService) Do(ctx context.Context, opts ...request.RequestOption) (res *QuerySubAccountSpotAssetsSummaryResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/spotSummary",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("email", s.email)
	r.SetParam("page", s.page)
	r.SetParam("size", s.size)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountSpotAssetsSummaryResp)
	err = json.Unmarshal(data, &res)
	return
}

type QuerySubAccountSpotAssetsSummaryResp struct {
	TotalCount                int64  `json:"totalCount"`
	MasterAccountTotalAsset   string `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []struct {
		Email   string `json:"email"`
		ToAsset string `json:"toAsset"`
	} `json:"spotSubUserAssetBtcVoList"`
}

//gen:new_service
type SubAccountDepositAddressService struct {
	C       *connector.Connector
	email   string
	coin    string
	network *string
}

func (s *SubAccountDepositAddressService) Email(email string) *SubAccountDepositAddressService {
	s.email = email
	return s
}

func (s *SubAccountDepositAddressService) Coin(coin string) *SubAccountDepositAddressService {
	s.coin = coin
	return s
}

func (s *SubAccountDepositAddressService) Network(network string) *SubAccountDepositAddressService {
	s.network = &network
	return s
}

func (s *SubAccountDepositAddressService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubAccountDepositAddressResp, err error) {

	r := request.New(
		"/sapi/v1/capital/deposit/subAddress",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	if s.coin == "" {
		err = fmt.Errorf("%w: coin", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("coin", s.coin)
	r.SetParam("network", s.network)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountDepositAddressResp)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountDepositAddressResp struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

// Get Sub-account Deposit History (For Master Account)

//gen:new_service
type SubAccountDepositHistoryService struct {
	C         *connector.Connector
	email     string
	coin      string
	status    *int64
	startTime *uint64
	endTime   *uint64
	limit     *int
	offset    *int64
}

func (s *SubAccountDepositHistoryService) Email(email string) *SubAccountDepositHistoryService {
	s.email = email
	return s
}

func (s *SubAccountDepositHistoryService) Coin(coin string) *SubAccountDepositHistoryService {
	s.coin = coin
	return s
}

func (s *SubAccountDepositHistoryService) Status(status int64) *SubAccountDepositHistoryService {
	s.status = &status
	return s
}

func (s *SubAccountDepositHistoryService) StartTime(startTime uint64) *SubAccountDepositHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountDepositHistoryService) EndTime(endTime uint64) *SubAccountDepositHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountDepositHistoryService) Limit(limit int) *SubAccountDepositHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountDepositHistoryService) Offset(offset int64) *SubAccountDepositHistoryService {
	s.offset = &offset
	return s
}

func (s *SubAccountDepositHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res []*SubAccountDepositHistoryResponse, err error) {

	r := request.New(
		"/sapi/v1/capital/deposit/subHisrec",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.coin == "" {
		err = fmt.Errorf("%w: coin", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("coin", s.coin)

	r.SetParam("status", s.status)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)
	r.SetParam("offset", s.offset)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*SubAccountDepositHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountDepositHistoryResponse struct {
	Id            int64  `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int64  `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    uint64 `json:"insertTime"`
	TransferType  int64  `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int64  `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}

// Get Sub-account's Status on Margin/Futures (For Master Account)

//gen:new_service
type SubAccountStatusService struct {
	C     *connector.Connector
	email *string
}

func (s *SubAccountStatusService) Email(email string) *SubAccountStatusService {
	s.email = &email
	return s
}

func (s *SubAccountStatusService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubAccountStatusResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/status",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("email", s.email)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountStatusResp)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountStatusResp struct {
	Email            string `json:"email"`
	IsSubUserEnabled bool   `json:"isSubUserEnabled"`
	IsUserActive     bool   `json:"isUserActive"`
	InsertTime       uint64 `json:"insertTime"`
	IsMarginEnabled  bool   `json:"isMarginEnabled"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
	Mobile           int64  `json:"mobile"`
}

// Enable Margin for Sub-account (For Master Account)

//gen:new_service
type EnableMarginForSubAccountService struct {
	C     *connector.Connector
	email string
}

func (s *EnableMarginForSubAccountService) Email(email string) *EnableMarginForSubAccountService {
	s.email = email
	return s
}

func (s *EnableMarginForSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *EnableMarginForSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/margin/enable",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableMarginForSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type EnableMarginForSubAccountResp struct {
	Email           string `json:"email"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
}

// Get Detail on Sub-account's Margin Account (For Master Account)

//gen:new_service
type DetailOnSubAccountMarginAccountService struct {
	C     *connector.Connector
	email string
}

func (s *DetailOnSubAccountMarginAccountService) Email(email string) *DetailOnSubAccountMarginAccountService {
	s.email = email
	return s
}

func (s *DetailOnSubAccountMarginAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *DetailOnSubAccountMarginAccountResp, err error) {
	r := request.New(
		"/sapi/v1/sub-account/margin/account",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DetailOnSubAccountMarginAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type DetailOnSubAccountMarginAccountResp struct {
	Email               string `json:"email"`
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo  struct {
		ForceLiquidationRate string `json:"forceLiquidationRate"`
		MarginCallBar        string `json:"marginCallBar"`
		NormalBar            string `json:"normalBar"`
	} `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	}
}

// Get Summary of Sub-account's Margin Account (For Master Account)

//gen:new_service
type SummaryOfSubAccountMarginAccountService struct {
	C *connector.Connector
}

func (s *SummaryOfSubAccountMarginAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *SummaryOfSubAccountMarginAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/margin/accountSummary",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SummaryOfSubAccountMarginAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type SummaryOfSubAccountMarginAccountResp struct {
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	SubAccountList      []struct {
		Email               string `json:"email"`
		TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
		TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
		TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	} `json:"subAccountList"`
}

// Enable Futures for Sub-account (For Master Account)

//gen:new_service
type EnableFuturesForSubAccountService struct {
	C     *connector.Connector
	email string
}

func (s *EnableFuturesForSubAccountService) Email(email string) *EnableFuturesForSubAccountService {
	s.email = email
	return s
}

func (s *EnableFuturesForSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *EnableFuturesForSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/enable",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesForSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type EnableFuturesForSubAccountResp struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

// Get Detail on Sub-account's Futures Account (For Master Account)

//gen:new_service
type DetailOnSubAccountFuturesAccountService struct {
	C     *connector.Connector
	email string
}

func (s *DetailOnSubAccountFuturesAccountService) Email(email string) *DetailOnSubAccountFuturesAccountService {
	s.email = email
	return s
}

func (s *DetailOnSubAccountFuturesAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *DetailOnSubAccountFuturesAccountResp, err error) {
	r := request.New(
		"/sapi/v1/sub-account/futures/account",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DetailOnSubAccountFuturesAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type DetailOnSubAccountFuturesAccountResp struct {
	Email  string `json:"email"`
	Asset  string `json:"asset"`
	Assets []struct {
		Asset                  string `json:"asset"`
		InitialMargin          string `json:"initialMargin"`
		MaintenanceMargin      string `json:"maintenanceMargin"`
		MarginBalance          string `json:"marginBalance"`
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		WalletBalance          string `json:"walletBalance"`
	} `json:"assets"`
	CanDeposit                  bool   `json:"canDeposit"`
	CanTrade                    bool   `json:"canTrade"`
	CanWithdraw                 bool   `json:"canWithdraw"`
	FeeTier                     int    `json:"feeTier"`
	MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	UpdateTime                  uint64 `json:"updateTime"`
}

// Get Summary of Sub-account's Futures Account (For Master Account)

//gen:new_service
type SummaryOfSubAccountFuturesAccountService struct {
	C *connector.Connector
}

func (s *SummaryOfSubAccountFuturesAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *SummaryOfSubAccountFuturesAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/accountSummary",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SummaryOfSubAccountFuturesAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type SummaryOfSubAccountFuturesAccountResp struct {
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
	Asset                       string `json:"asset"`
	SubAccountList              []struct {
		Email                       string `json:"email"`
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		Asset                       string `json:"asset"`
	} `json:"subAccountList"`
}

// Get Futures Position-Risk of Sub-account (For Master Account)

//gen:new_service
type FuturesPositionRiskOfSubAccountService struct {
	C     *connector.Connector
	email string
}

func (s *FuturesPositionRiskOfSubAccountService) Email(email string) *FuturesPositionRiskOfSubAccountService {
	s.email = email
	return s
}

func (s *FuturesPositionRiskOfSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *FuturesPositionRiskOfSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/positionRisk",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesPositionRiskOfSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type FuturesPositionRiskOfSubAccountResp struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MaxNotional      string `json:"maxNotional"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmount   string `json:"positionAmount"`
	Symbol           string `json:"symbol"`
	UnrealizedProfit string `json:"unrealizedProfit"`
}

// Futures Transfer for Sub-account (For Master Account)

//gen:new_service
type FuturesTransferForSubAccountService struct {
	C            *connector.Connector
	email        string
	asset        string
	amount       float64
	transferType int
}

func (s *FuturesTransferForSubAccountService) Email(email string) *FuturesTransferForSubAccountService {
	s.email = email
	return s
}

func (s *FuturesTransferForSubAccountService) Asset(asset string) *FuturesTransferForSubAccountService {
	s.asset = asset
	return s
}

func (s *FuturesTransferForSubAccountService) Amount(amount float64) *FuturesTransferForSubAccountService {
	s.amount = amount
	return s
}

func (s *FuturesTransferForSubAccountService) TransferType(transferType int) *FuturesTransferForSubAccountService {
	s.transferType = transferType
	return s
}

func (s *FuturesTransferForSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *FuturesTransferForSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/transfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}

	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	if s.transferType == 0 {
		err = fmt.Errorf("%w: transferType", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("type", s.transferType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesTransferForSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type FuturesTransferForSubAccountResp struct {
	TxnId int `json:"txnId"`
}

// Margin Transfer for Sub-account (For Master Account)

//gen:new_service
type MarginTransferForSubAccountService struct {
	C            *connector.Connector
	email        string
	asset        string
	amount       float32
	transferType int
}

func (s *MarginTransferForSubAccountService) Email(email string) *MarginTransferForSubAccountService {
	s.email = email
	return s
}

func (s *MarginTransferForSubAccountService) Asset(asset string) *MarginTransferForSubAccountService {
	s.asset = asset
	return s
}

func (s *MarginTransferForSubAccountService) Amount(amount float32) *MarginTransferForSubAccountService {
	s.amount = amount
	return s
}

func (s *MarginTransferForSubAccountService) TransferType(transferType int) *MarginTransferForSubAccountService {
	s.transferType = transferType
	return s
}

func (s *MarginTransferForSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *MarginTransferForSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/margin/transfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}
	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	if s.transferType == 0 {
		err = fmt.Errorf("%w: transferType", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("type", s.transferType)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginTransferForSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type MarginTransferForSubAccountResp struct {
	TxnId int `json:"txnId"`
}

// Transfer to Sub-account of Same Master (For Sub-account)

//gen:new_service
type TransferToSubAccountOfSameMasterService struct {
	C       *connector.Connector
	toEmail string
	asset   string
	amount  float64
}

func (s *TransferToSubAccountOfSameMasterService) ToEmail(toEmail string) *TransferToSubAccountOfSameMasterService {
	s.toEmail = toEmail
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Asset(asset string) *TransferToSubAccountOfSameMasterService {
	s.asset = asset
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Amount(amount float64) *TransferToSubAccountOfSameMasterService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountOfSameMasterService) Do(ctx context.Context, opts ...request.RequestOption) (res *TransferToSubAccountOfSameMasterResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/transfer/subToSub",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.toEmail == "" {
		err = fmt.Errorf("%w: toEmail", apierrors.ErrMissingParameter)
		return
	}
	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}
	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("toEmail", s.toEmail)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransferToSubAccountOfSameMasterResp)
	err = json.Unmarshal(data, &res)
	return
}

type TransferToSubAccountOfSameMasterResp struct {
	TxnId int `json:"txnId"`
}

// Transfer to Master (For Sub-account)

//gen:new_service
type TransferToMasterService struct {
	C      *connector.Connector
	asset  string
	amount float64
}

func (s *TransferToMasterService) Asset(asset string) *TransferToMasterService {
	s.asset = asset
	return s
}

func (s *TransferToMasterService) Amount(amount float64) *TransferToMasterService {
	s.amount = amount
	return s
}

func (s *TransferToMasterService) Do(ctx context.Context, opts ...request.RequestOption) (res *TransferToMasterResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/transfer/subToMaster",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)
	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}
	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransferToMasterResp)
	err = json.Unmarshal(data, &res)
	return
}

type TransferToMasterResp struct {
	TxnId int `json:"txnId"`
}

// Sub-account Transfer History (For Sub-account)

//gen:new_service
type SubAccountTransferHistoryService struct {
	C            *connector.Connector
	asset        *string
	transferType *int
	startTime    *uint64
	endTime      *uint64
	limit        *int
}

func (s *SubAccountTransferHistoryService) Asset(asset string) *SubAccountTransferHistoryService {
	s.asset = &asset
	return s
}

func (s *SubAccountTransferHistoryService) TransferType(transferType int) *SubAccountTransferHistoryService {
	s.transferType = &transferType
	return s
}

func (s *SubAccountTransferHistoryService) StartTime(startTime uint64) *SubAccountTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SubAccountTransferHistoryService) EndTime(endTime uint64) *SubAccountTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SubAccountTransferHistoryService) Limit(limit int) *SubAccountTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *SubAccountTransferHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res *SubAccountTransferHistoryResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/transfer/subUserHistory",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("type", s.transferType)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountTransferHistoryResp)
	err = json.Unmarshal(data, &res)
	return
}

type SubAccountTransferHistoryResp struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int    `json:"type"`
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	TranId          int    `json:"tranId"`
	Time            uint64 `json:"time"`
}

// Universal Transfer (For Master Account)

//gen:new_service
type UniversalTransferService struct {
	C               *connector.Connector
	asset           string
	amount          float64
	fromAccountType string
	toAccountType   string
	fromEmail       *string
	toEmail         *string
	clientTranId    *string
	symbol          *string
}

func (s *UniversalTransferService) FromEmail(fromEmail string) *UniversalTransferService {
	s.fromEmail = &fromEmail
	return s
}

func (s *UniversalTransferService) ToEmail(toEmail string) *UniversalTransferService {
	s.toEmail = &toEmail
	return s
}

func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}

func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}

func (s *UniversalTransferService) ClientTranId(clientTranId string) *UniversalTransferService {
	s.clientTranId = &clientTranId
	return s
}

func (s *UniversalTransferService) Symbol(symbol string) *UniversalTransferService {
	s.symbol = &symbol
	return s
}

func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}

func (s *UniversalTransferService) Amount(amount float64) *UniversalTransferService {
	s.amount = amount
	return s
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...request.RequestOption) (res *UniversalTransferResp, err error) {

	r := request.New(
		"/sapi/v1/asset/universalTransfer",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("fromAccountType", s.fromAccountType)
	r.SetParam("toAccountType", s.toAccountType)

	r.SetParam("fromEmail", s.fromEmail)
	r.SetParam("toEmail", s.toEmail)
	r.SetParam("clientTranId", s.clientTranId)
	r.SetParam("symbol", s.symbol)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UniversalTransferResp)
	err = json.Unmarshal(data, &res)
	return
}

type UniversalTransferResp struct {
	TranId       int    `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

// Query Universal Transfer History (For Master Account)

//gen:new_service
type QueryUniversalTransferHistoryService struct {
	C            *connector.Connector
	fromEmail    *string
	toEmail      *string
	clientTranId *string
	startTime    *uint64
	endTime      *uint64
	page         *int
	limit        *int
}

func (s *QueryUniversalTransferHistoryService) FromEmail(fromEmail string) *QueryUniversalTransferHistoryService {
	s.fromEmail = &fromEmail
	return s
}

func (s *QueryUniversalTransferHistoryService) ToEmail(toEmail string) *QueryUniversalTransferHistoryService {
	s.toEmail = &toEmail
	return s
}

func (s *QueryUniversalTransferHistoryService) ClientTranId(clientTranId string) *QueryUniversalTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}

func (s *QueryUniversalTransferHistoryService) StartTime(startTime uint64) *QueryUniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}

func (s *QueryUniversalTransferHistoryService) EndTime(endTime uint64) *QueryUniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}

func (s *QueryUniversalTransferHistoryService) Page(page int) *QueryUniversalTransferHistoryService {
	s.page = &page
	return s
}

func (s *QueryUniversalTransferHistoryService) Limit(limit int) *QueryUniversalTransferHistoryService {
	s.limit = &limit
	return s
}

func (s *QueryUniversalTransferHistoryService) Do(ctx context.Context, opts ...request.RequestOption) (res QueryUniversalTransferHistoryResp, err error) {
	r := request.New(
		"/sapi/v1/asset/universalTransfer",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("fromEmail", s.fromEmail)
	r.SetParam("toEmail", s.toEmail)
	r.SetParam("clientTranId", s.clientTranId)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res.Result = make([]*InternalUniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	return
}

type QueryUniversalTransferHistoryResp struct {
	Result     []*InternalUniversalTransfer `json:"result"`
	TotalCount int                          `json:"totalCount"`
}

type InternalUniversalTransfer struct {
	TranId          int64  `json:"tranId"`
	ClientTranId    string `json:"clientTranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	CreateTimeStamp uint64 `json:"createTimeStamp"`
}

// Get Detail on Sub-account's Futures Account V2 (For Master Account)

//gen:new_service
type DetailOnSubAccountFuturesAccountV2Service struct {
	C           *connector.Connector
	email       string
	futuresType int
}

func (s *DetailOnSubAccountFuturesAccountV2Service) Email(email string) *DetailOnSubAccountFuturesAccountV2Service {
	s.email = email
	return s
}

func (s *DetailOnSubAccountFuturesAccountV2Service) FuturesType(futuresType int) *DetailOnSubAccountFuturesAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *DetailOnSubAccountFuturesAccountV2Service) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {
	r := request.New(
		"/sapi/v1/sub-account/futures/account",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.futuresType == 0 {
		err = fmt.Errorf("%w: futuresType", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("futuresType", s.futuresType)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(DetailOnSubAccountFuturesAccountV2USDTResp)
	} else {
		res = new(DetailOnSubAccountFuturesAccountV2COINResp)
	}
	err = json.Unmarshal(data, &res)
	return
}

type FuturesAccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

type DetailOnSubAccountFuturesAccountV2USDTResp struct {
	FutureAccountResp struct {
		Email                       string                `json:"email"`
		Assets                      []FuturesAccountAsset `json:"assets"`
		CanDeposit                  bool                  `json:"canDeposit"`
		CanTrade                    bool                  `json:"canTrade"`
		CanWithdraw                 bool                  `json:"canWithdraw"`
		FeeTier                     int                   `json:"feeTier"`
		MaxWithdrawAmount           string                `json:"maxWithdrawAmount"`
		TotalInitialMargin          string                `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string                `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string                `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string                `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string                `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string                `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string                `json:"totalWalletBalance"`
		UpdateTime                  uint64                `json:"updateTime"`
	} `json:"futureAccountResp"`
}

type DetailOnSubAccountFuturesAccountV2COINResp struct {
	DeliveryAccountResp struct {
		Email       string                `json:"email"`
		Assets      []FuturesAccountAsset `json:"assets"`
		CanDeposit  bool                  `json:"canDeposit"`
		CanTrade    bool                  `json:"canTrade"`
		CanWithdraw bool                  `json:"canWithdraw"`
		FeeTier     int                   `json:"feeTier"`
		UpdateTime  uint64                `json:"updateTime"`
	} `json:"deliveryAccountResp"`
}

// Get Summary of Sub-account's Futures Account V2 (For Master Account)

//gen:new_service
type SummaryOfSubAccountFuturesAccountV2Service struct {
	C           *connector.Connector
	futuresType int
	page        *int
	limit       *int
}

func (s *SummaryOfSubAccountFuturesAccountV2Service) FuturesType(futuresType int) *SummaryOfSubAccountFuturesAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *SummaryOfSubAccountFuturesAccountV2Service) Page(page int) *SummaryOfSubAccountFuturesAccountV2Service {
	s.page = &page
	return s
}

func (s *SummaryOfSubAccountFuturesAccountV2Service) Limit(limit int) *SummaryOfSubAccountFuturesAccountV2Service {
	s.limit = &limit
	return s
}

func (s *SummaryOfSubAccountFuturesAccountV2Service) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/accountSummary",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("futuresType", s.futuresType)

	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(SummaryOfSubAccountFuturesAccountV2USDTResp)
	} else {
		res = new(SummaryOfSubAccountFuturesAccountV2COINResp)
	}
	err = json.Unmarshal(data, &res)
	return
}

type SummaryOfSubAccountFuturesAccountV2USDTResp struct {
	FutureAccountSummaryResp struct {
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		Asset                       string `json:"asset"`
		SubAccountList              []struct {
			Email                       string `json:"email"`
			TotalInitialMargin          string `json:"totalInitialMargin"`
			TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
			TotalMarginBalance          string `json:"totalMarginBalance"`
			TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
			TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
			TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
			TotalWalletBalance          string `json:"totalWalletBalance"`
			Asset                       string `json:"asset"`
		} `json:"subAccountList"`
	} `json:"futureAccountSummaryResp"`
}

type SummaryOfSubAccountFuturesAccountV2COINResp struct {
	DeliveryAccountSummaryResp struct {
		TotalMarginBalanceOfBTC    string `json:"totalMarginBalanceOfBTC"`
		TotalUnrealizedProfitOfBTC string `json:"totalUnrealizedProfitOfBTC"`
		TotalWalletBalanceOfBTC    string `json:"totalWalletBalanceOfBTC"`
		Asset                      string `json:"asset"`
		SubAccountList             []struct {
			Email                 string `json:"email"`
			TotalMarginBalance    string `json:"totalMarginBalance"`
			TotalUnrealizedProfit string `json:"totalUnrealizedProfit"`
			TotalWalletBalance    string `json:"totalWalletBalance"`
			Asset                 string `json:"asset"`
		} `json:"subAccountList"`
	} `json:"deliveryAccountSummaryResp"`
}

// Get Futures Position-Risk of Sub-account V2 (For Master Account)

//gen:new_service
type FuturesPositionRiskOfSubAccountV2Service struct {
	C           *connector.Connector
	email       string
	futuresType int
}

func (s *FuturesPositionRiskOfSubAccountV2Service) Email(email string) *FuturesPositionRiskOfSubAccountV2Service {
	s.email = email
	return s
}

func (s *FuturesPositionRiskOfSubAccountV2Service) FuturesType(futuresType int) *FuturesPositionRiskOfSubAccountV2Service {
	s.futuresType = futuresType
	return s
}

func (s *FuturesPositionRiskOfSubAccountV2Service) Do(ctx context.Context, opts ...request.RequestOption) (res interface{}, err error) {

	r := request.New(
		"/sapi/v1/sub-account/futures/positionRisk",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.futuresType == 0 {
		err = fmt.Errorf("%w: futuresType", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("futuresType", s.futuresType)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.futuresType == 1 {
		res = new(FuturesPositionRiskOfSubAccountV2USDTResp)
	} else {
		res = new(FuturesPositionRiskOfSubAccountV2COINResp)
	}
	err = json.Unmarshal(data, &res)
	return
}

type FuturesPositionRiskOfSubAccountV2USDTResp struct {
	FuturePositionRiskVos []struct {
		EntryPrice       string `json:"entryPrice"`
		Leverage         string `json:"leverage"`
		MaxNotional      string `json:"maxNotional"`
		LiquidationPrice string `json:"liquidationPrice"`
		MarkPrice        string `json:"markPrice"`
		PositionAmount   string `json:"positionAmount"`
		Symbol           string `json:"symbol"`
		UnrealizedProfit string `json:"unrealizedProfit"`
	} `json:"futurePositionRiskVos"`
}

type FuturesPositionRiskOfSubAccountV2COINResp struct {
	DeliveryPositionRiskVos []struct {
		EntryPrice       string `json:"entryPrice"`
		MarkPrice        string `json:"markPrice"`
		Leverage         string `json:"leverage"`
		Isolated         string `json:"isolated"`
		IsolatedWallet   string `json:"isolatedWallet"`
		IsolatedMargin   string `json:"isolatedMargin"`
		IsAutoAddMargin  string `json:"isAutoAddMargin"`
		PositionSide     string `json:"positionSide"`
		PositionAmount   string `json:"positionAmount"`
		Symbol           string `json:"symbol"`
		UnrealizedProfit string `json:"unrealizedProfit"`
	} `json:"deliveryPositionRiskVos"`
}

// Enable Leverage Token for Sub-account (For Master Account)
const (
	enableLeverageTokenForSubAccountEndpoint = "/sapi/v1/sub-account/blvt/enable"
)

//gen:new_service
type EnableLeverageTokenForSubAccountService struct {
	C          *connector.Connector
	email      string
	enableBlvt bool
}

func (s *EnableLeverageTokenForSubAccountService) Email(email string) *EnableLeverageTokenForSubAccountService {
	s.email = email
	return s
}

func (s *EnableLeverageTokenForSubAccountService) EnableBlvt(enableBlvt bool) *EnableLeverageTokenForSubAccountService {
	s.enableBlvt = enableBlvt
	return s
}

func (s *EnableLeverageTokenForSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *EnableLeverageTokenForSubAccountResp, err error) {

	r := request.New(
		enableLeverageTokenForSubAccountEndpoint,
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	// TODO: this is weird
	if !s.enableBlvt {
		err = fmt.Errorf("%w: enableBlvt", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)
	r.SetParam("enableBlvt", s.enableBlvt)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableLeverageTokenForSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type EnableLeverageTokenForSubAccountResp struct {
	Email      string `json:"email"`
	EnableBlvt bool   `json:"enableBlvt"`
}

// Get IP Restriction for a Sub-account API Key (For Master Account)

//gen:new_service
type IPRestrictionForSubAccountAPIKeyService struct {
	C                *connector.Connector
	email            string
	subAccountApiKey string
}

func (s *IPRestrictionForSubAccountAPIKeyService) Email(email string) *IPRestrictionForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *IPRestrictionForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *IPRestrictionForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *IPRestrictionForSubAccountAPIKeyService) Do(ctx context.Context, opts ...request.RequestOption) (res *IPRestrictionForSubAccountAPIKeyResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/subaccountApi/ipRestriction",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	if s.subAccountApiKey == "" {
		err = fmt.Errorf("%w: subAccountApiKey", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("subAccountApiKey", s.subAccountApiKey)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IPRestrictionForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, &res)
	return
}

type IPRestrictionForSubAccountAPIKeyResp struct {
	IpRestrict string `json:"ipRestrict"`
	IpList     []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Delete IP List For a Sub-account API Key (For Master Account)

//gen:new_service
type DeleteIPListForSubAccountAPIKeyService struct {
	C                *connector.Connector
	email            string
	subAccountApiKey string
	ipAddress        *string
}

func (s *DeleteIPListForSubAccountAPIKeyService) Email(email string) *DeleteIPListForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteIPListForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) IpAddress(ipAddress string) *DeleteIPListForSubAccountAPIKeyService {
	s.ipAddress = &ipAddress
	return s
}

func (s *DeleteIPListForSubAccountAPIKeyService) Do(ctx context.Context, opts ...request.RequestOption) (res *DeleteIPListForSubAccountAPIKeyResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/subaccountApi/ipRestriction/ipList",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return

	}
	if s.subAccountApiKey == "" {
		err = fmt.Errorf("%w: subAccountApiKey", apierrors.ErrMissingParameter)
		return

	}

	r.SetParam("email", s.email)
	r.SetParam("subAccountApiKey", s.subAccountApiKey)
	r.SetParam("ipAddress", s.ipAddress)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DeleteIPListForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, &res)
	return
}

type DeleteIPListForSubAccountAPIKeyResp struct {
	IpRestrict string `json:"ipRestrict"`
	IpList     []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Update IP Restriction for Sub-Account API key (For Master Account)

//gen:new_service
type UpdateIPRestrictionForSubAccountAPIKeyService struct {
	C                *connector.Connector
	email            string
	subAccountApiKey string
	status           string
	ipAddress        *string
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Email(email string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.email = email
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Status(status string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.status = status
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) IpAddress(ipAddress string) *UpdateIPRestrictionForSubAccountAPIKeyService {
	s.ipAddress = &ipAddress
	return s
}

func (s *UpdateIPRestrictionForSubAccountAPIKeyService) Do(ctx context.Context, opts ...request.RequestOption) (res *UpdateIPRestrictionForSubAccountAPIKeyResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/subaccountApi/ipRestriction",
		request.Method(http.MethodPut),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	if s.subAccountApiKey == "" {
		err = fmt.Errorf("%w: subAccountApiKey", apierrors.ErrMissingParameter)
		return
	}
	if s.status == "" {
		err = fmt.Errorf("%w: status", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)
	r.SetParam("subAccountApiKey", s.subAccountApiKey)
	r.SetParam("status", s.status)

	r.SetParam("ipAddress", s.ipAddress)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UpdateIPRestrictionForSubAccountAPIKeyResp)
	err = json.Unmarshal(data, &res)
	return
}

type UpdateIPRestrictionForSubAccountAPIKeyResp struct {
	Status string `json:"status"`
	IpList []struct {
		Ip string `json:"ip"`
	} `json:"ipList"`
	UpdateTime uint64 `json:"updateTime"`
	ApiKey     string `json:"apiKey"`
}

// Deposit Assets Into The Managed Sub-account（For Investor Master Account）

//gen:new_service
type DepositAssetsIntoTheManagedSubAccountService struct {
	C       *connector.Connector
	toEmail string
	asset   string
	amount  float64
}

func (s *DepositAssetsIntoTheManagedSubAccountService) ToEmail(toEmail string) *DepositAssetsIntoTheManagedSubAccountService {
	s.toEmail = toEmail
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Asset(asset string) *DepositAssetsIntoTheManagedSubAccountService {
	s.asset = asset
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Amount(amount float64) *DepositAssetsIntoTheManagedSubAccountService {
	s.amount = amount
	return s
}

func (s *DepositAssetsIntoTheManagedSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *DepositAssetsIntoTheManagedSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/deposit",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)
	if s.toEmail == "" {
		err = fmt.Errorf("%w: toEmail", apierrors.ErrMissingParameter)
		return

	}
	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}

	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return

	}
	r.SetParam("toEmail", s.toEmail)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DepositAssetsIntoTheManagedSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type DepositAssetsIntoTheManagedSubAccountResp struct {
	TranId int64 `json:"tranId"`
}

// Query Managed Sub-account Asset Details（For Investor Master Account）

//gen:new_service
type QueryManagedSubAccountAssetDetailsService struct {
	C     *connector.Connector
	email string
}

func (s *QueryManagedSubAccountAssetDetailsService) Email(email string) *QueryManagedSubAccountAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountAssetDetailsService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountAssetDetailsResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/managed-subaccount/asset",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountAssetDetailsResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountAssetDetailsResp struct {
	AssetDetail []struct {
		Coin             string `json:"coin"`
		Name             string `json:"name"`
		TotalBalance     string `json:"totalBalance"`
		AvailableBalance string `json:"availableBalance"`
		InOrder          string `json:"inOrder"`
		BtcValue         string `json:"btcValue"`
	} `json:"assetDetail"`
}

// Withdrawl Assets From The Managed Sub-account（For Investor Master Account）

//gen:new_service
type WithdrawAssetsFromTheManagedSubAccountService struct {
	C            *connector.Connector
	fromEmail    string
	asset        string
	amount       float32
	transferDate *int64
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) FromEmail(fromEmail string) *WithdrawAssetsFromTheManagedSubAccountService {
	s.fromEmail = fromEmail
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Asset(asset string) *WithdrawAssetsFromTheManagedSubAccountService {
	s.asset = asset
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Amount(amount float32) *WithdrawAssetsFromTheManagedSubAccountService {
	s.amount = amount
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) TransferDate(transferDate int64) *WithdrawAssetsFromTheManagedSubAccountService {
	s.transferDate = &transferDate
	return s
}

func (s *WithdrawAssetsFromTheManagedSubAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *WithdrawAssetsFromTheManagedSubAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/managed-subaccount/withdraw",
		request.Method(http.MethodPost),
		request.SecType(request.SecTypeSigned),
	)

	if s.fromEmail == "" {
		err = fmt.Errorf("%w: fromEmail", apierrors.ErrMissingParameter)
		return

	}
	if s.asset == "" {
		err = fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
		return
	}

	if s.amount == 0 {
		err = fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("fromEmail", s.fromEmail)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)
	r.SetParam("transferDate", s.transferDate)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(WithdrawAssetsFromTheManagedSubAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type WithdrawAssetsFromTheManagedSubAccountResp struct {
	TranId int64 `json:"tranId"`
}

// Query Managed Sub-account Snapshot（For Investor Master Account）

//gen:new_service
type QueryManagedSubAccountSnapshotService struct {
	C         *connector.Connector
	email     string
	subType   string
	startTime *uint64
	endTime   *uint64
	limit     *int
}

func (s *QueryManagedSubAccountSnapshotService) Email(email string) *QueryManagedSubAccountSnapshotService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountSnapshotService) SubType(subType string) *QueryManagedSubAccountSnapshotService {
	s.subType = subType
	return s
}

func (s *QueryManagedSubAccountSnapshotService) StartTime(startTime uint64) *QueryManagedSubAccountSnapshotService {
	s.startTime = &startTime
	return s
}

func (s *QueryManagedSubAccountSnapshotService) EndTime(endTime uint64) *QueryManagedSubAccountSnapshotService {
	s.endTime = &endTime
	return s
}

func (s *QueryManagedSubAccountSnapshotService) Limit(limit int) *QueryManagedSubAccountSnapshotService {
	s.limit = &limit
	return s
}

func (s *QueryManagedSubAccountSnapshotService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountSnapshotResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/accountSnapshot",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return

	}
	if s.subType == "" {
		err = fmt.Errorf("%w: subType", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("type", s.subType)

	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountSnapshotResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountSnapshotResp struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data []struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime uint64 `json:"updateTime"`
	} `json:"snapshotVos"`
}

// Query Managed Sub Account Transfer Log (Investor) (USER_DATA)

//gen:new_service
type QueryManagedSubAccountTransferLogService struct {
	C                           *connector.Connector
	email                       string
	startTime                   uint64
	endTime                     uint64
	page                        int
	limit                       int
	transfers                   *string
	transferFunctionAccountType *string
}

func (s *QueryManagedSubAccountTransferLogService) Email(email string) *QueryManagedSubAccountTransferLogService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountTransferLogService) StartTime(startTime uint64) *QueryManagedSubAccountTransferLogService {
	s.startTime = startTime
	return s
}

func (s *QueryManagedSubAccountTransferLogService) EndTime(endTime uint64) *QueryManagedSubAccountTransferLogService {
	s.endTime = endTime
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Page(page int) *QueryManagedSubAccountTransferLogService {
	s.page = page
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Limit(limit int) *QueryManagedSubAccountTransferLogService {
	s.limit = limit
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Transfers(transfers string) *QueryManagedSubAccountTransferLogService {
	s.transfers = &transfers
	return s
}

func (s *QueryManagedSubAccountTransferLogService) TransferFunctionAccountType(transferFunctionAccountType string) *QueryManagedSubAccountTransferLogService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *QueryManagedSubAccountTransferLogService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountTransferLogResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/queryTransLogForInvestor",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	if s.startTime == 0 {
		err = fmt.Errorf("%w: startTime", apierrors.ErrMissingParameter)
		return

	}
	if s.endTime == 0 {
		err = fmt.Errorf("%w: endTime", apierrors.ErrMissingParameter)
		return

	}
	if s.page == 0 {
		err = fmt.Errorf("%w: page", apierrors.ErrMissingParameter)
		return

	}
	if s.limit == 0 {
		err = fmt.Errorf("%w: limit", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	r.SetParam("transfers", s.transfers)
	r.SetParam("transferFunctionAccountType", s.transferFunctionAccountType)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountTransferLogResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountTransferLogResp struct {
	ManagerSubTransferHistoryVos []struct {
		FromEmail       string `json:"fromEmail"`
		FromAccountType string `json:"fromAccountType"`
		ToEmail         string `json:"toEmail"`
		ToAccountType   string `json:"toAccountType"`
		Asset           string `json:"asset"`
		Amount          int    `json:"amount"`
		ScheduledData   int    `json:"scheduledData"`
		CreateTime      uint64 `json:"createTime"`
		Status          string `json:"status"`
	} `json:"managerSubTransferHistoryVos"`
	Count int `json:"count"`
}

// Query Managed Sub-account Futures Asset Details（For Investor Master Account）(USER_DATA)

//gen:new_service
type QueryManagedSubAccountFuturesAssetDetailsService struct {
	C     *connector.Connector
	email string
}

func (s *QueryManagedSubAccountFuturesAssetDetailsService) Email(email string) *QueryManagedSubAccountFuturesAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountFuturesAssetDetailsService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountFuturesAssetDetailsResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/fetch-future-asset",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return

	}
	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountFuturesAssetDetailsResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountFuturesAssetDetailsResp struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	SnapshotVos []struct {
		Type       string `json:"type"`
		UpdateTime uint64 `json:"updateTime"`
		Data       struct {
			Assets []struct {
				Asset         string `json:"asset"`
				MarginBalance int64  `json:"marginBalance"`
				WalletBalance int64  `json:"walletBalance"`
			} `json:"assets"`
			Position []struct {
				Symbol      string `json:"symbol"`
				EntryPrice  int64  `json:"entryPrice"`
				MarkPrice   int64  `json:"markPrice"`
				PositionAmt int64  `json:"positionAmt"`
			} `json:"position"`
		} `json:"data"`
	} `json:"snapshotVos"`
}

// Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)

//gen:new_service
type QueryManagedSubAccountMarginAssetDetailsService struct {
	C     *connector.Connector
	email string
}

func (s *QueryManagedSubAccountMarginAssetDetailsService) Email(email string) *QueryManagedSubAccountMarginAssetDetailsService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountMarginAssetDetailsService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountMarginAssetDetailsResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/marginAsset",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountMarginAssetDetailsResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountMarginAssetDetailsResp struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
	UserAssets          []struct {
		Asset    string `json:"asset"`
		Borrowed string `json:"borrowed"`
		Free     string `json:"free"`
		Interest string `json:"interest"`
		Locked   string `json:"locked"`
		NetAsset string `json:"netAsset"`
	} `json:"userAssets"`
}

// Query Managed Sub Account Transfer Log (Trading Team) (USER_DATA)

//gen:new_service
type QueryManagedSubAccountTransferLogForTradingTeamService struct {
	C                           *connector.Connector
	email                       string
	startTime                   uint64
	endTime                     uint64
	page                        int
	limit                       int
	transfers                   *string
	transferFunctionAccountType *string
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Email(email string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.email = email
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) StartTime(startTime uint64) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.startTime = startTime
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) EndTime(endTime uint64) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.endTime = endTime
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Page(page int) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.page = page
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Limit(limit int) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.limit = limit
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Transfers(transfers string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.transfers = &transfers
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) TransferFunctionAccountType(transferFunctionAccountType string) *QueryManagedSubAccountTransferLogForTradingTeamService {
	s.transferFunctionAccountType = &transferFunctionAccountType
	return s
}

func (s *QueryManagedSubAccountTransferLogForTradingTeamService) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountTransferLogForTradingTeamResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/queryTransLogForTradeParent",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	if s.startTime == 0 {
		err = fmt.Errorf("%w: startTime", apierrors.ErrMissingParameter)
		return
	}
	if s.endTime == 0 {
		err = fmt.Errorf("%w: endTime", apierrors.ErrMissingParameter)
		return
	}
	if s.page == 0 {
		err = fmt.Errorf("%w: page", apierrors.ErrMissingParameter)
		return
	}

	if s.limit == 0 {
		err = fmt.Errorf("%w: limit", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	r.SetParam("transfers", s.transfers)
	r.SetParam("transferFunctionAccountType", s.transferFunctionAccountType)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountTransferLogForTradingTeamResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountTransferLogForTradingTeamResp struct {
	ManagerSubTransferHistoryVos []struct {
		FromEmail       string `json:"fromEmail"`
		FromAccountType string `json:"fromAccountType"`
		ToEmail         string `json:"toEmail"`
		ToAccountType   string `json:"toAccountType"`
		Asset           string `json:"asset"`
		Amount          string `json:"amount"`
		ScheduledData   int64  `json:"scheduledData"`
		CreateTime      uint64 `json:"createTime"`
		Status          string `json:"status"`
	} `json:"managerSubTransferHistoryVos"`
	Count int `json:"count"`
}

// Query Sub-account Assets (For Master Account)(USER_DATA)

//gen:new_service
type QuerySubAccountAssetsForMasterAccountService struct {
	C     *connector.Connector
	email string
}

func (s *QuerySubAccountAssetsForMasterAccountService) Email(email string) *QuerySubAccountAssetsForMasterAccountService {
	s.email = email
	return s
}

func (s *QuerySubAccountAssetsForMasterAccountService) Do(ctx context.Context, opts ...request.RequestOption) (res *QuerySubAccountAssetsForMasterAccountResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/assets",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountAssetsForMasterAccountResp)
	err = json.Unmarshal(data, &res)
	return
}

type QuerySubAccountAssetsForMasterAccountResp struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

// Query Managed Sub-account List

//gen:new_service
type QueryManagedSubAccountList struct {
	C     *connector.Connector
	email *string
	page  *int
	limit *int
}

func (s *QueryManagedSubAccountList) Email(email string) *QueryManagedSubAccountList {
	s.email = &email
	return s
}

func (s *QueryManagedSubAccountList) Page(page int) *QueryManagedSubAccountList {
	s.page = &page
	return s
}

func (s *QueryManagedSubAccountList) Limit(limit int) *QueryManagedSubAccountList {
	s.limit = &limit
	return s
}

func (s *QueryManagedSubAccountList) Do(ctx context.Context, opts ...request.RequestOption) (res *QueryManagedSubAccountListResp, err error) {
	r := request.New(
		"/sapi/v1/managed-subaccount/info",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("email", s.email)
	r.SetParam("page", s.page)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QueryManagedSubAccountListResp)
	err = json.Unmarshal(data, &res)
	return
}

type QueryManagedSubAccountListResp struct {
	Total                    int `json:"total"`
	ManagerSubUserInfoVoList []struct {
		RootUserId               int    `json:"rootUserId"`
		ManagersubUserId         int    `json:"managersubUserId"`
		BindParentUserId         int    `json:"bindParentUserId"`
		Email                    string `json:"email"`
		InsertTimestamp          uint64 `json:"insertTimestamp"`
		BindParentEmail          string `json:"bindParentEmail"`
		IsSubUserEnabled         bool   `json:"isSubUserEnabled"`
		IsUserActive             bool   `json:"isUserActive"`
		IsMarginEnabled          bool   `json:"isMarginEnabled"`
		IsFutureEnabled          bool   `json:"isFutureEnabled"`
		IsSignedLVTRiskAgreement bool   `json:"isSignedLVTRiskAgreement"`
	} `json:"managerSubUserInfoVoList"`
}

// Query Sub-account Transaction Tatistics (For Master Account) (USER_DATA)

//gen:new_service
type QuerySubAccountTransactionStatistics struct {
	C     *connector.Connector
	email string
}

func (s *QuerySubAccountTransactionStatistics) Email(email string) *QuerySubAccountTransactionStatistics {
	s.email = email
	return s
}

func (s *QuerySubAccountTransactionStatistics) Do(ctx context.Context, opts ...request.RequestOption) (res *QuerySubAccountTransactionStatisticsResp, err error) {

	r := request.New(
		"/sapi/v1/sub-account/transaction-statistics",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return
	}
	r.SetParam("email", s.email)
	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(QuerySubAccountTransactionStatisticsResp)
	err = json.Unmarshal(data, &res)
	return
}

type QuerySubAccountTransactionStatisticsResp struct {
	Recent30BtcTotal         string `json:"recent30BtcTotal"`
	Recent30BtcFuturesTotal  string `json:"recent30BtcFuturesTotal"`
	Recent30BtcMarginTotal   string `json:"recent30BtcMarginTotal"`
	Recent30BusdTotal        string `json:"recent30BusdTotal"`
	Recent30BusdFuturesTotal string `json:"recent30BusdFuturesTotal"`
	Recent30BusdMarginTotal  string `json:"recent30BusdMarginTotal"`
	TradeInfoVos             []struct {
		UserId      int64 `json:"userId"`
		Btc         int   `json:"btc"`
		BtcFutures  int   `json:"btcFutures"`
		BtcMargin   int   `json:"btcMargin"`
		Busd        int   `json:"busd"`
		BusdFutures int   `json:"busdFutures"`
		BusdMargin  int   `json:"busdMargin"`
		Date        int64 `json:"date"`
	} `json:"tradeInfoVos"`
}

// Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)

//gen:new_service
type ManagedSubAccountDepositAddressService struct {
	C       *connector.Connector
	email   string
	coin    string
	network *string
}

func (s *ManagedSubAccountDepositAddressService) Email(email string) *ManagedSubAccountDepositAddressService {
	s.email = email
	return s
}

func (s *ManagedSubAccountDepositAddressService) Coin(coin string) *ManagedSubAccountDepositAddressService {
	s.coin = coin
	return s
}

func (s *ManagedSubAccountDepositAddressService) Network(network string) *ManagedSubAccountDepositAddressService {
	s.network = &network
	return s
}

func (s *ManagedSubAccountDepositAddressService) Do(ctx context.Context, opts ...request.RequestOption) (res *ManagedSubAccountDepositAddressResp, err error) {

	r := request.New(
		"/sapi/v1/managed-subaccount/deposit/address",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)
	if s.email == "" {
		err = fmt.Errorf("%w: email", apierrors.ErrMissingParameter)
		return

	}
	if s.coin == "" {
		err = fmt.Errorf("%w: coin", apierrors.ErrMissingParameter)
		return

	}

	r.SetParam("email", s.email)
	r.SetParam("coin", s.coin)
	r.SetParam("network", s.network)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ManagedSubAccountDepositAddressResp)
	err = json.Unmarshal(data, &res)
	return
}

type ManagedSubAccountDepositAddressResp struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
