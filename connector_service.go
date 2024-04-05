
package binance_connector


func (c *Client) NewAccountService() *AccountService {
	return &AccountService{C: c.Connector}
}


func (c *Client) NewAccountCommissionService() *AccountCommissionService {
	return &AccountCommissionService{C: c.Connector}
}


func (c *Client) NewQueryCurrentOrderCountUsageService() *QueryCurrentOrderCountUsageService {
	return &QueryCurrentOrderCountUsageService{C: c.Connector}
}


func (c *Client) NewAccountAllocationsService() *AccountAllocationsService {
	return &AccountAllocationsService{C: c.Connector}
}


func (c *Client) NewQueryPreventedMatchesService() *QueryPreventedMatchesService {
	return &QueryPreventedMatchesService{C: c.Connector}
}


func (c *Client) NewAccountTradeListService() *AccountTradeListService {
	return &AccountTradeListService{C: c.Connector}
}


func (c *Client) NewAggTradesListService() *AggTradesList {
	return &AggTradesList{C: c.Connector}
}


func (c *Client) NewAvgPriceService() *AvgPrice {
	return &AvgPrice{C: c.Connector}
}


func (c *Client) NewTickerBookTickerService() *TickerBookTicker {
	return &TickerBookTicker{C: c.Connector}
}


func (c *Client) NewExchangeInfoService() *ExchangeInfo {
	return &ExchangeInfo{C: c.Connector}
}


func (c *Client) NewHistoricalTradeLookupService() *HistoricalTradeLookup {
	return &HistoricalTradeLookup{C: c.Connector}
}


func (c *Client) NewKlinesService() *Klines {
	return &Klines{C: c.Connector}
}


func (c *Client) NewOrderBookService() *OrderBook {
	return &OrderBook{C: c.Connector}
}


func (c *Client) NewPingService() *Ping {
	return &Ping{C: c.Connector}
}


func (c *Client) NewRecentTradesListService() *RecentTradesList {
	return &RecentTradesList{C: c.Connector}
}


func (c *Client) NewServerTimeService() *ServerTime {
	return &ServerTime{C: c.Connector}
}


func (c *Client) NewTickerService() *Ticker {
	return &Ticker{C: c.Connector}
}


func (c *Client) NewTicker24hrService() *Ticker24hr {
	return &Ticker24hr{C: c.Connector}
}


func (c *Client) NewTickerPriceService() *TickerPrice {
	return &TickerPrice{C: c.Connector}
}


func (c *Client) NewUiKlinesService() *UiKlines {
	return &UiKlines{C: c.Connector}
}


func (c *Client) NewCancelOCOService() *CancelOCOService {
	return &CancelOCOService{C: c.Connector}
}


func (c *Client) NewCancelOpenOrdersService() *CancelOpenOrdersService {
	return &CancelOpenOrdersService{C: c.Connector}
}


func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{C: c.Connector}
}


func (c *Client) NewCancelReplaceService() *CancelReplaceService {
	return &CancelReplaceService{C: c.Connector}
}


func (c *Client) NewCreateOrderService() *CreateOrderService {
	return &CreateOrderService{C: c.Connector}
}


func (c *Client) NewGetAllOrdersService() *GetAllOrdersService {
	return &GetAllOrdersService{C: c.Connector}
}


func (c *Client) NewGetOpenOrdersService() *GetOpenOrdersService {
	return &GetOpenOrdersService{C: c.Connector}
}


func (c *Client) NewGetOrderService() *GetOrderService {
	return &GetOrderService{C: c.Connector}
}


func (c *Client) NewNewOCOService() *NewOCOService {
	return &NewOCOService{C: c.Connector}
}


func (c *Client) NewCreateOrderListOCOService() *CreateOrderListOCOService {
	return &CreateOrderListOCOService{C: c.Connector}
}


func (c *Client) NewQueryAllOCOService() *QueryAllOCOService {
	return &QueryAllOCOService{C: c.Connector}
}


func (c *Client) NewQueryOCOService() *QueryOCOService {
	return &QueryOCOService{C: c.Connector}
}


func (c *Client) NewQueryOpenOCOService() *QueryOpenOCOService {
	return &QueryOpenOCOService{C: c.Connector}
}


