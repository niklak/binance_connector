package wallet

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/internal/request"
)

// Trade Fee (USER_DATA)

// TradeFeeService trade fee
type TradeFeeService struct {
	C      *connector.Connector
	symbol *string
}

// Symbol set symbol
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = &symbol
	return s
}

func (s *TradeFeeService) Do(ctx context.Context) (res []*TradeFeeResponse, err error) {

	r := request.New("/sapi/v1/asset/tradeFee",
		request.Method(http.MethodGet),
		request.SecType(request.SecTypeSigned),
	)

	r.SetParam("symbol", s.symbol)
	data, err := s.C.CallAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = make([]*TradeFeeResponse, 0)
	err = json.Unmarshal(data, &res)
	return
}

// TradeFeeResponse define response of TradeFeeService
type TradeFeeResponse struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
