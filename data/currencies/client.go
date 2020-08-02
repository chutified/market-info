package data

import (
	currency "github.com/chutified/currencies/protos/currency"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// CurrencyService allow the rpc calls. The service provides currency data.
type CurrencyService struct {
	client currency.CurrencyClient
	conn   *grpc.ClientConn
}

// NewCurrency is the CurrencyService constructor.
func NewCurrency() *CurrencyService {
	return &CurrencyService{}
}

func (cs *CurrencyService) Init(target string) error {

	// get connection
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "unable to dial ")
	}

	// register client
	cs.client = currency.NewCurrencyClient(conn)
	cs.conn = conn
	return nil
}

// Close cancels the connection between the client and the server.
func (cs *CurrencyService) Close() error {
	return cs.conn.Close()
}
