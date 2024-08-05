package util

import (
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
)

const (
	BaseUrl = "https://v1.hitokoto.cn/"
)

var client = common.DefaultHttpReq(BaseUrl)

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

func GetQuote() (text string, err error) {
	var res quote
	r, err := client.R().SetSuccessResult(&res).Get("/")
	if err != nil {
		return
	}
	if !r.IsSuccessState() {
		klog.Error("send message error ", r.String())
		err = errors.New("request api error")
		return
	}
	text = fmt.Sprintf("『%s』\n—— %s《%s》", res.Hitokoto, res.FromWho, res.From)
	return
}
