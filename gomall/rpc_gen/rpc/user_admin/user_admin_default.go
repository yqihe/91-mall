package user_admin

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user_admin.RegisterReq, callOptions ...callopt.Option) (resp *user_admin.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user_admin.LoginReq, callOptions ...callopt.Option) (resp *user_admin.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefreshToken(ctx context.Context, req *user_admin.RefreshTokenReq, callOptions ...callopt.Option) (resp *user_admin.RefreshTokenResp, err error) {
	resp, err = defaultClient.RefreshToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefreshToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAdminInfo(ctx context.Context, req *user_admin.GetAdminInfoReq, callOptions ...callopt.Option) (resp *user_admin.GetAdminInfoResp, err error) {
	resp, err = defaultClient.GetAdminInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAdminInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Logout(ctx context.Context, req *user_admin.LogoutReq, callOptions ...callopt.Option) (resp *user_admin.LogoutResp, err error) {
	resp, err = defaultClient.Logout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func List(ctx context.Context, req *user_admin.ListReq, callOptions ...callopt.Option) (resp *user_admin.ListResp, err error) {
	resp, err = defaultClient.List(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "List call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetItem(ctx context.Context, req *user_admin.GetItemReq, callOptions ...callopt.Option) (resp *user_admin.GetItemResp, err error) {
	resp, err = defaultClient.GetItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *user_admin.UpdateReq, callOptions ...callopt.Option) (resp *user_admin.UpdateResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePassword(ctx context.Context, req *user_admin.UpdatePasswordReq, callOptions ...callopt.Option) (resp *user_admin.UpdatePasswordResp, err error) {
	resp, err = defaultClient.UpdatePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *user_admin.DeleteReq, callOptions ...callopt.Option) (resp *user_admin.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateStatus(ctx context.Context, req *user_admin.UpdateStatusReq, callOptions ...callopt.Option) (resp *user_admin.UpdateStatusResp, err error) {
	resp, err = defaultClient.UpdateStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateRole(ctx context.Context, req *user_admin.UpdateRoleReq, callOptions ...callopt.Option) (resp *user_admin.UpdateRoleResp, err error) {
	resp, err = defaultClient.UpdateRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetRoleList(ctx context.Context, req *user_admin.GetRoleListReq, callOptions ...callopt.Option) (resp *user_admin.GetRoleListResp, err error) {
	resp, err = defaultClient.GetRoleList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetRoleList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
