package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Binance Query OCO (USER_DATA) (GET /api/v3/orderList)
// QueryOCOService query OCO order
type QueryOCOService struct {
	C                 *connector.Connector
	orderListId       *int64
	origClientOrderId *string
}

// OrderListId set orderListId
func (s *QueryOCOService) OrderListId(orderListId int64) *QueryOCOService {
	s.orderListId = &orderListId
	return s
}

// OrigClientOrderId set origClientOrderId
func (s *QueryOCOService) OrigClientOrderId(origClientOrderId string) *QueryOCOService {
	s.origClientOrderId = &origClientOrderId
	return s
}

// Do send request
func (s *QueryOCOService) Do(ctx context.Context, opts ...request.RequestOption) (res *OCOResponse, err error) {
	r := &request.Request{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/orderList",
		SecType:  request.SecTypeSigned,
	}

	r.Init()

	r.SetParam("orderListId", s.orderListId)
	r.SetParam("origClientOrderId", s.origClientOrderId)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OCOResponse)
	err = json.Unmarshal(data, res)
	return
}

// Create OCOResponse
type OCOResponse struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   uint64 `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}
