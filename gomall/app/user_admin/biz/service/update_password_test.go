package service

import (
	"context"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
	"testing"
)

func TestUpdatePassword_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdatePasswordService(ctx)
	// init req and assert value

	req := &user_admin.UpdatePasswordReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
