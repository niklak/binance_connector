package binanceservice

import (
	"context"
	"encoding/json"

	"github.com/niklak/binance_connector/internal/connector"
	"github.com/niklak/binance_connector/request"
)

type Service struct {
	C *connector.Connector
}

func (s *Service) Do(ctx context.Context, endpoint string, res any, opts ...request.RequestOption) (err error) {
	r := request.New(endpoint, opts...)

	data, err := s.C.CallAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, res)
	return
}
