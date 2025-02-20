package mymodel

import (
	"context"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
)

func InsertLoginLog(ctx context.Context, loginLog *model.UmsAdminLoginLog) error {
	return mysql.Query.UmsAdminLoginLog.WithContext(ctx).Create(loginLog)
}
