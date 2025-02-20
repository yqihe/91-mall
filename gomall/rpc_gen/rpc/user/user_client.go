package user

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"

	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	RefreshToken(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.RefreshTokenResp, err error)
	GetAdminInfo(ctx context.Context, Req *user.GetAdminInfoReq, callOptions ...callopt.Option) (r *user.GetAdminInfoResp, err error)
	Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error)
	List(ctx context.Context, Req *user.ListReq, callOptions ...callopt.Option) (r *user.ListResp, err error)
	GetItem(ctx context.Context, Req *user.GetItemReq, callOptions ...callopt.Option) (r *user.GetItemResp, err error)
	Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error)
	UpdatePassword(ctx context.Context, Req *user.UpdatePasswordReq, callOptions ...callopt.Option) (r *user.UpdatePasswordResp, err error)
	Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error)
	UpdateStatus(ctx context.Context, Req *user.UpdateStatusReq, callOptions ...callopt.Option) (r *user.UpdateStatusResp, err error)
	UpdateRole(ctx context.Context, Req *user.UpdateRoleReq, callOptions ...callopt.Option) (r *user.UpdateRoleResp, err error)
	GetRoleList(ctx context.Context, Req *user.GetRoleListReq, callOptions ...callopt.Option) (r *user.GetRoleListResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
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
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) RefreshToken(ctx context.Context, Req *user.RefreshTokenReq, callOptions ...callopt.Option) (r *user.RefreshTokenResp, err error) {
	return c.kitexClient.RefreshToken(ctx, Req, callOptions...)
}

func (c *clientImpl) GetAdminInfo(ctx context.Context, Req *user.GetAdminInfoReq, callOptions ...callopt.Option) (r *user.GetAdminInfoResp, err error) {
	return c.kitexClient.GetAdminInfo(ctx, Req, callOptions...)
}

func (c *clientImpl) Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error) {
	return c.kitexClient.Logout(ctx, Req, callOptions...)
}

func (c *clientImpl) List(ctx context.Context, Req *user.ListReq, callOptions ...callopt.Option) (r *user.ListResp, err error) {
	return c.kitexClient.List(ctx, Req, callOptions...)
}

func (c *clientImpl) GetItem(ctx context.Context, Req *user.GetItemReq, callOptions ...callopt.Option) (r *user.GetItemResp, err error) {
	return c.kitexClient.GetItem(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdatePassword(ctx context.Context, Req *user.UpdatePasswordReq, callOptions ...callopt.Option) (r *user.UpdatePasswordResp, err error) {
	return c.kitexClient.UpdatePassword(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateStatus(ctx context.Context, Req *user.UpdateStatusReq, callOptions ...callopt.Option) (r *user.UpdateStatusResp, err error) {
	return c.kitexClient.UpdateStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateRole(ctx context.Context, Req *user.UpdateRoleReq, callOptions ...callopt.Option) (r *user.UpdateRoleResp, err error) {
	return c.kitexClient.UpdateRole(ctx, Req, callOptions...)
}

func (c *clientImpl) GetRoleList(ctx context.Context, Req *user.GetRoleListReq, callOptions ...callopt.Option) (r *user.GetRoleListResp, err error) {
	return c.kitexClient.GetRoleList(ctx, Req, callOptions...)
}
