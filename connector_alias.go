
package binance_connector

import (
	"github.com/niklak/binance_connector/api/staking"
	"github.com/niklak/binance_connector/api/subaccount"
	"github.com/niklak/binance_connector/api/margin"
	"github.com/niklak/binance_connector/api/ethstaking"
	"github.com/niklak/binance_connector/api/account"
	"github.com/niklak/binance_connector/api/market"
	"github.com/niklak/binance_connector/api/spot"
	"github.com/niklak/binance_connector/api/wallet"
)

//file: api/account/account.go
type AccountService = account.AccountService
type AccountResponse = account.AccountResponse
type Balance = account.Balance

//file: api/account/comission.go
type AccountCommissionService = account.AccountCommissionService
type Commission = account.Commission
type CommissionDiscount = account.CommissionDiscount
type AccountCommissionResponse = account.AccountCommissionResponse

//file: api/account/currentordercount.go
type QueryCurrentOrderCountUsageService = account.QueryCurrentOrderCountUsageService
type QueryCurrentOrderCountUsageResponse = account.QueryCurrentOrderCountUsageResponse

//file: api/account/myallocations.go
type AccountAllocationsService = account.AccountAllocationsService
type AccountAllocationsResponse = account.AccountAllocationsResponse

//file: api/account/preventedmatches.go
type QueryPreventedMatchesService = account.QueryPreventedMatchesService
type QueryPreventedMatchesResponse = account.QueryPreventedMatchesResponse

//file: api/account/tradelist.go
type AccountTradeListService = account.AccountTradeListService
type AccountTradeListResponse = account.AccountTradeListResponse

//file: api/market/aggtrades.go
type AggTradesList = market.AggTradesList
type AggTradesListResponse = market.AggTradesListResponse

//file: api/market/avgprice.go
type AvgPrice = market.AvgPrice
type AvgPriceResponse = market.AvgPriceResponse

//file: api/market/bookticker.go
type TickerBookTicker = market.TickerBookTicker
type TickerBookTickerResponse = market.TickerBookTickerResponse

//file: api/market/exchangeinfo.go
type ExchangeInfo = market.ExchangeInfo
type ExchangeInfoResponse = market.ExchangeInfoResponse
type RateLimit = market.RateLimit
type ExchangeFilter = market.ExchangeFilter
type SymbolInfo = market.SymbolInfo
type SymbolFilter = market.SymbolFilter

//file: api/market/historicaltrades.go
type HistoricalTradeLookup = market.HistoricalTradeLookup

//file: api/market/klines.go
type Klines = market.Klines
type KlinesResponse = market.KlinesResponse

//file: api/market/orderbook.go
type OrderBook = market.OrderBook
type OrderBookResponse = market.OrderBookResponse

//file: api/market/ping.go
type Ping = market.Ping

//file: api/market/recenttrades.go
type RecentTradesList = market.RecentTradesList
type RecentTradesListResponse = market.RecentTradesListResponse

//file: api/market/servertime.go
type ServerTime = market.ServerTime
type ServerTimeResponse = market.ServerTimeResponse

//file: api/market/ticker.go
type Ticker = market.Ticker
type TickerResponse = market.TickerResponse

//file: api/market/ticker24.go
type Ticker24hr = market.Ticker24hr
type Ticker24hrResponse = market.Ticker24hrResponse

//file: api/market/tickerprice.go
type TickerPrice = market.TickerPrice
type TickerPriceResponse = market.TickerPriceResponse

//file: api/market/uiklines.go
type UiKlines = market.UiKlines

//file: api/spot/canceloco.go
type CancelOCOService = spot.CancelOCOService

//file: api/spot/cancelopenorders.go
type CancelOpenOrdersService = spot.CancelOpenOrdersService

//file: api/spot/cancelorder.go
type CancelOrderService = spot.CancelOrderService

//file: api/spot/cancelreplaceservice.go
type CancelReplaceService = spot.CancelReplaceService
type OrderData = spot.OrderData
type CancelReplaceResponse = spot.CancelReplaceResponse

