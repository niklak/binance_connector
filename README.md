# Binance Connector

[![Go Reference](https://pkg.go.dev/badge/github.com/niklak/binance_connector.svg)](https://pkg.go.dev/github.com/niklak/binance_connector)

Based on [binance-connector-go](https://github.com/binance/binance-connector-go).

Right now it contains only a few methods useful only for me.
In the future, I will add the rest methods from the original package.

## Differences between this package and `binance-connector-go`
- only client (connector) and api (services) are exposed.
- Use `zerolog` package instead of standard `log` package
- Add more convenient way to set client timeout
- `TickerPrice.Do` returns a `[]*TickerPriceResponse` instead of `*TickerPriceResponse`
- `Ticker24hr.Do` returns a `[]*Ticker24hrResponse` instead of `*Ticker24hrResponse`
- Reduced `if` blocks for optional parameters, The check performs only at the one place now.
- Added more constructors, which can be convenient if you need to set timeout or http client with your settings.
- `AllOrdersResponse` and `NewOpenOrdersResponse` are combined into one struct `OrderResponse`