package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/pkg/utils/auth"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type GetAdminInfoService struct {
	ctx context.Context
} // NewGetAdminInfoService new GetAdminInfoService
func NewGetAdminInfoService(ctx context.Context) *GetAdminInfoService {
	return &GetAdminInfoService{ctx: ctx}
}

// Run create note info
func (s *GetAdminInfoService) Run(req *user_admin.GetAdminInfoReq) (resp *user_admin.GetAdminInfoResp, err error) {
	// Finish your business logic.
	resp = new(user_admin.GetAdminInfoResp)

	// 检验 Token
	username, err := auth.CheckAuth(s.ctx)
	if err != nil {
		klog.Errorf("[user_admin-service] GetAdminInfo CheckAuth failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	umsAdmin, err := getAdminByUsername(s.ctx, username)
	if err != nil {
		klog.Errorf("[user_admin-service] GetAdminInfo getAdminByUsername failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}

	return
}