//file: api/spot/createorder.go
type CreateOrderService = spot.CreateOrderService
type CreateOrderResponseACK = spot.CreateOrderResponseACK
type CreateOrderResponseRESULT = spot.CreateOrderResponseRESULT
type CreateOrderResponseFULL = spot.CreateOrderResponseFULL

//file: api/spot/getallorders.go
type GetAllOrdersService = spot.GetAllOrdersService

//file: api/spot/getopenorders.go
type GetOpenOrdersService = spot.GetOpenOrdersService

//file: api/spot/getorder.go
type GetOrderService = spot.GetOrderService

//file: api/spot/newoco.go
type NewOCOService = spot.NewOCOService
type OrderOCOResponse = spot.OrderOCOResponse

//file: api/spot/neworderlistoco.go
type CreateOrderListOCOService = spot.CreateOrderListOCOService

//file: api/spot/orderresponse.go
type OrderResponse = spot.OrderResponse

//file: api/spot/queryalloco.go
type QueryAllOCOService = spot.QueryAllOCOService

//file: api/spot/queryoco.go
type QueryOCOService = spot.QueryOCOService
type OCOResponse = spot.OCOResponse

//file: api/spot/queryopenoco.go
type QueryOpenOCOService = spot.QueryOpenOCOService

//file: api/spot/testneworder.go
type TestNewOrder = spot.TestNewOrder
type AccountOrderBookResponse = spot.AccountOrderBookResponse

//file: api/wallet/accountstatus.go
type AccountStatusService = wallet.AccountStatusService
type AccountStatusResponse = wallet.AccountStatusResponse

//file: api/wallet/accounttradingstatus.go
type AccountApiTradingStatusService = wallet.AccountApiTradingStatusService
type AccountApiTradingStatusResponse = wallet.AccountApiTradingStatusResponse

//file: api/wallet/allcoins.go
type AllCoinsInfoService = wallet.AllCoinsInfoService
type CoinInfo = wallet.CoinInfo

//file: api/wallet/apikeypermission.go
type APIKeyPermissionService = wallet.APIKeyPermissionService
type APIKeyPermissionResponse = wallet.APIKeyPermissionResponse

//file: api/wallet/assetdetail.go
type AssetDetailV2Service = wallet.AssetDetailV2Service
type AssetDetailV2Response = wallet.AssetDetailV2Response

//file: api/wallet/assetdivident.go
type AssetDividendRecordService = wallet.AssetDividendRecordService
type AssetDividendRecordResponse = wallet.AssetDividendRecordResponse

//file: api/wallet/autoconverting.go
type AutoConvertStableCoinService = wallet.AutoConvertStableCoinService
type AutoConvertStableCoinResponse = wallet.AutoConvertStableCoinResponse

//file: api/wallet/cloudmininghistory.go
type CloudMiningPaymentHistoryService = wallet.CloudMiningPaymentHistoryService
type CloudMiningPaymentHistoryResponse = wallet.CloudMiningPaymentHistoryResponse

//file: api/wallet/dailiyaccountsnapshot.go
type AccountSnapshotService = wallet.AccountSnapshotService
type AccountSnapshotResponse = wallet.AccountSnapshotResponse

//file: api/wallet/depositaddress.go
type DepositAddressService = wallet.DepositAddressService
type DepositAddressResponse = wallet.DepositAddressResponse

//file: api/wallet/deposithistory.go
type DepositHistoryService = wallet.DepositHistoryService
type DepositHistoryResponse = wallet.DepositHistoryResponse

//file: api/wallet/disablefastwithdraw.go
type DisableFastWithdrawSwitchService = wallet.DisableFastWithdrawSwitchService
type DisableFastWithdrawSwitchResponse = wallet.DisableFastWithdrawSwitchResponse

//file: api/wallet/dustbtc.go
type AssetDetailService = wallet.AssetDetailService
type AssetDetailResponse = wallet.AssetDetailResponse

//file: api/wallet/dustlog.go
type DustLogService = wallet.DustLogService
type DustLogResponse = wallet.DustLogResponse

//file: api/wallet/dusttransfer.go
type DustTransferService = wallet.DustTransferService
type DustTransferResponse = wallet.DustTransferResponse
type DustTransferResult = wallet.DustTransferResult

