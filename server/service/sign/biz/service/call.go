package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/sign/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/sign/biz/util"
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
	text := common.ExtractText(message)
	var point, pointNow, rank int64
	var isAlreadySignedToday bool
	msg := common.NewMsg(message)
	switch text {
	case "签到":
		text = `说明：
发送[]内的文字即可执行对应签到操作
[保底签到]：签到获取5积分
[试试手气]：随机获取积分`
		goto SEND
	case "保底签到":
		point = 5
	case "试试手气":
		point = util.GenerateScore()
	default:
		return
	}

	isAlreadySignedToday, pointNow, rank, err = mysql.SignIn(message.UserId, point)
	if err != nil {
		klog.Error(err)
		return
	}
	if isAlreadySignedToday {
		text = fmt.Sprintf("\n今天已经签到过了，明天再来吧\n目前你积分为%d\n排名第%d位", pointNow, rank)
	} else {
		text = fmt.Sprintf("\n签到成功，获得%d积分\n目前你积分为%d\n排名第%d位", point, pointNow, rank)
	}
SEND:
	msg.Message = text
	msg.At().Reply().SendMessage()
	return
}
