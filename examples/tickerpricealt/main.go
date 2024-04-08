package main

import (
	"context"
	"encoding/json"
	"os"

	bc "github.com/niklak/binance_connector"
	"github.com/niklak/binance_connector/request"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var err error

	apiKey := os.Getenv("BINANCE_KEY")
	apiSecret := os.Getenv("BINANCE_SECRET")

	// some of binance api endpoints does not require authentication such as price ticker

	client := bc.NewClientWithOptions(
		apiKey,
		apiSecret,
		bc.BaseURL("https://testnet.binance.vision"),
	)

	type symbolPrice struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}

	var tickerPrice []*symbolPrice

	symbols := []string{"BTCUSDT", "ETHUSDT"}

	symbolsVal, err := json.Marshal(&symbols)

	if err != nil {
		panic(err)
	}

	err = client.NewService().Do(
		context.Background(),
		"/api/v3/ticker/price",
		&tickerPrice,
		request.SetParam("symbols", string(symbolsVal)),
	)

	if err != nil {
		panic(err)
	}

	log.Info().Interface("ticker_price", tickerPrice).Msg("")
}
