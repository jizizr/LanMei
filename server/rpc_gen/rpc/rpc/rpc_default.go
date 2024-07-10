package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
)

func Call(ctx context.Context, message *bot.Message, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.Call(ctx, message, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Call call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Type(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (resp rpc.CmdType, err error) {
	resp, err = defaultClient.Type(ctx, empty, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Type call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Command(ctx context.Context, empty *rpc.Empty, callOptions ...callopt.Option) (resp *rpc.Cmd, err error) {
	resp, err = defaultClient.Command(ctx, empty, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Command call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
