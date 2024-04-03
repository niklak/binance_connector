package wallet

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Daily Account Snapshot (USER_DATA)

// AccountSnapshotService get all orders from account
//
//gen:new_service
type AccountSnapshotService struct {
	C          *connector.Connector
	marketType string
	startTime  *uint64
	endTime    *uint64
	limit      *int
}

// MarketType set market type
func (s *AccountSnapshotService) MarketType(marketType string) *AccountSnapshotService {
	s.marketType = marketType
	return s
}

// StartTime set start time
func (s *AccountSnapshotService) StartTime(startTime uint64) *AccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *AccountSnapshotService) EndTime(endTime uint64) *AccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *AccountSnapshotService) Limit(limit int) *AccountSnapshotService {
	s.limit = &limit
	return s
}

func (s *AccountSnapshotService) Do(ctx context.Context) (res *AccountSnapshotResponse, err error) {

	r := request.New(
		"/sapi/v1/accountSnapshot",
		request.SecType(request.SecTypeSigned),
	)

	if s.marketType == "" {
		err = fmt.Errorf("%w: marketType", apierrors.ErrMissingParameter)
		return
	}

	r.SetParam("type", s.marketType)
	r.SetParam("startTime", s.startTime)
	r.SetParam("endTime", s.endTime)
	r.SetParam("limit", s.limit)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return
	}
	res = new(AccountSnapshotResponse)
	err = json.Unmarshal(data, res)
	return
}

// AccountSnapshotResponse define response of AccountSnapshotService
type AccountSnapshotResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	SnapshotVos []struct {
		Data struct {
			Balances []struct {
				Asset  string `json:"asset"`
				Free   string `json:"free"`
				Locked string `json:"locked"`
			} `json:"balances"`
			TotalAssetOfBtc string `json:"totalAssetOfBtc"`
		} `json:"data"`
		Type       string `json:"type"`
		UpdateTime uint64 `json:"updateTime"`
	} `json:"snapshotVos"`
}
