package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
)

func Call(ctx context.Context, message *bot.Message, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.Call(ctx, message, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Call call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
