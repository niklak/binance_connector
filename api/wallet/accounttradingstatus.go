package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Account API Trading Status (USER_DATA)

// AccountApiTradingStatusService account api trading status
type AccountApiTradingStatusService struct {
	C *connector.Connector
}

func (s *AccountApiTradingStatusService) Do(ctx context.Context) (res *AccountApiTradingStatusResponse, err error) {

	r := request.New("/sapi/v1/account/apiTradingStatus",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AccountApiTradingStatusResponse)
	err = json.Unmarshal(data, res)
	return
}

// AccountApiTradingStatusResponse define response of AccountApiTradingStatusService
type AccountApiTradingStatusResponse struct {
	Data struct {
		IsLocked           bool  `json:"isLocked"`
		PlannedRecoverTime int64 `json:"plannedRecoverTime"`
		TriggerCondition   struct {
			GCR  int `json:"GCR"`
			IFER int `json:"IFER"`
			UFR  int `json:"UFR"`
		} `json:"triggerCondition"`
		UpdateTime uint64 `json:"updateTime"`
	} `json:"data"`
}
