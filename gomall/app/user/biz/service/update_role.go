package service

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

type UpdateRoleService struct {
	ctx context.Context
} // NewUpdateRoleService new UpdateRoleService
func NewUpdateRoleService(ctx context.Context) *UpdateRoleService {
	return &UpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleService) Run(req *user.UpdateRoleReq) (resp *user.UpdateRoleResp, err error) {
	// Finish your business logic.

	return
}