//file: api/wallet/enablefastwithdraw.go
type EnableFastWithdrawSwitchService = wallet.EnableFastWithdrawSwitchService
type EnableFastWithdrawSwitchResponse = wallet.EnableFastWithdrawSwitchResponse

//file: api/wallet/fundingwallet.go
type FundingWalletService = wallet.FundingWalletService
type FundingWalletResponse = wallet.FundingWalletResponse

//file: api/wallet/systemstatus.go
type SystemStatusService = wallet.SystemStatusService
type SystemStatusResponse = wallet.SystemStatusResponse

//file: api/wallet/tradefee.go
type TradeFeeService = wallet.TradeFeeService
type TradeFeeResponse = wallet.TradeFeeResponse

//file: api/wallet/universaltransfer.go
type UserUniversalTransferService = wallet.UserUniversalTransferService
type UserUniversalTransferResponse = wallet.UserUniversalTransferResponse

//file: api/wallet/universaltransferhistory.go
type UserUniversalTransferHistoryService = wallet.UserUniversalTransferHistoryService
type UserUniversalTransferHistoryResponse = wallet.UserUniversalTransferHistoryResponse

//file: api/wallet/userasset.go
type UserAssetService = wallet.UserAssetService
type UserAssetResponse = wallet.UserAssetResponse

//file: api/wallet/withdraw.go
type WithdrawService = wallet.WithdrawService
type WithdrawResponse = wallet.WithdrawResponse

//file: api/wallet/withdrawhistory.go
type WithdrawHistoryService = wallet.WithdrawHistoryService
type WithdrawHistoryResponse = wallet.WithdrawHistoryResponse

//file: api/staking/personalleftquota.go
type PersonalLeftQuotaService = staking.PersonalLeftQuotaService
type PersonalLeftQuotaResponse = staking.PersonalLeftQuotaResponse

//file: api/staking/purchasestakingproduct.go
type PurchaseStakingProductService = staking.PurchaseStakingProductService
type PurchaseStakingProductResponse = staking.PurchaseStakingProductResponse

//file: api/staking/redeemstakingproduct.go
type RedeemStakingProductService = staking.RedeemStakingProductService
type RedeemStakingProductResponse = staking.RedeemStakingProductResponse

//file: api/staking/setautostaking.go
type SetAutoStakingService = staking.SetAutoStakingService
type SetAutoStakingResponse = staking.SetAutoStakingResponse

//file: api/staking/stakinghistory.go
type StakingHistoryService = staking.StakingHistoryService
type StakingHistoryResponse = staking.StakingHistoryResponse

//file: api/staking/stakingproductlist.go
type StakingProductListService = staking.StakingProductListService
type StakingProductListResponse = staking.StakingProductListResponse

//file: api/staking/stakingproductposition.go
type StakingProductPositionService = staking.StakingProductPositionService
type StakingProductPositionResponse = staking.StakingProductPositionResponse

