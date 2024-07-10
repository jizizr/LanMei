package service

import (
	"context"
	"github.com/jizizr/LanMei/server/common"
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
	resp = true
	var msg common.Msg
	msg.MessageType = message.MessageType
	msg.GroupID = message.GroupId
	msg.UserID = message.UserId
	msg.Message = message.RawMessage
	msg.Send()
	return
}
