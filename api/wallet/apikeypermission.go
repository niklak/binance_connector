package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Get API Key Permission (USER_DATA)

// APIKeyPermissionService get api key permission
//
//gen:new_service
type APIKeyPermissionService struct {
	C *connector.Connector
}

func (s *APIKeyPermissionService) Do(ctx context.Context) (res *APIKeyPermissionResponse, err error) {

	r := request.New("/sapi/v1/account/apiRestrictions",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(APIKeyPermissionResponse)
	err = json.Unmarshal(data, res)
	return
}

// APIKeyPermissionResponse define response of APIKeyPermissionService
type APIKeyPermissionResponse struct {
	IPRestrict                     bool   `json:"ipRestrict"`
	CreateTime                     uint64 `json:"createTime"`
	EnableWithdrawals              bool   `json:"enableWithdrawals"`
	EnableInternalTransfer         bool   `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool   `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool   `json:"enableVanillaOptions"`
	EnableReading                  bool   `json:"enableReading"`
	EnableFutures                  bool   `json:"enableFutures"`
	EnableMargin                   bool   `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool   `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime uint64 `json:"tradingAuthorityExpirationTime"`
}
