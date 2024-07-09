package dal

import (
	"github.com/jizizr/LanMei/server/service/bot/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/bot/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
