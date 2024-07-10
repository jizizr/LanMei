package rpc

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc/rpcservice"
)

type RPCClient interface {
	KitexClient() rpcservice.Client
	Service() string
	Call(ctx context.Context, message *bot.Message, callOptions ...callopt.Option) (r bool, err error)
	Type(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (r rpc.CmdType, err error)
	Command(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (r *rpc.Cmd, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := rpcservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient rpcservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() rpcservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Call(ctx context.Context, message *bot.Message, callOptions ...callopt.Option) (r bool, err error) {
	return c.kitexClient.Call(ctx, message, callOptions...)
}

func (c *clientImpl) Type(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (r rpc.CmdType, err error) {
	return c.kitexClient.Type(ctx, empty, callOptions...)
}

func (c *clientImpl) Command(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (r *rpc.Cmd, err error) {
	return c.kitexClient.Command(ctx, empty, callOptions...)
}
