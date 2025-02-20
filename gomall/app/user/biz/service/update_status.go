package service

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

type UpdateStatusService struct {
	ctx context.Context
} // NewUpdateStatusService new UpdateStatusService
func NewUpdateStatusService(ctx context.Context) *UpdateStatusService {
	return &UpdateStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateStatusService) Run(req *user.UpdateStatusReq) (resp *user.UpdateStatusResp, err error) {
	// Finish your business logic.

	return
}
