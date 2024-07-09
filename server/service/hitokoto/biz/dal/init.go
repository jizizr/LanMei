package dal

import (
	"github.com/jizizr/LanMei/server/service/hitokoto/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/hitokoto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
