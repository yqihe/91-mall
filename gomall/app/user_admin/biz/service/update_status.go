package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type UpdateStatusService struct {
	ctx context.Context
} // NewUpdateStatusService new UpdateStatusService
func NewUpdateStatusService(ctx context.Context) *UpdateStatusService {
	return &UpdateStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateStatusService) Run(req *user_admin.UpdateStatusReq) (resp *user_admin.UpdateStatusResp, err error) {
	// Finish your business logic.

	return
}
