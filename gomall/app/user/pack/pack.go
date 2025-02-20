package pack

import (
	"errors"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
	"github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

func BuildBaseResp(err error) *user.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *user.BaseResponse {
	return &user.BaseResponse{
		Code:    err.ErrorCode,
		Message: err.ErrorMsg,
	}
}
