package main

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
	"github.com/jizizr/LanMei/server/service/history/biz/service"
)

// RpcServiceImpl implements the last service interface defined in the IDL.
type RpcServiceImpl struct{}

// Call implements the RpcServiceImpl interface.
func (s *RpcServiceImpl) Call(ctx context.Context, message *bot.Message) (resp bool, err error) {
	resp, err = service.NewCallService(ctx).Run(message)

	return resp, err
}

// Type implements the RpcServiceImpl interface.
func (s *RpcServiceImpl) Type(ctx context.Context, empty *rpc.Empty) (resp rpc.CmdType, err error) {
	resp, err = service.NewTypeService(ctx).Run(empty)

	return resp, err
}

// Command implements the RpcServiceImpl interface.
func (s *RpcServiceImpl) Command(ctx context.Context, empty *rpc.Empty) (resp *rpc.Cmd, err error) {
	resp, err = service.NewCommandService(ctx).Run(empty)

	return resp, err
}
