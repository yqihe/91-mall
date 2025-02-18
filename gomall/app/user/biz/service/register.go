package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/pkg/utils"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	var adminResp *user.UmsAdmin
	// 查询是否有相同用户名的用户
	admin, err := mysql.Query.UmsAdmin.
		WithContext(s.ctx).
		Where(mysql.Query.UmsAdmin.Username.Eq(req.Username)).
		First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Errorf("[user-service] Register 30 failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	if admin != nil {
		return nil, errno.AdminExistedError
	}
	// 将密码进行加密操作
	umsAdmin := new(user.UmsAdmin)
	_ = copier.Copy(umsAdmin, admin)
	encodePassword := utils.MD5(req.Password)
	return
}