func (c *Client) NewTestNewOrderService() *TestNewOrder {
	return &TestNewOrder{C: c.Connector}
}


func (c *Client) NewAccountStatusService() *AccountStatusService {
	return &AccountStatusService{C: c.Connector}
}


func (c *Client) NewAccountApiTradingStatusService() *AccountApiTradingStatusService {
	return &AccountApiTradingStatusService{C: c.Connector}
}


func (c *Client) NewAllCoinsInfoService() *AllCoinsInfoService {
	return &AllCoinsInfoService{C: c.Connector}
}


func (c *Client) NewAPIKeyPermissionService() *APIKeyPermissionService {
	return &APIKeyPermissionService{C: c.Connector}
}


func (c *Client) NewAssetDetailV2Service() *AssetDetailV2Service {
	return &AssetDetailV2Service{C: c.Connector}
}


func (c *Client) NewAssetDividendRecordService() *AssetDividendRecordService {
	return &AssetDividendRecordService{C: c.Connector}
}


func (c *Client) NewAutoConvertStableCoinService() *AutoConvertStableCoinService {
	return &AutoConvertStableCoinService{C: c.Connector}
}


func (c *Client) NewCloudMiningPaymentHistoryService() *CloudMiningPaymentHistoryService {
	return &CloudMiningPaymentHistoryService{C: c.Connector}
}


func (c *Client) NewAccountSnapshotService() *AccountSnapshotService {
	return &AccountSnapshotService{C: c.Connector}
}


func (c *Client) NewDepositAddressService() *DepositAddressService {
	return &DepositAddressService{C: c.Connector}
}


func (c *Client) NewDepositHistoryService() *DepositHistoryService {
	return &DepositHistoryService{C: c.Connector}
}


func (c *Client) NewDisableFastWithdrawSwitchService() *DisableFastWithdrawSwitchService {
	return &DisableFastWithdrawSwitchService{C: c.Connector}
}


func (c *Client) NewAssetDetailService() *AssetDetailService {
	return &AssetDetailService{C: c.Connector}
}


func (c *Client) NewDustLogService() *DustLogService {
	return &DustLogService{C: c.Connector}
}


func (c *Client) NewDustTransferService() *DustTransferService {
	return &DustTransferService{C: c.Connector}
}


func (c *Client) NewEnableFastWithdrawSwitchService() *EnableFastWithdrawSwitchService {
	return &EnableFastWithdrawSwitchService{C: c.Connector}
}


func (c *Client) NewFundingWalletService() *FundingWalletService {
	return &FundingWalletService{C: c.Connector}
}


func (c *Client) NewSystemStatusService() *SystemStatusService {
	return &SystemStatusService{C: c.Connector}
}


func (c *Client) NewTradeFeeService() *TradeFeeService {
	return &TradeFeeService{C: c.Connector}
}


func (c *Client) NewUserUniversalTransferService() *UserUniversalTransferService {
	return &UserUniversalTransferService{C: c.Connector}
}


func (c *Client) NewUserUniversalTransferHistoryService() *UserUniversalTransferHistoryService {
	return &UserUniversalTransferHistoryService{C: c.Connector}
}


func (c *Client) NewUserAssetService() *UserAssetService {
	return &UserAssetService{C: c.Connector}
}


func (c *Client) NewWithdrawService() *WithdrawService {
	return &WithdrawService{C: c.Connector}
}


func (c *Client) NewWithdrawHistoryService() *WithdrawHistoryService {
	return &WithdrawHistoryService{C: c.Connector}
}


func (c *Client) NewPersonalLeftQuotaService() *PersonalLeftQuotaService {
	return &PersonalLeftQuotaService{C: c.Connector}
}


func (c *Client) NewPurchaseStakingProductService() *PurchaseStakingProductService {
	return &PurchaseStakingProductService{C: c.Connector}
}


func (c *Client) NewRedeemStakingProductService() *RedeemStakingProductService {
	return &RedeemStakingProductService{C: c.Connector}
}


func (c *Client) NewSetAutoStakingService() *SetAutoStakingService {
	return &SetAutoStakingService{C: c.Connector}
}


func (c *Client) NewStakingHistoryService() *StakingHistoryService {
	return &StakingHistoryService{C: c.Connector}
}


func (c *Client) NewStakingProductListService() *StakingProductListService {
	return &StakingProductListService{C: c.Connector}
}


