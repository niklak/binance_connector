package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Cancel OCO (TRADE) (DELETE /api/v3/orderList)
// CancelOCOService cancel OCO order
//
//gen:new_service
type CancelOCOService struct {
	C                 *connector.Connector
	symbol            string
	orderListId       *int
	listClientOrderId *string
	newClientOrderId  *string
}

// Symbol set symbol
func (s *CancelOCOService) Symbol(symbol string) *CancelOCOService {
	s.symbol = symbol
	return s
}

// OrderListId set orderListId
func (s *CancelOCOService) OrderListId(orderListId int) *CancelOCOService {
	s.orderListId = &orderListId
	return s
}

// ListClientId set listClientId
func (s *CancelOCOService) ListClientOrderId(ListClientOrderId string) *CancelOCOService {
	s.listClientOrderId = &ListClientOrderId
	return s
}

// NewClientOrderId set newClientOrderId
func (s *CancelOCOService) NewClientOrderId(newClientOrderId string) *CancelOCOService {
	s.newClientOrderId = &newClientOrderId
	return s
}

// Do send request
func (s *CancelOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OrderOCOResponse, err error) {

	r := request.New(
		"/api/v3/orderList",
		request.Method(http.MethodDelete),
		request.SecType(request.SecTypeSigned),
		request.RequiredParams("symbol"),
		request.RequiredOneOfParams([]string{"orderListId", "listClientOrderId"}),
	)

	r.SetParam("symbol", s.symbol)
	r.SetParam("orderListId", s.orderListId)
	r.SetParam("listClientOrderId", s.listClientOrderId)
	r.SetParam("newClientOrderId", s.newClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderOCOResponse)
	err = json.Unmarshal(data, res)
	return
}