//file: api/subaccount/subaccount.go
type CreateSubAccountService = subaccount.CreateSubAccountService
type CreateSubAccountResp = subaccount.CreateSubAccountResp
type QuerySubAccountListService = subaccount.QuerySubAccountListService
type SubAccountListResp = subaccount.SubAccountListResp
type SubAccount = subaccount.SubAccount
type QuerySubAccountSpotAssetTransferHistoryService = subaccount.QuerySubAccountSpotAssetTransferHistoryService
type SubAccountTransferHistoryResponse = subaccount.SubAccountTransferHistoryResponse
type QuerySubAccountFuturesAssetTransferHistoryService = subaccount.QuerySubAccountFuturesAssetTransferHistoryService
type QuerySubAccountFuturesAssetTransferHistoryResp = subaccount.QuerySubAccountFuturesAssetTransferHistoryResp
type SubAccountFuturesAssetTransferService = subaccount.SubAccountFuturesAssetTransferService
type SubAccountFuturesAssetTransferResp = subaccount.SubAccountFuturesAssetTransferResp
type QuerySubAccountAssetsService = subaccount.QuerySubAccountAssetsService
type QuerySubAccountAssetsResp = subaccount.QuerySubAccountAssetsResp
type QuerySubAccountSpotAssetsSummaryService = subaccount.QuerySubAccountSpotAssetsSummaryService
type QuerySubAccountSpotAssetsSummaryResp = subaccount.QuerySubAccountSpotAssetsSummaryResp
type SubAccountDepositAddressService = subaccount.SubAccountDepositAddressService
type SubAccountDepositAddressResp = subaccount.SubAccountDepositAddressResp
type SubAccountDepositHistoryService = subaccount.SubAccountDepositHistoryService
type SubAccountDepositHistoryResponse = subaccount.SubAccountDepositHistoryResponse
type SubAccountStatusService = subaccount.SubAccountStatusService
type SubAccountStatusResp = subaccount.SubAccountStatusResp
type EnableMarginForSubAccountService = subaccount.EnableMarginForSubAccountService
type EnableMarginForSubAccountResp = subaccount.EnableMarginForSubAccountResp
type DetailOnSubAccountMarginAccountService = subaccount.DetailOnSubAccountMarginAccountService
type DetailOnSubAccountMarginAccountResp = subaccount.DetailOnSubAccountMarginAccountResp
type SummaryOfSubAccountMarginAccountService = subaccount.SummaryOfSubAccountMarginAccountService
type SummaryOfSubAccountMarginAccountResp = subaccount.SummaryOfSubAccountMarginAccountResp
type EnableFuturesForSubAccountService = subaccount.EnableFuturesForSubAccountService
type EnableFuturesForSubAccountResp = subaccount.EnableFuturesForSubAccountResp
type DetailOnSubAccountFuturesAccountService = subaccount.DetailOnSubAccountFuturesAccountService
type DetailOnSubAccountFuturesAccountResp = subaccount.DetailOnSubAccountFuturesAccountResp
type SummaryOfSubAccountFuturesAccountService = subaccount.SummaryOfSubAccountFuturesAccountService
type SummaryOfSubAccountFuturesAccountResp = subaccount.SummaryOfSubAccountFuturesAccountResp
type FuturesPositionRiskOfSubAccountService = subaccount.FuturesPositionRiskOfSubAccountService
type FuturesPositionRiskOfSubAccountResp = subaccount.FuturesPositionRiskOfSubAccountResp
type FuturesTransferForSubAccountService = subaccount.FuturesTransferForSubAccountService
type FuturesTransferForSubAccountResp = subaccount.FuturesTransferForSubAccountResp
type MarginTransferForSubAccountService = subaccount.MarginTransferForSubAccountService
type MarginTransferForSubAccountResp = subaccount.MarginTransferForSubAccountResp
type TransferToSubAccountOfSameMasterService = subaccount.TransferToSubAccountOfSameMasterService
type TransferToSubAccountOfSameMasterResp = subaccount.TransferToSubAccountOfSameMasterResp
type TransferToMasterService = subaccount.TransferToMasterService
type TransferToMasterResp = subaccount.TransferToMasterResp
type SubAccountTransferHistoryService = subaccount.SubAccountTransferHistoryService
type SubAccountTransferHistoryResp = subaccount.SubAccountTransferHistoryResp
type UniversalTransferService = subaccount.UniversalTransferService
type UniversalTransferResp = subaccount.UniversalTransferResp
type QueryUniversalTransferHistoryService = subaccount.QueryUniversalTransferHistoryService
type QueryUniversalTransferHistoryResp = subaccount.QueryUniversalTransferHistoryResp
type InternalUniversalTransfer = subaccount.InternalUniversalTransfer
type DetailOnSubAccountFuturesAccountV2Service = subaccount.DetailOnSubAccountFuturesAccountV2Service
type FuturesAccountAsset = subaccount.FuturesAccountAsset
type DetailOnSubAccountFuturesAccountV2USDTResp = subaccount.DetailOnSubAccountFuturesAccountV2USDTResp
type DetailOnSubAccountFuturesAccountV2COINResp = subaccount.DetailOnSubAccountFuturesAccountV2COINResp
type SummaryOfSubAccountFuturesAccountV2Service = subaccount.SummaryOfSubAccountFuturesAccountV2Service
type SummaryOfSubAccountFuturesAccountV2USDTResp = subaccount.SummaryOfSubAccountFuturesAccountV2USDTResp
type SummaryOfSubAccountFuturesAccountV2COINResp = subaccount.SummaryOfSubAccountFuturesAccountV2COINResp
type FuturesPositionRiskOfSubAccountV2Service = subaccount.FuturesPositionRiskOfSubAccountV2Service
type FuturesPositionRiskOfSubAccountV2USDTResp = subaccount.FuturesPositionRiskOfSubAccountV2USDTResp
type FuturesPositionRiskOfSubAccountV2COINResp = subaccount.FuturesPositionRiskOfSubAccountV2COINResp
type EnableLeverageTokenForSubAccountService = subaccount.EnableLeverageTokenForSubAccountService
type EnableLeverageTokenForSubAccountResp = subaccount.EnableLeverageTokenForSubAccountResp
type IPRestrictionForSubAccountAPIKeyService = subaccount.IPRestrictionForSubAccountAPIKeyService
type IPRestrictionForSubAccountAPIKeyResp = subaccount.IPRestrictionForSubAccountAPIKeyResp
type DeleteIPListForSubAccountAPIKeyService = subaccount.DeleteIPListForSubAccountAPIKeyService
type DeleteIPListForSubAccountAPIKeyResp = subaccount.DeleteIPListForSubAccountAPIKeyResp
type UpdateIPRestrictionForSubAccountAPIKeyService = subaccount.UpdateIPRestrictionForSubAccountAPIKeyService
type UpdateIPRestrictionForSubAccountAPIKeyResp = subaccount.UpdateIPRestrictionForSubAccountAPIKeyResp
type DepositAssetsIntoTheManagedSubAccountService = subaccount.DepositAssetsIntoTheManagedSubAccountService
type DepositAssetsIntoTheManagedSubAccountResp = subaccount.DepositAssetsIntoTheManagedSubAccountResp
type QueryManagedSubAccountAssetDetailsService = subaccount.QueryManagedSubAccountAssetDetailsService
type QueryManagedSubAccountAssetDetailsResp = subaccount.QueryManagedSubAccountAssetDetailsResp
type WithdrawAssetsFromTheManagedSubAccountService = subaccount.WithdrawAssetsFromTheManagedSubAccountService
type WithdrawAssetsFromTheManagedSubAccountResp = subaccount.WithdrawAssetsFromTheManagedSubAccountResp
type QueryManagedSubAccountSnapshotService = subaccount.QueryManagedSubAccountSnapshotService
type QueryManagedSubAccountSnapshotResp = subaccount.QueryManagedSubAccountSnapshotResp
type QueryManagedSubAccountTransferLogService = subaccount.QueryManagedSubAccountTransferLogService
type QueryManagedSubAccountTransferLogResp = subaccount.QueryManagedSubAccountTransferLogResp
type QueryManagedSubAccountFuturesAssetDetailsService = subaccount.QueryManagedSubAccountFuturesAssetDetailsService
type QueryManagedSubAccountFuturesAssetDetailsResp = subaccount.QueryManagedSubAccountFuturesAssetDetailsResp
type QueryManagedSubAccountMarginAssetDetailsService = subaccount.QueryManagedSubAccountMarginAssetDetailsService
type QueryManagedSubAccountMarginAssetDetailsResp = subaccount.QueryManagedSubAccountMarginAssetDetailsResp
type QueryManagedSubAccountTransferLogForTradingTeamService = subaccount.QueryManagedSubAccountTransferLogForTradingTeamService
type QueryManagedSubAccountTransferLogForTradingTeamResp = subaccount.QueryManagedSubAccountTransferLogForTradingTeamResp
type QuerySubAccountAssetsForMasterAccountService = subaccount.QuerySubAccountAssetsForMasterAccountService
type QuerySubAccountAssetsForMasterAccountResp = subaccount.QuerySubAccountAssetsForMasterAccountResp
type QueryManagedSubAccountList = subaccount.QueryManagedSubAccountList
type QueryManagedSubAccountListResp = subaccount.QueryManagedSubAccountListResp
type QuerySubAccountTransactionStatistics = subaccount.QuerySubAccountTransactionStatistics
type QuerySubAccountTransactionStatisticsResp = subaccount.QuerySubAccountTransactionStatisticsResp
type ManagedSubAccountDepositAddressService = subaccount.ManagedSubAccountDepositAddressService
type ManagedSubAccountDepositAddressResp = subaccount.ManagedSubAccountDepositAddressResp

