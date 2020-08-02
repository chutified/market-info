package data

import (
	currency "github.com/chutified/currencies/protos/currency"
	"google.golang.org/grpc"
)

// CurrencyServicE allow the rpc calls. The service provides currency data.
type CurrencyService struct {
	client currency.CurrencyClient
	conn   *grpc.ClientConn
}

// NewCurrency is the CurrencyService constructor.
func NewCurrency() *CurrencyService {
	return &CurrencyService{}
}
