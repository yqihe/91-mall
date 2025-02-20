package mysql

import (
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/query"
	"github.com/yqihe/91-mall/gomall/app/user_admin/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Query *query.Query
	err   error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	query.SetDefault(DB)
	Query = query.Q
}
