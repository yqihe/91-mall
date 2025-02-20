package mymodel

import (
	"context"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
)

func GetAdminByUsername(ctx context.Context, username string) (*model.UmsAdmin, error) {
	return mysql.Query.UmsAdmin.
		WithContext(ctx).
		Where(mysql.Query.UmsAdmin.Username.Eq(username)).
		First()
}

func GetAdminByID(ctx context.Context, id int64) (*model.UmsAdmin, error) {
	return mysql.Query.UmsAdmin.
		WithContext(ctx).
		Where(mysql.Query.UmsAdmin.ID.Eq(id)).
		First()
}

func InsertAdmin(ctx context.Context, umsAdmin *model.UmsAdmin) error {
	return mysql.Query.UmsAdmin.WithContext(ctx).Create(umsAdmin)
}
