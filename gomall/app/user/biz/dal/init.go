package dal

import (
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
