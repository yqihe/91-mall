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

func GetResourceListByAdminId(ctx context.Context, adminId int64) ([]*model.UmsResource, error) {
	q := mysql.Query
	r, ar, rrr, ur := q.UmsRole, q.UmsAdminRoleRelation, q.UmsRoleResourceRelation, q.UmsResource
	var resources []*model.UmsResource
	if err := ar.
		WithContext(ctx).
		Select(ur.ID, ur.CreateTime, ur.Name, ur.URL, ur.Description, ur.CategoryID).
		LeftJoin(r, ar.RoleID.EqCol(r.ID)).
		LeftJoin(rrr, r.ID.EqCol(rrr.RoleID)).
		LeftJoin(ur, ur.ID.EqCol(rrr.ResourceID)).
		Where(ar.AdminID.Eq(adminId)).
		Where(ur.ID.IsNotNull()).
		Group(ur.ID).
		Scan(resources); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		klog.Errorf("[model-UmsResource] GetResourceListByAdminId Query resource list failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	return resources, nil
}
