package data

import (
	crypto "github.com/chutified/crypto-currencies/protos/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// CryptoService the rpc calls. The service provides cryptocurrency data.
type CryptoService struct {
	client crypto.CryptoClient
	conn   *grpc.ClientConn
}

// NewCrypto is the CryptoService constructor.
func NewCrypto() *CryptoService {
	return &CryptoService{}
}

// Init starts the server connection and enables rpc calls.
func (cs *CryptoService) Init(target string) error {

	// get connection
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "unable to dial commodity service")
	}

	// register client
	cs.client = crypto.NewCryptoClient(conn)
	cs.conn = conn
	return nil
}

// Close cancels the connection between the client and the server.
func (cs *CryptoService) Close() error {
	return cs.conn.Close()
}
