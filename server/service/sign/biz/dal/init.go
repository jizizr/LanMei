package dal

import (
	"github.com/jizizr/LanMei/server/service/sign/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/sign/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
