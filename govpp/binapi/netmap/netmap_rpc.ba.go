// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package netmap

import (
	"context"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service netmap.
type RPCService interface {
	NetmapCreate(ctx context.Context, in *NetmapCreate) (*NetmapCreateReply, error)
	NetmapDelete(ctx context.Context, in *NetmapDelete) (*NetmapDeleteReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) NetmapCreate(ctx context.Context, in *NetmapCreate) (*NetmapCreateReply, error) {
	out := new(NetmapCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NetmapDelete(ctx context.Context, in *NetmapDelete) (*NetmapDeleteReply, error) {
	out := new(NetmapDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}