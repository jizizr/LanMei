package init

import (
	"github.com/jizizr/LanMei/server/service/reply/biz/global"
	"github.com/jizizr/LanMei/server/service/reply/biz/util"
	"time"
)

func initToken() {
	for {
		time.Sleep(time.Duration(util.UpdateTenantAccessToken(&global.Token)) * time.Second)
	}
}
