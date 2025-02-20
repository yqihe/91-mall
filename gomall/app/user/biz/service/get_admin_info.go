package service

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

type GetAdminInfoService struct {
	ctx context.Context
} // NewGetAdminInfoService new GetAdminInfoService
func NewGetAdminInfoService(ctx context.Context) *GetAdminInfoService {
	return &GetAdminInfoService{ctx: ctx}
}

// Run create note info
func (s *GetAdminInfoService) Run(req *user.GetAdminInfoReq) (resp *user.GetAdminInfoResp, err error) {
	// Finish your business logic.

	return
}
