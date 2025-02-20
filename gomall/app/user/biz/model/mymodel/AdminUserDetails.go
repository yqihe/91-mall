package mymodel

import (
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/model"
)

type AdminUserDetails struct {
	UmsAdmin     *model.UmsAdmin
	ResourceList []*model.UmsResource
}

func (ud *AdminUserDetails) IsEnabled() bool {
	return ud.UmsAdmin.Status == 1
}
