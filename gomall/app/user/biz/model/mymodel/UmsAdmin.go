package mymodel

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/model"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"gorm.io/gorm"
)

func GetAdminByUsername(ctx context.Context, username string) (*model.UmsAdmin, error) {
	admin, err := mysql.Query.UmsAdmin.
		WithContext(ctx).
		Where(mysql.Query.UmsAdmin.Username.Eq(username)).
		First()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNameOrPasswordError
		}
		klog.Errorf("[modedl-UmsAdmin] GetAdminByUsername Query admin failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	return admin, nil
}
