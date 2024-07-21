package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/rust_func"
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
	if message.GroupId == nil {
		return
	}
	if common.IsBot(message.UserId) {
		return
	}
	words, err := mysql.GetWord(*message.GroupId)
	if err != nil {
		klog.Error(err)
		return
	}
	msg := common.NewMsg(message)
	if len(words) == 0 {
		msg.Message = "群里太冷清了，热闹一点吧！"
	} else {
		picBase64 := rust_func.Wcloud(words)
		msg.Message = fmt.Sprintf("[CQ:image,file=base64://%s]", picBase64)
	}
	_, err = msg.SendMessage()
	return
}
