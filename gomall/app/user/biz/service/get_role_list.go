package service

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

type GetRoleListService struct {
	ctx context.Context
} // NewGetRoleListService new GetRoleListService
func NewGetRoleListService(ctx context.Context) *GetRoleListService {
	return &GetRoleListService{ctx: ctx}
}

// Run create note info
func (s *GetRoleListService) Run(req *user.GetRoleListReq) (resp *user.GetRoleListResp, err error) {
	// Finish your business logic.

	return
}
