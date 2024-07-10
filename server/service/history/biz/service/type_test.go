package service

import (
	"context"
	rpc "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/rpc"
	"testing"
)

func TestType_Run(t *testing.T) {
	ctx := context.Background()
	s := NewTypeService(ctx)
	// init req and assert value

	empty := &rpc.Empty{}
	resp, err := s.Run(empty)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
