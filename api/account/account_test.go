package account

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/niklak/binance_connector/internal/connector"
)

type AccountTestSuite struct {
	suite.Suite
	testServer       *httptest.Server
	binanceConnector *connector.Connector
}

func (s *AccountTestSuite) SetupSuite() {

	//zerolog.SetGlobalLevel(zerolog.Level(1))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/api/v3/account":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
			{
				"makerCommission": 15,
				"takerCommission": 15,
				"buyerCommission": 0,
				"sellerCommission": 0,
				"commissionRates": {
				  "maker": "0.00150000",
				  "taker": "0.00150000",
				  "buyer": "0.00000000",
				  "seller": "0.00000000"
				},
				"canTrade": true,
				"canWithdraw": true,
				"canDeposit": true,
				"brokered": false,
				"requireSelfTradePrevention": false,
				"preventSor": false,
				"updateTime": 123456789,
				"accountType": "SPOT",
				"balances": [
				  {
					"asset": "BTC",
					"free": "4723846.89208129",
					"locked": "0.00000000"
				  },
				  {
					"asset": "LTC",
					"free": "4763368.68006011",
					"locked": "0.00000000"
				  }
				],
				"permissions": [
				  "SPOT"
				],
				"uid": 354937868
			  }
			`))

		case "/api/v3/myTrades":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
			[
				{
					"symbol": "BNBBTC",
					"id": 28457,
					"orderId": 100234,
					"orderListId": -1,
					"price": "4.00000100",
					"qty": "12.00000000",
					"quoteQty": "48.000012",
					"commission": "10.10000000",
					"commissionAsset": "BNB",
					"time": 1499865549590,
					"isBuyer": true,
					"isMaker": false,
					"isBestMatch": true
				}
			]			
			`))

		case "/api/v3/rateLimit/order":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
			[
				{
					"rateLimitType": "ORDERS",
					"interval": "SECOND",
					"intervalNum": 10,
					"limit": 10000,
					"count": 0
				},
				{
					"rateLimitType": "ORDERS",
					"interval": "DAY",
					"intervalNum": 1,
					"limit": 20000,
					"count": 0
				}
			]
			`))

		case "/api/v3/myPreventedMatches":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
			[
				{
					"symbol": "BTCUSDT",
					"preventedMatchId": 1,
					"takerOrderId": 5,
					"makerOrderId": 3,
					"tradeGroupId": 1,
					"selfTradePreventionMode": "EXPIRE_MAKER",
					"price": "1.100000",
					"makerPreventedQuantity": "1.300000",
					"transactTime": 1669101687094
				}
			]
			`))

		default:
			http.NotFound(w, r)
		}
		// Serve the request
	})
	s.testServer = httptest.NewServer(handler)
	s.binanceConnector = connector.NewConnector("test-key", "test-secret", connector.BaseURL(s.testServer.URL))
}

func (s *AccountTestSuite) TearDownSuite() {
	// Setup the test
	s.testServer.Close()
}

func (s *AccountTestSuite) TestAccountService() {
	accountService := AccountService{C: s.binanceConnector}
	accountResponse, err := accountService.Do(context.Background())
	s.NoError(err)

	expected := &AccountResponse{
		MakerCommission:  15,
		TakerCommission:  15,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      true,
		CanDeposit:       true,
		UpdateTime:       123456789,
		AccountType:      "SPOT",
		Balances: []Balance{
			{Asset: "BTC", Free: "4723846.89208129", Locked: "0.00000000"},
			{Asset: "LTC", Free: "4763368.68006011", Locked: "0.00000000"},
		},
		Permissions: []string{"SPOT"}}
	assert.Equal(s.T(), accountResponse, expected)
}

func (s *AccountTestSuite) TestTradesList() {
	accountTradeListService := AccountTradeListService{C: s.binanceConnector}
	accountTradeListResponse, err := accountTradeListService.Symbol("BNBBTC").Do(context.Background())
	s.NoError(err)

	expected := []*AccountTradeListResponse{
		{
			Symbol:          "BNBBTC",
			Id:              28457,
			OrderId:         100234,
			OrderListId:     -1,
			Price:           "4.00000100",
			Quantity:        "12.00000000",
			QuoteQuantity:   "48.000012",
			Commission:      "10.10000000",
			CommissionAsset: "BNB",
			Time:            1499865549590,
			IsBuyer:         true,
			IsMaker:         false,
			IsBestMatch:     true,
		},
	}
	assert.Equal(s.T(), accountTradeListResponse, expected)
}

func (s *AccountTestSuite) TestRateLimit() {
	rateLimitService := QueryCurrentOrderCountUsageService{C: s.binanceConnector}
	rateLimitResponse, err := rateLimitService.Do(context.Background())
	s.NoError(err)

	expected := []*QueryCurrentOrderCountUsageResponse{
		{
			RateLimitType: "ORDERS",
			Interval:      "SECOND",
			IntervalNum:   10,
			Limit:         10000,
			Count:         0,
		},
		{
			RateLimitType: "ORDERS",
			Interval:      "DAY",
			IntervalNum:   1,
			Limit:         20000,
			Count:         0,
		},
	}
	assert.Equal(s.T(), rateLimitResponse, expected)
}

func (s *AccountTestSuite) TestQueryPreventedMatches() {
	queryPreventedMatchesService := QueryPreventedMatchesService{C: s.binanceConnector}
	queryPreventedMatchesResponse, err := queryPreventedMatchesService.Symbol("BTCUSDT").Do(context.Background())
	s.NoError(err)

	expected := []*QueryPreventedMatchesResponse{
		{
			Symbol:                  "BTCUSDT",
			PreventedMatchId:        1,
			TakerOrderId:            5,
			MakerOrderId:            3,
			TradeGroupId:            1,
			SelfTradePreventionMode: "EXPIRE_MAKER",
			Price:                   "1.100000",
			MakerPreventedQuantity:  "1.300000",
			TransactTime:            1669101687094,
		},
	}
	assert.Equal(s.T(), queryPreventedMatchesResponse, expected)

}
func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
