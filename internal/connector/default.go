package connector

import "time"

// Globals
const (
	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

const defaultTimeout = 1 * time.Minute

const defaultApiURL = "https://api.binance.com"