//file: api/margin/margin.go
type TransferService = margin.TransferService
type TransferResponse = margin.TransferResponse
type BorrowService = margin.BorrowService
type BorrowResponse = margin.BorrowResponse
type RepayService = margin.RepayService
type RepayResponse = margin.RepayResponse
type QueryMarginAssetService = margin.QueryMarginAssetService
type QueryMarginAssetResponse = margin.QueryMarginAssetResponse
type QueryCrossMarginPairService = margin.QueryCrossMarginPairService
type QueryCrossMarginPairResponse = margin.QueryCrossMarginPairResponse
type GetAllMarginAssetsService = margin.GetAllMarginAssetsService
type GetAllMarginAssetsResponse = margin.GetAllMarginAssetsResponse
type GetAllMarginPairsService = margin.GetAllMarginPairsService
type GetAllMarginPairsResponse = margin.GetAllMarginPairsResponse
type QueryMarginPriceIndexService = margin.QueryMarginPriceIndexService
type QueryMarginPriceIndexResponse = margin.QueryMarginPriceIndexResponse
type MarginAccountNewOrderService = margin.MarginAccountNewOrderService
type MarginAccountNewOrderResponseACK = margin.MarginAccountNewOrderResponseACK
type MarginAccountNewOrderResponseRESULT = margin.MarginAccountNewOrderResponseRESULT
type MarginAccountNewOrderResponseFULL = margin.MarginAccountNewOrderResponseFULL
type MarginAccountCancelOrderService = margin.MarginAccountCancelOrderService
type MarginAccountCancelOrderResponse = margin.MarginAccountCancelOrderResponse
type MarginAccountCancelAllOrdersService = margin.MarginAccountCancelAllOrdersService
type MarginAccountCancelAllOrdersResponse = margin.MarginAccountCancelAllOrdersResponse
type CrossMarginTransferHistoryResponse = margin.CrossMarginTransferHistoryResponse
type QueryMarginBorrowRepayService = margin.QueryMarginBorrowRepayService
type QueryMarginBorrowRepay = margin.QueryMarginBorrowRepay
type InterestHistoryService = margin.InterestHistoryService
type InterestHistoryResponse = margin.InterestHistoryResponse
type ForceLiquidationRecordService = margin.ForceLiquidationRecordService
type ForceLiquidationRecordResponse = margin.ForceLiquidationRecordResponse
type CrossMarginAccountDetailService = margin.CrossMarginAccountDetailService
type CrossMarginAccountDetailResponse = margin.CrossMarginAccountDetailResponse
type MarginAccountOrderService = margin.MarginAccountOrderService
type MarginAccountOrderResponse = margin.MarginAccountOrderResponse
type MarginAccountOpenOrderService = margin.MarginAccountOpenOrderService
type MarginAccountOpenOrderResponse = margin.MarginAccountOpenOrderResponse
type MarginAccountAllOrderService = margin.MarginAccountAllOrderService
type MarginAccountAllOrderResponse = margin.MarginAccountAllOrderResponse
type MarginAccountNewOCOService = margin.MarginAccountNewOCOService
type MarginAccountNewOCOResponse = margin.MarginAccountNewOCOResponse
type MarginAccountCancelOCOService = margin.MarginAccountCancelOCOService
type MarginAccountCancelOCOResponse = margin.MarginAccountCancelOCOResponse
type MarginAccountQueryOCOService = margin.MarginAccountQueryOCOService
type MarginAccountQueryOCOResponse = margin.MarginAccountQueryOCOResponse
type MarginAccountQueryAllOCOService = margin.MarginAccountQueryAllOCOService
type MarginAccountQueryAllOCOResponse = margin.MarginAccountQueryAllOCOResponse
type MarginAccountQueryOpenOCOService = margin.MarginAccountQueryOpenOCOService
type MarginAccountQueryOpenOCOResponse = margin.MarginAccountQueryOpenOCOResponse
type MarginAccountQueryTradeListService = margin.MarginAccountQueryTradeListService
type MarginAccountQueryTradeListResponse = margin.MarginAccountQueryTradeListResponse
type MarginAccountQueryMaxBorrowService = margin.MarginAccountQueryMaxBorrowService
type MarginAccountQueryMaxBorrowResponse = margin.MarginAccountQueryMaxBorrowResponse
type MarginAccountQueryMaxTransferOutAmountService = margin.MarginAccountQueryMaxTransferOutAmountService
type MarginAccountQueryMaxTransferOutAmountResponse = margin.MarginAccountQueryMaxTransferOutAmountResponse
type MarginAccountSummaryService = margin.MarginAccountSummaryService
type MarginAccountSummaryResponse = margin.MarginAccountSummaryResponse
type MarginIsolatedAccountTransferService = margin.MarginIsolatedAccountTransferService
type MarginIsolatedAccountTransferResponse = margin.MarginIsolatedAccountTransferResponse
type MarginIsolatedAccountTransferHistoryService = margin.MarginIsolatedAccountTransferHistoryService
type MarginIsolatedAccountTransferHistoryResponse = margin.MarginIsolatedAccountTransferHistoryResponse
type MarginIsolatedAccountInfoService = margin.MarginIsolatedAccountInfoService
type MarginIsolatedAccountInfoAssets = margin.MarginIsolatedAccountInfoAssets
type MarginIsolatedAccountInfoResponseSymbols = margin.MarginIsolatedAccountInfoResponseSymbols
type MarginIsolatedAccountInfoResponse = margin.MarginIsolatedAccountInfoResponse
type MarginIsolatedAccountDisableService = margin.MarginIsolatedAccountDisableService
type MarginIsolatedAccountDisableResponse = margin.MarginIsolatedAccountDisableResponse
type MarginIsolatedAccountEnableService = margin.MarginIsolatedAccountEnableService
type MarginIsolatedAccountEnableResponse = margin.MarginIsolatedAccountEnableResponse
type MarginIsolatedAccountLimitService = margin.MarginIsolatedAccountLimitService
type MarginIsolatedAccountLimitResponse = margin.MarginIsolatedAccountLimitResponse
type MarginIsolatedSymbolService = margin.MarginIsolatedSymbolService
type MarginIsolatedSymbolResponse = margin.MarginIsolatedSymbolResponse
type AllIsolatedMarginSymbolService = margin.AllIsolatedMarginSymbolService
type MarginToggleBnbBurnService = margin.MarginToggleBnbBurnService
type MarginToggleBnbBurnResponse = margin.MarginToggleBnbBurnResponse
type MarginBnbBurnStatusService = margin.MarginBnbBurnStatusService
type MarginBnbBurnStatusResponse = margin.MarginBnbBurnStatusResponse
type MarginInterestRateHistoryService = margin.MarginInterestRateHistoryService
type MarginInterestRateHistoryResponse = margin.MarginInterestRateHistoryResponse
type MarginCrossMarginFeeService = margin.MarginCrossMarginFeeService
type MarginCrossMarginFeeResponse = margin.MarginCrossMarginFeeResponse
type MarginIsolatedMarginFeeService = margin.MarginIsolatedMarginFeeService
type MarginIsolatedMarginFeeResponse = margin.MarginIsolatedMarginFeeResponse
type MarginIsolatedMarginTierService = margin.MarginIsolatedMarginTierService
type MarginIsolatedMarginTierResponse = margin.MarginIsolatedMarginTierResponse
type MarginCurrentOrderCountService = margin.MarginCurrentOrderCountService
type MarginCurrentOrderCountResponse = margin.MarginCurrentOrderCountResponse
type MarginDustlogService = margin.MarginDustlogService
type MarginDustlogResponse = margin.MarginDustlogResponse
type UserAssetDribblet = margin.UserAssetDribblet
type UserAssetDribbletDetail = margin.UserAssetDribbletDetail
type MarginCrossCollateralRatioService = margin.MarginCrossCollateralRatioService
type MarginCrossCollateralRatioResponse = margin.MarginCrossCollateralRatioResponse
type MarginSmallLiabilityExchangeCoinListService = margin.MarginSmallLiabilityExchangeCoinListService
type MarginSmallLiabilityExchangeCoinListResponse = margin.MarginSmallLiabilityExchangeCoinListResponse
type MarginSmallLiabilityExchangeService = margin.MarginSmallLiabilityExchangeService
type MarginSmallLiabilityExchangeResponse = margin.MarginSmallLiabilityExchangeResponse
type MarginSmallLiabilityExchangeHistoryService = margin.MarginSmallLiabilityExchangeHistoryService
type MarginSmallLiabilityExchangeHistoryResponse = margin.MarginSmallLiabilityExchangeHistoryResponse

