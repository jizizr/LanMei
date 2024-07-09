package service

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
)

type CallService struct {
	ctx context.Context
} // NewCallService new CallService
func NewCallService(ctx context.Context) *CallService {
	return &CallService{ctx: ctx}
}

// Run create note info
func (s *CallService) Run(message *bot.Message) (resp bool, err error) {
	// Finish your business logic.

	return
}
