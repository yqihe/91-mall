package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	redisClient "github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/redis"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/mymodel"
	"github.com/yqihe/91-mall/gomall/pkg/constants"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"gorm.io/gorm"
)

func getAdminByUsername(ctx context.Context, username string) (*model.UmsAdmin, error) {
	// 先从缓存中获取数据
	key := constants.REDIS_DATABASE_UMS + ":" + constants.REDIS_KEY_ADMIN + ":" + username
	adminBytes, err := redisClient.GetDataFromCache(ctx, key)
	if err == nil && len(adminBytes) > 0 {
		var admin *model.UmsAdmin
		if err = json.Unmarshal(adminBytes, &admin); err != nil {
			klog.Errorf("[user_admin-service] getAdminByUsername Unmarshal admin failed, err: %s", err.Error())
			return nil, errno.ServiceInternalError
		}
		return admin, nil
	}
	// 缓存中没有从数据库中获取
	admin, err := mymodel.GetAdminByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNameOrPasswordError
		}
		klog.Errorf("[user_admin-service] getAdminByUsername GetAdminByUsername failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	// 将数据库中的数据存入缓存中
	adminBytes, _ = json.Marshal(admin)
	if err = redisClient.SetCache(ctx, key, adminBytes); err != nil {
		klog.Errorf("[user_admin-service] getAdminByUsername Redis Set admin failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	return admin, nil
}
