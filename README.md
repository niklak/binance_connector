# Binance Connector

[![Go Reference](https://pkg.go.dev/badge/github.com/niklak/binance_connector.svg)](https://pkg.go.dev/github.com/niklak/binance_connector)
[![Go](https://github.com/niklak/binance_connector/actions/workflows/go.yml/badge.svg)](https://github.com/niklak/binance_connector/actions/workflows/go.yml)

This package is a wrapper for the Binance API. It is based on the [binance-connector-go](https://github.com/binance/binance-connector-go).

Supported API endpoints:
- Account
- Market Data
- Spot Trading
- Margin Account / Trading
- Sub-Account
- Wallet
- Staking
- Staking (ETH)

## Installation
```
go get github.com/niklak/binance_connector
```

## Examples

### REST API

```go
package main

import (
    "fmt"
    "time"

    bc "github.com/niklak/binance_connector"
)

func main() {
    // with default options
    //client := bc.NewClient("your-api-key", "your-api-secret")

    // with all options
    client := bc.NewClientWithOptions(
        "your-api-key", 
        "your-api-secret",
        bc.BaseURL("https://testnet.binance.vision"),
        bc.Timeout(1 * time.Minute),
        bc.HTTPClient(&http.Client{}),
        bc.TimeOffset(-1000),
	)

    // Set debug with zerolog
    // zerolog.SetGlobalLevel(zerolog.Level(0))

    // Create new order
	newOrder, err := client.NewCreateOrderService().
		Symbol("BTCUSDT").
		Side("BUY").
		Type("MARKET").
		Quantity(0.001).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("order: %#v\n", newOrder)
}


```

## Service for any kind of the Binance API endpoint
Binance API has a lot of endpoints with different parameters. Many endpoints are missing in this package but it has a flexible service that allows you to call any endpoint with functional options.
 - [examples/neworder](./examples/neworder/main.go)
 - [examples/tickerpricealt](./examples/tickerpricealt/main.go)


## Differences between this package and `binance-connector-go`
- Only a client (connector) and api (services) are exposed.
- Using `zerolog` package instead of standard `log` package.
- Added more convenient way to set client timeout
- `TickerPrice.Do` returns a `[]*TickerPriceResponse` instead of `*TickerPriceResponse`
- `Ticker24hr.Do` returns a `[]*Ticker24hrResponse` instead of `*Ticker24hrResponse`
- Reduced `if` blocks for optional parameters, The check performs only at the one place now.
- Added constructor `NewClientWithOptions` which allow to set optional parameters such as instance of `http.Client` or request timeout.
- `AllOrdersResponse` and `NewOpenOrdersResponse` are combined into one struct `OrderResponse`
- If a service name was started with prefix `Get` then this prefix was removed. For example `GetAccountService` was renamed to `AccountService`
- BUSD endpoints are omitted.
- `QuerySubAccountSpotAssetTransferHistoryService`  returns `[]*SubAccountAssetTransferHistoryResponse` instead of `QuerySubAccountSpotAssetTransferHistoryResp`
- (`Get`)`SubAccountDepositHistoryService` returns `[]*SubAccountDepositHistoryResponse` instead of `GetSubAccountDepositHistoryResp`
- This wrapper doesn't include websocket API handlers at all.
- This package doesn't contain pretty-print functions, you can use your favorite package for this purpose.

## Important

- Some endpoints may not work because Binance could remove these endpoints from the API. 
Please visit [binance api docs](https://binance-docs.github.io/apidocs/spot/en/#change-log).

- This package has no tests. Tests will be added in the future.

- This package will not contain websocket API handlers. Because it is a different approach with different dependencies.
I doubt that I will create another package for websocket API. 
