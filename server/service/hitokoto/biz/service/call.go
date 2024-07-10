package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imroc/req/v3"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"time"
)

type CallService struct {
	ctx context.Context
} // NewCallService new CallService
func NewCallService(ctx context.Context) *CallService {
	return &CallService{ctx: ctx}
}

const (
	BaseUrl = "https://v1.hitokoto.cn/"
)

var client = req.C().
	SetBaseURL(BaseUrl).
	R().
	SetRetryCount(3).
	SetRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)

type quote struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUid int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

// Run create note info
func (s *CallService) Run(message *bot.Message) (resp bool, err error) {
	// Finish your business logic.
	resp = true
	var res quote
	r, err := client.SetSuccessResult(&res).Get("/")
	if err != nil {
		klog.Error("send message error ", err)
		return
	}
	if !r.IsSuccessState() {
		klog.Error("send message error ", r.String())
		err = errors.New("request api error")
		return
	}
	text := fmt.Sprintf("『%s』\n—— %s《%s》", res.Hitokoto, res.FromWho, res.From)
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
