package setup

import (
	"github.com/huyinghuan/app/config"
	"github.com/huyinghuan/app/irisapp/database"
)

func Do() {
	conf := config.Get()
	database.InitMysql(&conf.Db)
	database.InitRedis(&conf.Redis)
}
