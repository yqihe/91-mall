package user

import (
	"context"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetItem(ctx context.Context, req *user.GetItemReq, callOptions ...callopt.Option) (resp *user.GetItemResp, err error) {
	resp, err = defaultClient.GetItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
