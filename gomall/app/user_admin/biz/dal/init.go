package dal

import (
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
