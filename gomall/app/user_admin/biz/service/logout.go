package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user_admin.LogoutReq) (resp *user_admin.LogoutResp, err error) {
	// Finish your business logic.

	return
}
