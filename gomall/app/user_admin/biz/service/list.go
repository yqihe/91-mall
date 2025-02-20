package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

type ListService struct {
	ctx context.Context
} // NewListService new ListService
func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

// Run create note info
func (s *ListService) Run(req *user_admin.ListReq) (resp *user_admin.ListResp, err error) {
	// Finish your business logic.

	return
}
