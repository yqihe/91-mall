.PHONY: gen-client
gen-client:
	@cd rpc_gen && cwgo client --type RPC --server_name ${svc} --module github.com/yqihe/91-mall/gomall/rpc_gen -I ../idl --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server:
	@cd app/${svc} && cwgo server --type RPC --server_name ${svc} --module github.com/yqihe/91-mall/gomall/app/${svc} --pass "-use github.com/yqihe/91-mall/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/${svc}.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --server_name frontend --module github.com/yqihe/91-mall/gomall/app/frontend --idl ../../idl/frontend/${svc}.proto