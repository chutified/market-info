package data

import (
	"github.com/chutified/resource-finder/protos/commodity"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// CommodityService defines the service's client and allows
// the rpc calls. The service provides commodity data.
type CommodityService struct {
	client commodity.CommodityClient
	conn   *grpc.ClientConn
}

// New is the CommodityService constructor.
func New() *CommodityService {
	return &CommodityService{}
}

// Init starts the server connection and enables rpc calls.
func (cs *CommodityService) Init(target string) error {

	// get connection
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "unable to dial commodity service")
	}

	// register client
	cs.client = commodity.NewCommodityClient(conn)
	cs.conn = conn
	return nil
}

// Close cancels the connection between the client and the server.
func (cs *CommodityService) Close() error {
	return cs.conn.Close()
}
