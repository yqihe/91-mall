package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user_admin/biz/service"
	"github.com/yqihe/91-mall/gomall/app/user_admin/pack"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user_admin"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetItem implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetItem(ctx context.Context, req *user_admin.GetItemReq) (resp *user_admin.GetItemResp, err error) {
	klog.Infof("[user_admin-handler] GetItem, req: %#v", req)
	if req.Id == 0 {
		resp = new(user_admin.GetItemResp)
		resp.Resp = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	resp, err = service.NewGetItemService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user_admin-handler] GetItem, err: %v", err.Error())
		resp = new(user_admin.GetItemResp)
		resp.Resp = pack.BuildBaseResp(err)
		return resp, nil
	}
	return resp, err
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user_admin.RegisterReq) (resp *user.RegisterResp, err error) {
	klog.Infof("[user-handler] Register, req: %#v", req)
	if len(req.Username) == 0 || len(req.Password) == 0 || len(req.Username) > 255 || len(req.Password) > 255 {
		resp = new(user.RegisterResp)
		resp.Resp = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	resp, err = service.NewRegisterService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user-handler] Register, err: %v", err.Error())
		resp = new(user.RegisterResp)
		resp.Resp = pack.BuildBaseResp(err)
		return resp, nil
	}
	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	klog.Infof("[user-handler] Login, req: %#v", req)
	if len(req.Username) == 0 || len(req.Password) == 0 || len(req.Username) > 255 || len(req.Password) > 255 {
		resp = new(user.LoginResp)
		resp.Resp = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	resp, err = service.NewLoginService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user-handler] Login, err: %v", err.Error())
		resp = new(user.LoginResp)
		resp.Resp = pack.BuildBaseResp(err)
		return resp, nil
	}
	return resp, nil
}

// RefreshToken implements the UserServiceImpl interface.
func (s *UserServiceImpl) RefreshToken(ctx context.Context, req *user.RefreshTokenReq) (resp *user.RefreshTokenResp, err error) {
	resp, err = service.NewRefreshTokenService(ctx).Run(req)

	return resp, err
}

// GetAdminInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAdminInfo(ctx context.Context, req *user.GetAdminInfoReq) (resp *user.GetAdminInfoResp, err error) {
	klog.Infof("[user-handler] GetAdminInfo, req: %#v", req)
	resp, err = service.NewGetAdminInfoService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user-handler] GetAdminInfo, err: %v", err.Error())
		resp = new(user.GetAdminInfoResp)
		resp.Resp = pack.BuildBaseResp(err)
		return resp, nil
	}
	return resp, nil
}

// Logout implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}

// List implements the UserServiceImpl interface.
func (s *UserServiceImpl) List(ctx context.Context, req *user.ListReq) (resp *user.ListResp, err error) {
	resp, err = service.NewListService(ctx).Run(req)

	return resp, err
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	resp, err = service.NewUpdatePasswordService(ctx).Run(req)

	return resp, err
}

// Delete implements the UserServiceImpl interface.
func (s *UserServiceImpl) Delete(ctx context.Context, req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// UpdateStatus implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateStatus(ctx context.Context, req *user.UpdateStatusReq) (resp *user.UpdateStatusResp, err error) {
	resp, err = service.NewUpdateStatusService(ctx).Run(req)

	return resp, err
}

// UpdateRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateRole(ctx context.Context, req *user.UpdateRoleReq) (resp *user.UpdateRoleResp, err error) {
	resp, err = service.NewUpdateRoleService(ctx).Run(req)

	return resp, err
}

// GetRoleList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetRoleList(ctx context.Context, req *user.GetRoleListReq) (resp *user.GetRoleListResp, err error) {
	resp, err = service.NewGetRoleListService(ctx).Run(req)

	return resp, err
}
