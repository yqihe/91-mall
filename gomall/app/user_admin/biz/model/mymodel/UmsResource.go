package mymodel

import (
	"context"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
)

func GetResourceListByAdminId(ctx context.Context, adminId int64) ([]*model.UmsResource, error) {
	q := mysql.Query
	r, ar, rrr, ur := q.UmsRole, q.UmsAdminRoleRelation, q.UmsRoleResourceRelation, q.UmsResource
	var resources []*model.UmsResource
	err := ar.WithContext(ctx).
		Select(ur.ID, ur.CreateTime, ur.Name, ur.URL, ur.Description, ur.CategoryID).
		LeftJoin(r, ar.RoleID.EqCol(r.ID)).
		LeftJoin(rrr, r.ID.EqCol(rrr.RoleID)).
		LeftJoin(ur, ur.ID.EqCol(rrr.ResourceID)).
		Where(ar.AdminID.Eq(adminId)).
		Where(ur.ID.IsNotNull()).
		Group(ur.ID).
		Scan(resources)
	return resources, err
}
