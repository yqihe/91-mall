package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type UpdatePasswordService struct {
	ctx context.Context
} // NewUpdatePasswordService new UpdatePasswordService
func NewUpdatePasswordService(ctx context.Context) *UpdatePasswordService {
	return &UpdatePasswordService{ctx: ctx}
}

// Run create note info
func (s *UpdatePasswordService) Run(req *user_admin.UpdatePasswordReq) (resp *user_admin.UpdatePasswordResp, err error) {
	// Finish your business logic.

	return
}
