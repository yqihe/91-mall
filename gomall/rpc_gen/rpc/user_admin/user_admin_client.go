package user_admin

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"

	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin/useradminservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() useradminservice.Client
	Service() string
	Register(ctx context.Context, Req *user_admin.RegisterReq, callOptions ...callopt.Option) (r *user_admin.RegisterResp, err error)
	Login(ctx context.Context, Req *user_admin.LoginReq, callOptions ...callopt.Option) (r *user_admin.LoginResp, err error)
	RefreshToken(ctx context.Context, Req *user_admin.RefreshTokenReq, callOptions ...callopt.Option) (r *user_admin.RefreshTokenResp, err error)
	GetAdminInfo(ctx context.Context, Req *user_admin.GetAdminInfoReq, callOptions ...callopt.Option) (r *user_admin.GetAdminInfoResp, err error)
	Logout(ctx context.Context, Req *user_admin.LogoutReq, callOptions ...callopt.Option) (r *user_admin.LogoutResp, err error)
	List(ctx context.Context, Req *user_admin.ListReq, callOptions ...callopt.Option) (r *user_admin.ListResp, err error)
	GetItem(ctx context.Context, Req *user_admin.GetItemReq, callOptions ...callopt.Option) (r *user_admin.GetItemResp, err error)
	Update(ctx context.Context, Req *user_admin.UpdateReq, callOptions ...callopt.Option) (r *user_admin.UpdateResp, err error)
	UpdatePassword(ctx context.Context, Req *user_admin.UpdatePasswordReq, callOptions ...callopt.Option) (r *user_admin.UpdatePasswordResp, err error)
	Delete(ctx context.Context, Req *user_admin.DeleteReq, callOptions ...callopt.Option) (r *user_admin.DeleteResp, err error)
	UpdateStatus(ctx context.Context, Req *user_admin.UpdateStatusReq, callOptions ...callopt.Option) (r *user_admin.UpdateStatusResp, err error)
	UpdateRole(ctx context.Context, Req *user_admin.UpdateRoleReq, callOptions ...callopt.Option) (r *user_admin.UpdateRoleResp, err error)
	GetRoleList(ctx context.Context, Req *user_admin.GetRoleListReq, callOptions ...callopt.Option) (r *user_admin.GetRoleListResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := useradminservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient useradminservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() useradminservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user_admin.RegisterReq, callOptions ...callopt.Option) (r *user_admin.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user_admin.LoginReq, callOptions ...callopt.Option) (r *user_admin.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) RefreshToken(ctx context.Context, Req *user_admin.RefreshTokenReq, callOptions ...callopt.Option) (r *user_admin.RefreshTokenResp, err error) {
	return c.kitexClient.RefreshToken(ctx, Req, callOptions...)
}

func (c *clientImpl) GetAdminInfo(ctx context.Context, Req *user_admin.GetAdminInfoReq, callOptions ...callopt.Option) (r *user_admin.GetAdminInfoResp, err error) {
	return c.kitexClient.GetAdminInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) Logout(ctx context.Context, Req *user_admin.LogoutReq, callOptions ...callopt.Option) (r *user_admin.LogoutResp, err error) {
	return c.kitexClient.Logout(ctx, Req, callOptions...)
}

func (c *clientImpl) List(ctx context.Context, Req *user_admin.ListReq, callOptions ...callopt.Option) (r *user_admin.ListResp, err error) {
	return c.kitexClient.List(ctx, Req, callOptions...)
}

func (c *clientImpl) GetItem(ctx context.Context, Req *user_admin.GetItemReq, callOptions ...callopt.Option) (r *user_admin.GetItemResp, err error) {
	return c.kitexClient.GetItem(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *user_admin.UpdateReq, callOptions ...callopt.Option) (r *user_admin.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdatePassword(ctx context.Context, Req *user_admin.UpdatePasswordReq, callOptions ...callopt.Option) (r *user_admin.UpdatePasswordResp, err error) {
	return c.kitexClient.UpdatePassword(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user_admin.DeleteReq, callOptions ...callopt.Option) (r *user_admin.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateStatus(ctx context.Context, Req *user_admin.UpdateStatusReq, callOptions ...callopt.Option) (r *user_admin.UpdateStatusResp, err error) {
	return c.kitexClient.UpdateStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateRole(ctx context.Context, Req *user_admin.UpdateRoleReq, callOptions ...callopt.Option) (r *user_admin.UpdateRoleResp, err error) {
	return c.kitexClient.UpdateRole(ctx, Req, callOptions...)
}

func (c *clientImpl) GetRoleList(ctx context.Context, Req *user_admin.GetRoleListReq, callOptions ...callopt.Option) (r *user_admin.GetRoleListResp, err error) {
	return c.kitexClient.GetRoleList(ctx, Req, callOptions...)
}
