package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
	"github.com/jizizr/LanMei/server/common"
	bot2 "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/bot/biz/global"
	"github.com/jizizr/LanMei/server/service/bot/biz/utils"
	bot "github.com/jizizr/LanMei/server/service/bot/hertz_gen/bot"
	"strconv"
	"strings"
)

type GetMessageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetMessageService(Context context.Context, RequestContext *app.RequestContext) *GetMessageService {
	return &GetMessageService{RequestContext: RequestContext, Context: Context}
}

func ParseStringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func Help(m *bot2.Message) (err error) {
	textList := make([]string, 0)
	global.Manager.Command.Range(func(key, value interface{}) bool {
		rpcInfo := value.(utils.RpcClientWithDescription)
		textList = append(textList, fmt.Sprintf("/%s —— %s", key, rpcInfo.Desc))
		return true
	})
	if len(textList) == 0 {
		return
	}
	msg := common.NewMsg(m)
	msg.Message = fmt.Sprintf("下面是支持的命令：\n%s", strings.Join(textList, "\n"))
	_, err = msg.Reply().SendMessage()
	return
}

func (h *GetMessageService) Run(req *bot.Message) (resp *bot.Response, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	resp = &bot.Response{Success: true}
	//klog.Info(req)
	if req.Message == nil {
		return
	}
	for i, msg := range req.Message {
		if msg.Type != "text" {
			continue
		}
		command := strings.TrimSpace(*msg.Data.Text)
		var m bot2.Message
		err = copier.Copy(&m, req)
		if strings.HasPrefix(command, "/") || i != 0 && req.Message[i-1].Type == "at" && ParseStringToInt(*req.Message[i-1].Data.Qq) == req.SelfID {
			if err != nil {
				hlog.Error(h.Context, "copier.Copy", "err", err)
				return
			}
			command = strings.TrimSpace(
				strings.TrimPrefix(
					strings.SplitN(command, " ", 2)[0], "/",
				),
			)
			if command == "help" {
				err = Help(&m)
				if err != nil {
					hlog.Error(h.Context, "Help", "err", err)
				}
			} else {
				ok, err := global.Manager.CallCommand(command, &m)
				if err != nil {
					hlog.Error(h.Context, "global.Manager.CallCommand", "err", err)
				}
				if !ok {
					global.Manager.CallText(&m)
				}
			}
		} else {
			global.Manager.CallText(&m)
		}
		return
	}
	return
}
