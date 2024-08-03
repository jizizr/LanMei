package dal

import (
	"github.com/jizizr/LanMei/server/service/reply/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/reply/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
