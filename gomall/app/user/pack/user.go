package pack

import (
	"fmt"
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/model"
	"github.com/yqihe/91-mall/gomall/app/user/biz/model/mymodel"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

func UmsAdmin(data *model.UmsAdmin) *user.UmsAdmin {
	if data == nil {
		return nil
	}
	return &user.UmsAdmin{
		Id:         data.ID,
		Username:   data.Username,
		Password:   data.Password,
		Icon:       data.Icon,
		Email:      data.Email,
		NickName:   data.NickName,
		Note:       data.Note,
		CreateTime: fmt.Sprintf("%v", mymodel.MyTime(data.CreateTime)),
		LoginTime:  fmt.Sprintf("%v", mymodel.MyTime(data.LoginTime)),
		Status:     data.Status,
	}
}
