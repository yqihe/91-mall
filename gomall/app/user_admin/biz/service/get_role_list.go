package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type GetRoleListService struct {
	ctx context.Context
} // NewGetRoleListService new GetRoleListService
func NewGetRoleListService(ctx context.Context) *GetRoleListService {
	return &GetRoleListService{ctx: ctx}
}

// Run create note info
func (s *GetRoleListService) Run(req *user_admin.GetRoleListReq) (resp *user_admin.GetRoleListResp, err error) {
	// Finish your business logic.

	return
}