//file: api/ethstaking/bethrewardshistory.go
type BETHRewardsDistributionHistoryService = ethstaking.BETHRewardsDistributionHistoryService
type BETHRewardHistoryRow = ethstaking.BETHRewardHistoryRow
type BETHRewardsDistributionHistoryResponse = ethstaking.BETHRewardsDistributionHistoryResponse

//file: api/ethstaking/currentstakingquota.go
type CurrentETHStakingQuotaService = ethstaking.CurrentETHStakingQuotaService
type CurrentETHStakingQuotaResponse = ethstaking.CurrentETHStakingQuotaResponse

//file: api/ethstaking/redeemeth.go
type RedeemETHService = ethstaking.RedeemETHService
type RedeemETHResponse = ethstaking.RedeemETHResponse

//file: api/ethstaking/redemtionhistory.go
type ETHRedemptionHistoryService = ethstaking.ETHRedemptionHistoryService

//file: api/ethstaking/stakingaccount.go
type ETHStakingAccountService = ethstaking.ETHStakingAccountService
type ETHStakingAccountResponse = ethstaking.ETHStakingAccountResponse

//file: api/ethstaking/stakingaccountv2.go
type ETHStakingAccountV2Service = ethstaking.ETHStakingAccountV2Service
type ETHStakingAccountV2Response = ethstaking.ETHStakingAccountV2Response

