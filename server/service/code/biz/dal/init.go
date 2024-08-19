package dal

import (
	"github.com/jizizr/LanMei/server/service/code/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/code/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
