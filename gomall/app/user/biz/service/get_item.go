package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user/pack"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

type GetItemService struct {
	ctx context.Context
}

// NewGetItemService new GetItemService
func NewGetItemService(ctx context.Context) *GetItemService {
	return &GetItemService{ctx: ctx}
}

// Run create note info
func (s *GetItemService) Run(req *user.GetItemReq) (resp *user.GetItemResp, err error) {
	// Finish your business logic.
	var adminResp *user.UmsAdmin
	admin, err := mysql.Query.UmsAdmin.
		WithContext(s.ctx).
		Where(mysql.Query.UmsAdmin.ID.Eq(req.Id)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.AdminNotExistedError
		}
		klog.Errorf("[user-service] GetItem 35 failed, err: %s", err.Error())
		return nil, errno.ServiceInternalError
	}
	adminResp = pack.UmsAdmin(admin)
	resp.Data = adminResp
	return resp, nil
}
