package dal

import (
	"github.com/jizizr/LanMei/server/service/history/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/history/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
