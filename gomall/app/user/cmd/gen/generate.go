package main

import (
	"fmt"
	"github.com/yqihe/91-mall/gomall/app/user/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../biz/model/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	db, err := gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		/*
			记得修改 conf 里的 initConf() 里的 prefix 为 "../../conf"
		*/
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
