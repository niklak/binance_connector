# Binance Connector

[![Go Reference](https://pkg.go.dev/badge/github.com/niklak/binance_connector.svg)](https://pkg.go.dev/github.com/niklak/binance_connector)

Based on [binance-connector-go](https://github.com/binance/binance-connector-go).

Right now it contains only a few methods useful only for me.
In the future, I will add the rest methods from the original package.

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
	newOrder, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side("BUY").Type("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("order: %#v\n", newOrder)
}


```


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