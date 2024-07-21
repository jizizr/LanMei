package dal

import (
	"github.com/jizizr/LanMei/server/service/limiter/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/limiter/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
