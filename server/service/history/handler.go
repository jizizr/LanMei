package main

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/history/biz/service"
)

// RpcServiceImpl implements the last service interface defined in the IDL.
type RpcServiceImpl struct{}

// Call implements the RpcServiceImpl interface.
func (s *RpcServiceImpl) Call(ctx context.Context, message *bot.Message) (resp bool, err error) {
	resp, err = service.NewCallService(ctx).Run(message)

	return resp, err
}
