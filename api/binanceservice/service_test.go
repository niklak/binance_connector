package binanceservice

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Level(1))
}

func TestSimpleService_Do(t *testing.T) {
	// Checking server time
	assert := assert.New(t)

	apiKey := os.Getenv("BINANCE_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET")

	assert.NotZero(apiKey)
	assert.NotZero(apiSecret)

	c := connector.NewConnector(apiKey, apiSecret, connector.BaseURL("https://testnet.binance.vision"))

	s := &Service{C: c}

	type response struct {
		ServerTime int64 `json:"serverTime"`
	}

	var res response

	err := s.Do(context.Background(), "/api/v3/time", &res)

	assert.NoError(err)

	assert.NotZero(res.ServerTime)
}

func TestSignedService_Do(t *testing.T) {
	// Checking account information
	assert := assert.New(t)

	apiKey := os.Getenv("BINANCE_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET")

	assert.NotZero(apiKey)
	assert.NotZero(apiSecret)

	c := connector.NewConnector(apiKey, apiSecret, connector.BaseURL("https://testnet.binance.vision"))

	s := &Service{C: c}

	type response struct {
		CanTrade bool `json:"canTrade"`
	}

	var res response

	err := s.Do(context.Background(), "/api/v3/account", &res, request.Signed())

	assert.NoError(err)
	assert.True(res.CanTrade)
}

func TestNewOrderService_Do(t *testing.T) {
	// Creating a new order
	assert := assert.New(t)

	apiKey := os.Getenv("BINANCE_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET")

	assert.NotZero(apiKey)
	assert.NotZero(apiSecret)

	c := connector.NewConnector(apiKey, apiSecret, connector.BaseURL("https://testnet.binance.vision"))

	s := &Service{C: c}

	type response struct{}

	var res response

	err := s.Do(
		context.Background(),
		"/api/v3/order/test",
		&res,
		request.Method(http.MethodPost),
		request.Signed(),
		request.RequiredParams("symbol", "side", "type", "quantity", "price", "timeInForce"),
		request.SetParam("symbol", "ADAUSDT"),
		request.SetParam("side", "BUY"),
		request.SetParam("type", "LIMIT"),
		request.SetParam("price", "0.51"),
		request.SetParam("quantity", "500"),
		request.SetParam("timeInForce", "GTC"),
	)

	assert.NoError(err)
}
