package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/model"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/model/mymodel"
	"github.com/yqihe/91-mall/gomall/app/user_admin/pack"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/pkg/utils"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
	"gorm.io/gorm"
	"time"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user_admin.RegisterReq) (resp *user_admin.RegisterResp, err error) {
	// Finish your business logic.
	resp = new(user_admin.RegisterResp)

	umsAdmin := new(model.UmsAdmin)
	_ = copier.Copy(umsAdmin, req)
	umsAdmin.CreateTime = time.Now()
	umsAdmin.LoginTime = time.Now()
	umsAdmin.Status = 1

	// 查询是否有相同用户名的用户
	admin, err := mymodel.GetAdminByUsername(s.ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Errorf("[user_admin-service] Register GetAdminByUsername failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	if admin != nil {
		return nil, errno.AdminExistedError
	}
	// 将密码进行加密操作
	umsAdmin.Password = utils.MD5(req.Password)
	if err = mymodel.InsertAdmin(s.ctx, umsAdmin); err != nil {
		klog.Errorf("[user_admin-service] Register InsertAdmin failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}

	adminResp := &user_admin.UmsAdmin{}
	if err = copier.Copy(adminResp, umsAdmin); err != nil {
		klog.Errorf("[user_admin-service] Register Copy failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	resp.Resp = pack.BuildBaseResp(nil)
	resp.Data = adminResp
	return resp, nil
}
