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
	"strings"
)

type GetMessageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetMessageService(Context context.Context, RequestContext *app.RequestContext) *GetMessageService {
	return &GetMessageService{RequestContext: RequestContext, Context: Context}
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
	for _, msg := range req.Message {
		if msg.Type != "text" {
			continue
		}
		command := strings.TrimSpace(*msg.Data.Text)
		var m bot2.Message
		err = copier.Copy(&m, req)
		if command == "/help" {
			textList := make([]string, 0)
			global.Manager.Command.Range(func(key, value interface{}) bool {
				rpcInfo := value.(utils.RpcClientWithDescription)
				textList = append(textList, fmt.Sprintf("/%s —— %s", key, rpcInfo.Desc))
				return true
			})
			if len(textList) == 0 {
				return
			}
			msg := common.NewMsg(&m)
			msg.Message = fmt.Sprintf("下面是支持的命令：\n%s", strings.Join(textList, "\n"))
			_, err = msg.Reply().SendMessage()
			return
		}
		if strings.HasPrefix(command, "/") {
			if err != nil {
				hlog.Error(h.Context, "copier.Copy", "err", err)
				return
			}
			command = strings.SplitN(command, " ", 2)[0]
			_, err = global.Manager.CallCommand(strings.TrimPrefix(command, "/"), &m)
			if err != nil {
				hlog.Error(h.Context, "global.Manager.CallCommand", "err", err)
			}
		} else {
			global.Manager.CallText(&m)
		}
	}
	return
}
