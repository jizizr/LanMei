package service

import (
	"context"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
)

type TypeService struct {
	ctx context.Context
} // NewTypeService new TypeService
func NewTypeService(ctx context.Context) *TypeService {
	return &TypeService{ctx: ctx}
}

// Run create note info
func (s *TypeService) Run(empty *rpc.Empty) (resp rpc.CmdType, err error) {
	// Finish your business logic.
	resp = rpc.CmdType_COMMAND
	return
}
