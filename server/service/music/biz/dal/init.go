package dal

import (
	"github.com/jizizr/LanMei/server/service/music/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/music/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
