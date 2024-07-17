package dal

import (
	"github.com/jizizr/LanMei/server/service/cut/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/cut/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
