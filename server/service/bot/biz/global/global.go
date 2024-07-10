package global

import (
	"github.com/jizizr/LanMei/server/common"
	"github.com/jizizr/LanMei/server/service/bot/conf"
)

var Manager *common.ServiceManager

func init() {
	var err error
	Manager, err = common.NewServiceManager(conf.GetConf().Hertz.RegistryAddress)
	if err != nil {
		panic(err)
	}
}
