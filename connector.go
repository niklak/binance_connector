//go:generate go run internal/cmd/genalias/genalias.go  --src=api/account,api/market,api/spot,api/wallet,api/staking,api/subaccount,api/margin,api/ethstaking --dst-alias=connector_alias.go --dst-service=connector_service.go

package binance_connector

import (
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

type ConnectorOption = connector.ConnectorOption

var BaseURL = connector.BaseURL
var HTTPClient = connector.HTTPClient
var Timeout = connector.Timeout
var TimeOffset = connector.TimeOffset

var ErrRequestCantBeNil = connector.ErrRequestCantBeNil
var ErrMissingParam = request.ErrMissingParam

type Client struct {
	*connector.Connector
}

func NewClient(apiKey, secretKey string, baseURL ...string) *Client {

	var c *connector.Connector
	if len(baseURL) > 0 {
		c = connector.NewConnector(apiKey, secretKey, BaseURL(baseURL[0]))
	} else {
		c = connector.NewConnector(apiKey, secretKey)
	}

	return &Client{Connector: c}
}

func NewClientWithOptions(apiKey, secretKey string, options ...ConnectorOption) *Client {

	c := connector.NewConnector(apiKey, secretKey, options...)

	return &Client{Connector: c}
}
