package main

import (
	"context"
	"os"

	bc "github.com/niklak/binance_connector"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

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

	tickerPrice, err := client.NewTickerPriceService().
		Symbols([]string{"BTCUSDT", "ETHUSDT"}).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	log.Info().Interface("ticker_price", tickerPrice).Msg("")
}
