package dal

import (
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
