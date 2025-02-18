package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user/biz/service"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetItem implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetItem(ctx context.Context, req *user.GetItemReq) (resp *user.GetItemResp, err error) {
	klog.Infof("[user-handler] GetItem, req: %#v", req)
	if req.Id == 0 {
		return nil, errno.ParamError
	}
	resp, err = service.NewGetItemService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user-handler] GetItem, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	klog.Infof("[user-handler] Register, req: %#v", req)
	if req.Username == "" || req.Password == "" {
		return nil, errno.ParamError
	}
	resp, err = service.NewRegisterService(ctx).Run(req)
	if err != nil {
		klog.Errorf("[user-handler] Register, err: %v", err.Error())
		return nil, err
	}
	return resp, nil
}
