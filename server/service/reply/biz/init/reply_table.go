package init

import (
	"github.com/jizizr/LanMei/server/service/reply/biz/global"
	"github.com/jizizr/LanMei/server/service/reply/biz/util"
	"time"
)

func initReplyTable() {
	for global.Token == "" {
		time.Sleep(1 * time.Second)
	}
	for {
		util.UpdateReplyTable(global.ReplyTable, &global.Revision, global.Token)
		time.Sleep(30 * time.Second)
	}
}
