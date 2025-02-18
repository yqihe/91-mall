package main

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user/biz/service"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

var (
	ErrInvalidRequest = errors.New("无效的请求参数")
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetItem implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetItem(ctx context.Context, req *user.GetItemReq) (resp *user.GetItemResp, err error) {
	klog.Infof("[user-handler] GetItem, req: %#v", req)
	if req.Id == 0 {
		return nil, ErrInvalidRequest
	}
	resp, err = service.NewGetItemService(ctx).Run(req)
	return resp, err
}
