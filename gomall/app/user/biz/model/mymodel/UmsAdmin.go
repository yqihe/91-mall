package mymodel

import (
	"context"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/model"
)

func GetAdminByUsername(ctx context.Context, username string) (*model.UmsAdmin, error) {
	return mysql.Query.UmsAdmin.
		WithContext(ctx).
		Where(mysql.Query.UmsAdmin.Username.Eq(username)).
		First()
}
