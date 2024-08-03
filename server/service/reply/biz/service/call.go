package service

import (
	"context"
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/reply/biz/global"
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
	fmt.Println(123)
	replies := global.ReplyTable.Match(common.ExtractText(message))
	for _, reply := range replies {
		msg := common.NewMsg(message)
		msg.Message = reply
		msg.Reply().SendMessage()
	}
	return
}
