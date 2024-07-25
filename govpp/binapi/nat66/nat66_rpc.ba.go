// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nat66

import (
	"context"
	"fmt"
	"govpp/binapi/memclnt"
	"io"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service nat66.
type RPCService interface {
	Nat66AddDelInterface(ctx context.Context, in *Nat66AddDelInterface) (*Nat66AddDelInterfaceReply, error)
	Nat66AddDelStaticMapping(ctx context.Context, in *Nat66AddDelStaticMapping) (*Nat66AddDelStaticMappingReply, error)
	Nat66InterfaceDump(ctx context.Context, in *Nat66InterfaceDump) (RPCService_Nat66InterfaceDumpClient, error)
	Nat66PluginEnableDisable(ctx context.Context, in *Nat66PluginEnableDisable) (*Nat66PluginEnableDisableReply, error)
	Nat66StaticMappingDump(ctx context.Context, in *Nat66StaticMappingDump) (RPCService_Nat66StaticMappingDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Nat66AddDelInterface(ctx context.Context, in *Nat66AddDelInterface) (*Nat66AddDelInterfaceReply, error) {
	out := new(Nat66AddDelInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat66AddDelStaticMapping(ctx context.Context, in *Nat66AddDelStaticMapping) (*Nat66AddDelStaticMappingReply, error) {
	out := new(Nat66AddDelStaticMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat66InterfaceDump(ctx context.Context, in *Nat66InterfaceDump) (RPCService_Nat66InterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat66InterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat66InterfaceDumpClient interface {
	Recv() (*Nat66InterfaceDetails, error)
	api.Stream
}

type serviceClient_Nat66InterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat66InterfaceDumpClient) Recv() (*Nat66InterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat66InterfaceDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat66PluginEnableDisable(ctx context.Context, in *Nat66PluginEnableDisable) (*Nat66PluginEnableDisableReply, error) {
	out := new(Nat66PluginEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat66StaticMappingDump(ctx context.Context, in *Nat66StaticMappingDump) (RPCService_Nat66StaticMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat66StaticMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat66StaticMappingDumpClient interface {
	Recv() (*Nat66StaticMappingDetails, error)
	api.Stream
}

type serviceClient_Nat66StaticMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat66StaticMappingDumpClient) Recv() (*Nat66StaticMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat66StaticMappingDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
