package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	bot2 "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/hitokoto/rpcservice"
	bot "github.com/jizizr/LanMei/server/service/bot/hertz_gen/bot"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
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
	klog.Info("GetMessageService.Run", "msg", req)
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	var msg bot2.Message
	copier.Copy(&msg, req)
	//fmt.Println(msg)
	//c, err := rpcservice.NewClient("hitokoto", client.WithResolver(r))
	//fmt.Println(c.Call(h.Context, &msg))
	//d, err := rpcservice.NewClient("history", client.WithResolver(r))
	//fmt.Println(d.Call(h.Context, &msg))
	c, err := rpcservice.NewClient("hitokoto", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Call(h.Context, &msg)
	if err != nil {
		klog.Error(err)
	}
	c, err = rpcservice.NewClient("history", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Call(h.Context, &msg)
	println(1)
	if err != nil {
		klog.Error(err)
	}
	resp.Success = true
	return
}
