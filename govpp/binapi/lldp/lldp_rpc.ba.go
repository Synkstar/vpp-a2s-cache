// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package lldp

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service lldp.
type RPCService interface {
	LldpConfig(ctx context.Context, in *LldpConfig) (*LldpConfigReply, error)
	LldpDump(ctx context.Context, in *LldpDump) (RPCService_LldpDumpClient, error)
	SwInterfaceSetLldp(ctx context.Context, in *SwInterfaceSetLldp) (*SwInterfaceSetLldpReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) LldpConfig(ctx context.Context, in *LldpConfig) (*LldpConfigReply, error) {
	out := new(LldpConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LldpDump(ctx context.Context, in *LldpDump) (RPCService_LldpDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LldpDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LldpDumpClient interface {
	Recv() (*LldpDetails, *LldpDumpReply, error)
	api.Stream
}

type serviceClient_LldpDumpClient struct {
	api.Stream
}

func (c *serviceClient_LldpDumpClient) Recv() (*LldpDetails, *LldpDumpReply, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, nil, err
	}
	switch m := msg.(type) {
	case *LldpDetails:
		return m, nil, nil
	case *LldpDumpReply:
		if err := api.RetvalToVPPApiError(m.Retval); err != nil {
			c.Stream.Close()
			return nil, m, err
		}
		err = c.Stream.Close()
		if err != nil {
			return nil, m, err
		}
		return nil, m, io.EOF
	default:
		return nil, nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceSetLldp(ctx context.Context, in *SwInterfaceSetLldp) (*SwInterfaceSetLldpReply, error) {
	out := new(SwInterfaceSetLldpReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}