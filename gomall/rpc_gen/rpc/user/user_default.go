package user

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefreshToken(ctx context.Context, req *user.RefreshTokenReq, callOptions ...callopt.Option) (resp *user.RefreshTokenResp, err error) {
	resp, err = defaultClient.RefreshToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefreshToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAdminInfo(ctx context.Context, req *user.GetAdminInfoReq, callOptions ...callopt.Option) (resp *user.GetAdminInfoResp, err error) {
	resp, err = defaultClient.GetAdminInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAdminInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Logout(ctx context.Context, req *user.LogoutReq, callOptions ...callopt.Option) (resp *user.LogoutResp, err error) {
	resp, err = defaultClient.Logout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func List(ctx context.Context, req *user.ListReq, callOptions ...callopt.Option) (resp *user.ListResp, err error) {
	resp, err = defaultClient.List(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "List call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetItem(ctx context.Context, req *user.GetItemReq, callOptions ...callopt.Option) (resp *user.GetItemResp, err error) {
	resp, err = defaultClient.GetItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *user.UpdateReq, callOptions ...callopt.Option) (resp *user.UpdateResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq, callOptions ...callopt.Option) (resp *user.UpdatePasswordResp, err error) {
	resp, err = defaultClient.UpdatePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *user.DeleteReq, callOptions ...callopt.Option) (resp *user.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateStatus(ctx context.Context, req *user.UpdateStatusReq, callOptions ...callopt.Option) (resp *user.UpdateStatusResp, err error) {
	resp, err = defaultClient.UpdateStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateRole(ctx context.Context, req *user.UpdateRoleReq, callOptions ...callopt.Option) (resp *user.UpdateRoleResp, err error) {
	resp, err = defaultClient.UpdateRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetRoleList(ctx context.Context, req *user.GetRoleListReq, callOptions ...callopt.Option) (resp *user.GetRoleListResp, err error) {
	resp, err = defaultClient.GetRoleList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetRoleList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
