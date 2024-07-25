// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vxlan_gpe

import (
	"context"
	"fmt"
	"govpp/binapi/memclnt"
	"io"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service vxlan_gpe.
type RPCService interface {
	SwInterfaceSetVxlanGpeBypass(ctx context.Context, in *SwInterfaceSetVxlanGpeBypass) (*SwInterfaceSetVxlanGpeBypassReply, error)
	VxlanGpeAddDelTunnel(ctx context.Context, in *VxlanGpeAddDelTunnel) (*VxlanGpeAddDelTunnelReply, error)
	VxlanGpeAddDelTunnelV2(ctx context.Context, in *VxlanGpeAddDelTunnelV2) (*VxlanGpeAddDelTunnelV2Reply, error)
	VxlanGpeTunnelDump(ctx context.Context, in *VxlanGpeTunnelDump) (RPCService_VxlanGpeTunnelDumpClient, error)
	VxlanGpeTunnelV2Dump(ctx context.Context, in *VxlanGpeTunnelV2Dump) (RPCService_VxlanGpeTunnelV2DumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SwInterfaceSetVxlanGpeBypass(ctx context.Context, in *SwInterfaceSetVxlanGpeBypass) (*SwInterfaceSetVxlanGpeBypassReply, error) {
	out := new(SwInterfaceSetVxlanGpeBypassReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) VxlanGpeAddDelTunnel(ctx context.Context, in *VxlanGpeAddDelTunnel) (*VxlanGpeAddDelTunnelReply, error) {
	out := new(VxlanGpeAddDelTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) VxlanGpeAddDelTunnelV2(ctx context.Context, in *VxlanGpeAddDelTunnelV2) (*VxlanGpeAddDelTunnelV2Reply, error) {
	out := new(VxlanGpeAddDelTunnelV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) VxlanGpeTunnelDump(ctx context.Context, in *VxlanGpeTunnelDump) (RPCService_VxlanGpeTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VxlanGpeTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VxlanGpeTunnelDumpClient interface {
	Recv() (*VxlanGpeTunnelDetails, error)
	api.Stream
}

type serviceClient_VxlanGpeTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_VxlanGpeTunnelDumpClient) Recv() (*VxlanGpeTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VxlanGpeTunnelDetails:
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

func (c *serviceClient) VxlanGpeTunnelV2Dump(ctx context.Context, in *VxlanGpeTunnelV2Dump) (RPCService_VxlanGpeTunnelV2DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VxlanGpeTunnelV2DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VxlanGpeTunnelV2DumpClient interface {
	Recv() (*VxlanGpeTunnelV2Details, error)
	api.Stream
}

type serviceClient_VxlanGpeTunnelV2DumpClient struct {
	api.Stream
}

func (c *serviceClient_VxlanGpeTunnelV2DumpClient) Recv() (*VxlanGpeTunnelV2Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VxlanGpeTunnelV2Details:
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
