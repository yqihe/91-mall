// Code generated by Kitex v0.9.1. DO NOT EDIT.
package useradminservice

import (
	server "github.com/cloudwego/kitex/server"
	user_admin "github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen/user-admin"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user_admin.UserAdminService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler user_admin.UserAdminService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
