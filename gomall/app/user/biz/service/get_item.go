package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/yqihe/91-mall/gomall/app/user/biz/dal/mysql"
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/mymodel"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

var (
	ErrAdminNotFound   = errors.New("未找到该用户")
	ErrUserServiceFail = errors.New("服务出错")
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
	admin, err := mysql.Query.UmsAdmin.
		WithContext(s.ctx).
		Where(mysql.Query.UmsAdmin.ID.Eq(req.Id)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAdminNotFound
		}
		klog.Errorf("[user-service] GetItem faile, err: %s", err.Error())
		return nil, ErrUserServiceFail
	}
	resp.Data = &user.UmsAdmin{
		Id:         admin.ID,
		Username:   admin.Username,
		Password:   admin.Password,
		Icon:       admin.Icon,
		Email:      admin.Email,
		NickName:   admin.NickName,
		Note:       admin.Note,
		CreateTime: fmt.Sprintf("%v", mymodel.MyTime(admin.CreateTime)),
		LoginTime:  fmt.Sprintf("%v", mymodel.MyTime(admin.LoginTime)),
		Status:     admin.Status,
	}
	return
}
