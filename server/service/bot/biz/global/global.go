package global

import (
	"github.com/jizizr/LanMei/server/service/bot/biz/utils"
	"github.com/jizizr/LanMei/server/service/bot/conf"
)

var Manager *utils.ServiceManager

func init() {
	var err error
	Manager, err = utils.NewServiceManager(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		panic(err)
	}
}
