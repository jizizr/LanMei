package bot

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jizizr/LanMei/server/service/bot/biz/service"
	"github.com/jizizr/LanMei/server/service/bot/biz/utils"
	bot "github.com/jizizr/LanMei/server/service/bot/hertz_gen/bot"
)

// GetMessage .
// @router /bot [POST]
func GetMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req bot.Message
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetMessageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