func (c *Client) NewStakingProductPositionService() *StakingProductPositionService {
	return &StakingProductPositionService{C: c.Connector}
}


func (c *Client) NewCreateSubAccountService() *CreateSubAccountService {
	return &CreateSubAccountService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountListService() *QuerySubAccountListService {
	return &QuerySubAccountListService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountSpotAssetTransferHistoryService() *QuerySubAccountSpotAssetTransferHistoryService {
	return &QuerySubAccountSpotAssetTransferHistoryService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountFuturesAssetTransferHistoryService() *QuerySubAccountFuturesAssetTransferHistoryService {
	return &QuerySubAccountFuturesAssetTransferHistoryService{C: c.Connector}
}


func (c *Client) NewSubAccountFuturesAssetTransferService() *SubAccountFuturesAssetTransferService {
	return &SubAccountFuturesAssetTransferService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountAssetsService() *QuerySubAccountAssetsService {
	return &QuerySubAccountAssetsService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountSpotAssetsSummaryService() *QuerySubAccountSpotAssetsSummaryService {
	return &QuerySubAccountSpotAssetsSummaryService{C: c.Connector}
}


func (c *Client) NewSubAccountDepositAddressService() *SubAccountDepositAddressService {
	return &SubAccountDepositAddressService{C: c.Connector}
}


func (c *Client) NewSubAccountDepositHistoryService() *SubAccountDepositHistoryService {
	return &SubAccountDepositHistoryService{C: c.Connector}
}


func (c *Client) NewSubAccountStatusService() *SubAccountStatusService {
	return &SubAccountStatusService{C: c.Connector}
}


func (c *Client) NewEnableMarginForSubAccountService() *EnableMarginForSubAccountService {
	return &EnableMarginForSubAccountService{C: c.Connector}
}


func (c *Client) NewDetailOnSubAccountMarginAccountService() *DetailOnSubAccountMarginAccountService {
	return &DetailOnSubAccountMarginAccountService{C: c.Connector}
}


func (c *Client) NewSummaryOfSubAccountMarginAccountService() *SummaryOfSubAccountMarginAccountService {
	return &SummaryOfSubAccountMarginAccountService{C: c.Connector}
}


func (c *Client) NewEnableFuturesForSubAccountService() *EnableFuturesForSubAccountService {
	return &EnableFuturesForSubAccountService{C: c.Connector}
}


func (c *Client) NewDetailOnSubAccountFuturesAccountService() *DetailOnSubAccountFuturesAccountService {
	return &DetailOnSubAccountFuturesAccountService{C: c.Connector}
}


func (c *Client) NewSummaryOfSubAccountFuturesAccountService() *SummaryOfSubAccountFuturesAccountService {
	return &SummaryOfSubAccountFuturesAccountService{C: c.Connector}
}


func (c *Client) NewFuturesPositionRiskOfSubAccountService() *FuturesPositionRiskOfSubAccountService {
	return &FuturesPositionRiskOfSubAccountService{C: c.Connector}
}


func (c *Client) NewFuturesTransferForSubAccountService() *FuturesTransferForSubAccountService {
	return &FuturesTransferForSubAccountService{C: c.Connector}
}


func (c *Client) NewMarginTransferForSubAccountService() *MarginTransferForSubAccountService {
	return &MarginTransferForSubAccountService{C: c.Connector}
}


func (c *Client) NewTransferToSubAccountOfSameMasterService() *TransferToSubAccountOfSameMasterService {
	return &TransferToSubAccountOfSameMasterService{C: c.Connector}
}


func (c *Client) NewTransferToMasterService() *TransferToMasterService {
	return &TransferToMasterService{C: c.Connector}
}


func (c *Client) NewSubAccountTransferHistoryService() *SubAccountTransferHistoryService {
	return &SubAccountTransferHistoryService{C: c.Connector}
}


func (c *Client) NewUniversalTransferService() *UniversalTransferService {
	return &UniversalTransferService{C: c.Connector}
}


func (c *Client) NewQueryUniversalTransferHistoryService() *QueryUniversalTransferHistoryService {
	return &QueryUniversalTransferHistoryService{C: c.Connector}
}


func (c *Client) NewDetailOnSubAccountFuturesAccountV2Service() *DetailOnSubAccountFuturesAccountV2Service {
	return &DetailOnSubAccountFuturesAccountV2Service{C: c.Connector}
}


func (c *Client) NewSummaryOfSubAccountFuturesAccountV2Service() *SummaryOfSubAccountFuturesAccountV2Service {
	return &SummaryOfSubAccountFuturesAccountV2Service{C: c.Connector}
}


func (c *Client) NewFuturesPositionRiskOfSubAccountV2Service() *FuturesPositionRiskOfSubAccountV2Service {
	return &FuturesPositionRiskOfSubAccountV2Service{C: c.Connector}
}


func (c *Client) NewEnableLeverageTokenForSubAccountService() *EnableLeverageTokenForSubAccountService {
	return &EnableLeverageTokenForSubAccountService{C: c.Connector}
}


func (c *Client) NewIPRestrictionForSubAccountAPIKeyService() *IPRestrictionForSubAccountAPIKeyService {
	return &IPRestrictionForSubAccountAPIKeyService{C: c.Connector}
}


func (c *Client) NewDeleteIPListForSubAccountAPIKeyService() *DeleteIPListForSubAccountAPIKeyService {
	return &DeleteIPListForSubAccountAPIKeyService{C: c.Connector}
}


func (c *Client) NewUpdateIPRestrictionForSubAccountAPIKeyService() *UpdateIPRestrictionForSubAccountAPIKeyService {
	return &UpdateIPRestrictionForSubAccountAPIKeyService{C: c.Connector}
}


func (c *Client) NewDepositAssetsIntoTheManagedSubAccountService() *DepositAssetsIntoTheManagedSubAccountService {
	return &DepositAssetsIntoTheManagedSubAccountService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountAssetDetailsService() *QueryManagedSubAccountAssetDetailsService {
	return &QueryManagedSubAccountAssetDetailsService{C: c.Connector}
}


func (c *Client) NewWithdrawAssetsFromTheManagedSubAccountService() *WithdrawAssetsFromTheManagedSubAccountService {
	return &WithdrawAssetsFromTheManagedSubAccountService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountSnapshotService() *QueryManagedSubAccountSnapshotService {
	return &QueryManagedSubAccountSnapshotService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountTransferLogService() *QueryManagedSubAccountTransferLogService {
	return &QueryManagedSubAccountTransferLogService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountFuturesAssetDetailsService() *QueryManagedSubAccountFuturesAssetDetailsService {
	return &QueryManagedSubAccountFuturesAssetDetailsService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountMarginAssetDetailsService() *QueryManagedSubAccountMarginAssetDetailsService {
	return &QueryManagedSubAccountMarginAssetDetailsService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountTransferLogForTradingTeamService() *QueryManagedSubAccountTransferLogForTradingTeamService {
	return &QueryManagedSubAccountTransferLogForTradingTeamService{C: c.Connector}
}


func (c *Client) NewQuerySubAccountAssetsForMasterAccountService() *QuerySubAccountAssetsForMasterAccountService {
	return &QuerySubAccountAssetsForMasterAccountService{C: c.Connector}
}


func (c *Client) NewQueryManagedSubAccountListService() *QueryManagedSubAccountList {
	return &QueryManagedSubAccountList{C: c.Connector}
}


func (c *Client) NewQuerySubAccountTransactionStatisticsService() *QuerySubAccountTransactionStatistics {
	return &QuerySubAccountTransactionStatistics{C: c.Connector}
}


func (c *Client) NewManagedSubAccountDepositAddressService() *ManagedSubAccountDepositAddressService {
	return &ManagedSubAccountDepositAddressService{C: c.Connector}
}


func (c *Client) NewTransferService() *TransferService {
	return &TransferService{C: c.Connector}
}


func (c *Client) NewBorrowService() *BorrowService {
	return &BorrowService{C: c.Connector}
}


func (c *Client) NewRepayService() *RepayService {
	return &RepayService{C: c.Connector}
}


func (c *Client) NewQueryMarginAssetService() *QueryMarginAssetService {
	return &QueryMarginAssetService{C: c.Connector}
}


func (c *Client) NewQueryCrossMarginPairService() *QueryCrossMarginPairService {
	return &QueryCrossMarginPairService{C: c.Connector}
}


func (c *Client) NewGetAllMarginAssetsService() *GetAllMarginAssetsService {
	return &GetAllMarginAssetsService{C: c.Connector}
}


func (c *Client) NewGetAllMarginPairsService() *GetAllMarginPairsService {
	return &GetAllMarginPairsService{C: c.Connector}
}


func (c *Client) NewQueryMarginPriceIndexService() *QueryMarginPriceIndexService {
	return &QueryMarginPriceIndexService{C: c.Connector}
}


func (c *Client) NewMarginAccountNewOrderService() *MarginAccountNewOrderService {
	return &MarginAccountNewOrderService{C: c.Connector}
}


func (c *Client) NewMarginAccountCancelOrderService() *MarginAccountCancelOrderService {
	return &MarginAccountCancelOrderService{C: c.Connector}
}


func (c *Client) NewMarginAccountCancelAllOrdersService() *MarginAccountCancelAllOrdersService {
	return &MarginAccountCancelAllOrdersService{C: c.Connector}
}


func (c *Client) NewQueryMarginBorrowRepayService() *QueryMarginBorrowRepayService {
	return &QueryMarginBorrowRepayService{C: c.Connector}
}


func (c *Client) NewInterestHistoryService() *InterestHistoryService {
	return &InterestHistoryService{C: c.Connector}
}


func (c *Client) NewForceLiquidationRecordService() *ForceLiquidationRecordService {
	return &ForceLiquidationRecordService{C: c.Connector}
}


func (c *Client) NewCrossMarginAccountDetailService() *CrossMarginAccountDetailService {
	return &CrossMarginAccountDetailService{C: c.Connector}
}


func (c *Client) NewMarginAccountOrderService() *MarginAccountOrderService {
	return &MarginAccountOrderService{C: c.Connector}
}


func (c *Client) NewMarginAccountOpenOrderService() *MarginAccountOpenOrderService {
	return &MarginAccountOpenOrderService{C: c.Connector}
}


func (c *Client) NewMarginAccountAllOrderService() *MarginAccountAllOrderService {
	return &MarginAccountAllOrderService{C: c.Connector}
}


func (c *Client) NewMarginAccountNewOCOService() *MarginAccountNewOCOService {
	return &MarginAccountNewOCOService{C: c.Connector}
}


func (c *Client) NewMarginAccountCancelOCOService() *MarginAccountCancelOCOService {
	return &MarginAccountCancelOCOService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryOCOService() *MarginAccountQueryOCOService {
	return &MarginAccountQueryOCOService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryAllOCOService() *MarginAccountQueryAllOCOService {
	return &MarginAccountQueryAllOCOService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryOpenOCOService() *MarginAccountQueryOpenOCOService {
	return &MarginAccountQueryOpenOCOService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryTradeListService() *MarginAccountQueryTradeListService {
	return &MarginAccountQueryTradeListService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryMaxBorrowService() *MarginAccountQueryMaxBorrowService {
	return &MarginAccountQueryMaxBorrowService{C: c.Connector}
}


func (c *Client) NewMarginAccountQueryMaxTransferOutAmountService() *MarginAccountQueryMaxTransferOutAmountService {
	return &MarginAccountQueryMaxTransferOutAmountService{C: c.Connector}
}


func (c *Client) NewMarginAccountSummaryService() *MarginAccountSummaryService {
	return &MarginAccountSummaryService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountTransferService() *MarginIsolatedAccountTransferService {
	return &MarginIsolatedAccountTransferService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountTransferHistoryService() *MarginIsolatedAccountTransferHistoryService {
	return &MarginIsolatedAccountTransferHistoryService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountInfoService() *MarginIsolatedAccountInfoService {
	return &MarginIsolatedAccountInfoService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountDisableService() *MarginIsolatedAccountDisableService {
	return &MarginIsolatedAccountDisableService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountEnableService() *MarginIsolatedAccountEnableService {
	return &MarginIsolatedAccountEnableService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedAccountLimitService() *MarginIsolatedAccountLimitService {
	return &MarginIsolatedAccountLimitService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedSymbolService() *MarginIsolatedSymbolService {
	return &MarginIsolatedSymbolService{C: c.Connector}
}


func (c *Client) NewAllIsolatedMarginSymbolService() *AllIsolatedMarginSymbolService {
	return &AllIsolatedMarginSymbolService{C: c.Connector}
}


func (c *Client) NewMarginToggleBnbBurnService() *MarginToggleBnbBurnService {
	return &MarginToggleBnbBurnService{C: c.Connector}
}


func (c *Client) NewMarginBnbBurnStatusService() *MarginBnbBurnStatusService {
	return &MarginBnbBurnStatusService{C: c.Connector}
}


func (c *Client) NewMarginInterestRateHistoryService() *MarginInterestRateHistoryService {
	return &MarginInterestRateHistoryService{C: c.Connector}
}


func (c *Client) NewMarginCrossMarginFeeService() *MarginCrossMarginFeeService {
	return &MarginCrossMarginFeeService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedMarginFeeService() *MarginIsolatedMarginFeeService {
	return &MarginIsolatedMarginFeeService{C: c.Connector}
}


func (c *Client) NewMarginIsolatedMarginTierService() *MarginIsolatedMarginTierService {
	return &MarginIsolatedMarginTierService{C: c.Connector}
}


func (c *Client) NewMarginCurrentOrderCountService() *MarginCurrentOrderCountService {
	return &MarginCurrentOrderCountService{C: c.Connector}
}


func (c *Client) NewMarginDustlogService() *MarginDustlogService {
	return &MarginDustlogService{C: c.Connector}
}


func (c *Client) NewMarginCrossCollateralRatioService() *MarginCrossCollateralRatioService {
	return &MarginCrossCollateralRatioService{C: c.Connector}
}


func (c *Client) NewMarginSmallLiabilityExchangeCoinListService() *MarginSmallLiabilityExchangeCoinListService {
	return &MarginSmallLiabilityExchangeCoinListService{C: c.Connector}
}


func (c *Client) NewMarginSmallLiabilityExchangeService() *MarginSmallLiabilityExchangeService {
	return &MarginSmallLiabilityExchangeService{C: c.Connector}
}


func (c *Client) NewMarginSmallLiabilityExchangeHistoryService() *MarginSmallLiabilityExchangeHistoryService {
	return &MarginSmallLiabilityExchangeHistoryService{C: c.Connector}
}


func (c *Client) NewBETHRewardsDistributionHistoryService() *BETHRewardsDistributionHistoryService {
	return &BETHRewardsDistributionHistoryService{C: c.Connector}
}


func (c *Client) NewCurrentETHStakingQuotaService() *CurrentETHStakingQuotaService {
	return &CurrentETHStakingQuotaService{C: c.Connector}
}


func (c *Client) NewRedeemETHService() *RedeemETHService {
	return &RedeemETHService{C: c.Connector}
}


func (c *Client) NewETHRedemptionHistoryService() *ETHRedemptionHistoryService {
	return &ETHRedemptionHistoryService{C: c.Connector}
}


func (c *Client) NewETHStakingAccountService() *ETHStakingAccountService {
	return &ETHStakingAccountService{C: c.Connector}
}


func (c *Client) NewETHStakingHistoryService() *ETHStakingHistoryService {
	return &ETHStakingHistoryService{C: c.Connector}
}


func (c *Client) NewSubscribeETHStakingService() *SubscribeETHStakingService {
	return &SubscribeETHStakingService{C: c.Connector}
}


func (c *Client) NewSubscribeETHStakingV2Service() *SubscribeETHStakingV2Service {
	return &SubscribeETHStakingV2Service{C: c.Connector}
}


func (c *Client) NewWBETHRateHistoryService() *WBETHRateHistoryService {
	return &WBETHRateHistoryService{C: c.Connector}
}


func (c *Client) NewWBETHRewardsHistoryService() *WBETHRewardsHistoryService {
	return &WBETHRewardsHistoryService{C: c.Connector}
}


func (c *Client) NewWBETHUnwrapHistoryService() *WBETHUnwrapHistoryService {
	return &WBETHUnwrapHistoryService{C: c.Connector}
}


func (c *Client) NewWBETHWrapHistoryService() *WBETHWrapHistoryService {
	return &WBETHWrapHistoryService{C: c.Connector}
}


func (c *Client) NewWrapBETHService() *WrapBETHService {
	return &WrapBETHService{C: c.Connector}
}

