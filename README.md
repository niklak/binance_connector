# Binance Connector

Based on [binance-connector-go](https://github.com/binance/binance-connector-go).

Right now it contains only a few methods useful only for me.
In the future, I will add the rest methods from the original package.

## Differences between this package and `binance-connector-go`
- only client (connector) and api (services) are exposed.
- Changed constructor `New` -- added optional timeout parameter and binance url required parameter (which still can be empty)
- Use `zerolog` package instead of standard `log` package
- Add more convenient way to set client timeout
- `TickerPrice.Do` returns a `[]*TickerPriceResponse` instead of `*TickerPriceResponse`
- Reduced `if` blocks for optional parameters, The check performs only at one place now. 