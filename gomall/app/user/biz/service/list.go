package service

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

type ListService struct {
	ctx context.Context
} // NewListService new ListService
func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

// Run create note info
func (s *ListService) Run(req *user.ListReq) (resp *user.ListResp, err error) {
	// Finish your business logic.

	return
}
