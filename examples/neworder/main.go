package main

import (
	"context"
	"net/http"
	"os"

	bc "github.com/niklak/binance_connector"
	"github.com/niklak/binance_connector/request"
)

func main() {

	apiKey := os.Getenv("BINANCE_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET")

	if apiKey == "" || apiSecret == "" {
		panic("API key or secret is not set")
	}

	client := bc.NewClientWithOptions(
		apiKey,
		apiSecret,
		bc.BaseURL("https://testnet.binance.vision"),
	)

	type response struct{}

	var res response

	err := client.NewService().Do(
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
		request.SetParam("newOrderRespType", "RESULT"),
	)

	if err != nil {
		panic(err)
	}

}
