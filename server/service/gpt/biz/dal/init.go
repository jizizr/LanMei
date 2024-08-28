package dal

import (
	"github.com/jizizr/LanMei/server/service/gpt/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/gpt/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
