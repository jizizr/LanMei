package mw

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jizizr/LanMei/server/service/bot/biz/utils"
	"github.com/jizizr/LanMei/server/service/bot/conf"
	"strings"
)

func SigAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		sig := ctx.Request.Header.Get("X-Signature")
		if sig == "" || !utils.VerifySignature(ctx.Request.Body(), strings.TrimPrefix(sig, "sha1="), conf.GetConf().Bot.Secret) {
			utils.SendErrResponse(c, ctx, 401, errors.New("unauthorized"))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
