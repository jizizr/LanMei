package util

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/service/reply/biz/model"
	"github.com/jizizr/LanMei/server/service/reply/conf"
)

func UpdateTenantAccessToken(token *string) int {
	tinfo := conf.GetConf().TableAppInfo
	app := model.AuthApp{
		AppID:     tinfo.AppID,
		AppSecret: tinfo.AppSecret,
	}
	// get tenant access token
	var authResp model.AuthResp
	_, err := client.SetBody(app).SetSuccessResult(&authResp).Post("/auth/v3/tenant_access_token/internal/")
	if err != nil {
		klog.Error("get tenant access token error ", err)
		return 1
	}
	*token = authResp.TenantAccessToken
	return authResp.Expire - 15*60 // 15 minutes before expire
}
