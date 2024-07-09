package hitokoto

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/hitokoto/rpcservice"
)

type RPCClient interface {
	KitexClient() rpcservice.Client
	Service() string
	Call(ctx context.Context, message *bot.Message, callOptions ...callopt.Option) (r bool, err error)
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