//file: api/ethstaking/stakinghistory.go
type ETHStakingHistoryService = ethstaking.ETHStakingHistoryService
type ETHHistoryRow = ethstaking.ETHHistoryRow
type ETHHistoryResponse = ethstaking.ETHHistoryResponse

//file: api/ethstaking/subscribestaking.go
type SubscribeETHStakingService = ethstaking.SubscribeETHStakingService
type SubscribeETHStakingResponse = ethstaking.SubscribeETHStakingResponse

//file: api/ethstaking/subscribestakingv2.go
type SubscribeETHStakingV2Service = ethstaking.SubscribeETHStakingV2Service
type SubscribeETHStakingV2Response = ethstaking.SubscribeETHStakingV2Response

//file: api/ethstaking/wbethratehistory.go
type WBETHRateHistoryService = ethstaking.WBETHRateHistoryService
type WBETHRateHistoryRow = ethstaking.WBETHRateHistoryRow
type WBETHRateHistoryResponse = ethstaking.WBETHRateHistoryResponse

//file: api/ethstaking/wbethrewardshistory.go
type WBETHRewardsHistoryService = ethstaking.WBETHRewardsHistoryService
type WBETHRewardsHistoryRow = ethstaking.WBETHRewardsHistoryRow
type WBETHRewardsHistoryResponse = ethstaking.WBETHRewardsHistoryResponse

//file: api/ethstaking/wbethunwraphistory.go
type WBETHUnwrapHistoryService = ethstaking.WBETHUnwrapHistoryService

//file: api/ethstaking/wbethwraphistory.go
type WBETHWrapHistoryService = ethstaking.WBETHWrapHistoryService
type WBETHHistoryRow = ethstaking.WBETHHistoryRow
type WBETHHistoryResponse = ethstaking.WBETHHistoryResponse

//file: api/ethstaking/wrapbeth.go
type WrapBETHService = ethstaking.WrapBETHService
type WrapBETHResponse = ethstaking.WrapBETHResponse

