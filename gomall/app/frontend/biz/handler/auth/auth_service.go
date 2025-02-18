package auth

import (
	"context"
	"github.com/yqihe/91-mall/gomall/app/frontend/biz/pack"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/yqihe/91-mall/gomall/app/frontend/biz/service"
	auth "github.com/yqihe/91-mall/gomall/app/frontend/hertz_gen/frontend/auth"
)

// Register .
// @router /admin/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := &auth.RegisterResp{}
	resp, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}
