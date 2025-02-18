package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yqihe/91-mall/gomall/pkg/errno"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func SendFailResponse(c *app.RequestContext, err error) {
	if err == nil {
		c.JSON(consts.StatusOK, Response{
			Code:    errno.SuccessCode,
			Message: errno.SuccessMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, Response{
		Code:    -1,
		Message: errno.ConvertErr(err).Error(),
	})
}

func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}
