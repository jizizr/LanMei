package service

import (
	"context"
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/limiter/biz/util"
	"time"
)

var limiter = util.NewLimiter(1500*time.Millisecond, 3)

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
	if limiter.IsOver(message.UserId) {
		msg := common.NewMsg(message)
		msg.SendBan(60)
		msg.Message = fmt.Sprintf("刷屏？")
		msg.Reply().SendMessage()
	}
	return
}
