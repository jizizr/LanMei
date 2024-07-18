package dal

import (
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
