package service

import (
	"context"
	"testing"
	user "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user"
)

func TestUpdateStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateStatusService(ctx)
	// init req and assert value

	req := &user.UpdateStatusReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
