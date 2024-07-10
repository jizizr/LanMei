package service

import (
	"context"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
)

type CommandService struct {
	ctx context.Context
} // NewCommandService new CommandService
func NewCommandService(ctx context.Context) *CommandService {
	return &CommandService{ctx: ctx}
}

// Run create note info
func (s *CommandService) Run(empty *rpc.Empty) (resp *rpc.Cmd, err error) {
	// Finish your business logic.
	resp = &rpc.Cmd{
		Cmd:         "quote",
		Description: "获取一句随机的名言",
	}
	return
}
