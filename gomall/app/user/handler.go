package main

import (
	"github.com/yqihe/91-mall/gomall/app/user/biz/service"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetItem implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetItem(ctx context.Context, req *user.GetItemReq) (resp *user.GetItemResp, err error) {
	resp, err = service.NewGetItemService(ctx).Run(req)

	return resp, err
}
