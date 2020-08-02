package data

import (
	crypto "github.com/chutified/crypto-currencies/protos/crypto"
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
