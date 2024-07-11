package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/hitokoto/biz/util"
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
	text, err := util.GetQuote()
	if err != nil {
		klog.Error(err)
	}
	var msg common.Msg
	msg.MessageType = message.MessageType
	msg.UserID = message.Sender.UserId
	msg.GroupID = message.GroupId
	msg.Message = text
	_, err = msg.Send()
	if err != nil {
		klog.Error("send message error ", err, msg)
	}
	return
}
