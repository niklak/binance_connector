package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/niklak/binance_connector/api/apierrors"
	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// User Universal Transfer (USER_DATA)

// UserUniversalTransferService user universal transfer
type UserUniversalTransferService struct {
	C            *connector.Connector
	transferType string
	asset        string
	amount       float64
	fromSymbol   *string
	toSymbol     *string
}

// TransferType set transferType
func (s *UserUniversalTransferService) TransferType(transferType string) *UserUniversalTransferService {
	s.transferType = transferType
	return s
}

// Asset set asset
func (s *UserUniversalTransferService) Asset(asset string) *UserUniversalTransferService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *UserUniversalTransferService) Amount(amount float64) *UserUniversalTransferService {
	s.amount = amount
	return s
}

// FromSymbol set fromSymbol
func (s *UserUniversalTransferService) FromSymbol(fromSymbol string) *UserUniversalTransferService {
	s.fromSymbol = &fromSymbol
	return s
}

// ToSymbol set toSymbol
func (s *UserUniversalTransferService) ToSymbol(toSymbol string) *UserUniversalTransferService {
	s.toSymbol = &toSymbol
	return s
}

func (s *UserUniversalTransferService) Do(ctx context.Context) (res *UserUniversalTransferResponse, err error) {
	r := &request.Request{
		Method:   http.MethodPost,
		Endpoint: "/sapi/v1/asset/transfer",
		SecType:  request.SecTypeSigned,
	}
	r.Init()

	if s.transferType == "" {
		return nil, fmt.Errorf("%w: transferType", apierrors.ErrMissingParameter)
	}
	if s.asset == "" {
		return nil, fmt.Errorf("%w: asset", apierrors.ErrMissingParameter)
	}

	if s.amount <= 0 {
		return nil, fmt.Errorf("%w: amount", apierrors.ErrMissingParameter)
	}
	r.SetParam("type", s.transferType)
	r.SetParam("asset", s.asset)
	r.SetParam("amount", s.amount)

	r.SetParam("fromSymbol", s.fromSymbol)
	r.SetParam("toSymbol", s.toSymbol)

	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(UserUniversalTransferResponse)
	err = json.Unmarshal(data, res)
	return
}

// UserUniversalTransferResponse define response of UserUniversalTransferService
type UserUniversalTransferResponse struct {
	TranId int64 `json:"tranId"`
}
