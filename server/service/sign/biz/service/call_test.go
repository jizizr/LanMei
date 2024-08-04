package service

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"testing"
)

func TestCall_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCallService(ctx)
	// init req and assert value

	message := &bot.Message{}
	resp, err := s.Run(message)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
