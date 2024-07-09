export ROOT_MOD=github.com/jizizr/LanMei
export APP_PATH=server/service
.PHONY: bot
bot:
	@cd ${APP_PATH}/bot && cwgo server --service bot --type HTTP --idl ../../../idl/bot.thrift --module ${ROOT_MOD}/${APP_PATH}/bot -I ../../../idl/
.PHONY: hitokoto
hitokoto:
	@cd ${APP_PATH}/hitokoto && go mod init ${ROOT_MOD}/${APP_PATH}/hitokoto && cwgo server --type RPC --service hitokoto --module ${ROOT_MOD}/${APP_PATH}/hitokoto --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
	@cd server/rpc_gen && go mod init ${ROOT_MOD}/server/rpc_gen && cwgo client --service hitokoto --type RPC --module ${ROOT_MOD}/server/rpc_gen --I ../../idl/ --idl ../../idl/rpc.thrift && rm go.mod go.sum
.PHONY: history
history:
	@cd ${APP_PATH}/history && go mod init ${ROOT_MOD}/${APP_PATH}/history && cwgo server --type RPC --service history --module ${ROOT_MOD}/${APP_PATH}/history --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
	@cd server/rpc_gen && go mod init ${ROOT_MOD}/server/rpc_gen && cwgo client --service history --type RPC --module ${ROOT_MOD}/server/rpc_gen --I ../../idl/ --idl ../../idl/rpc.thrift && rm go.mod go.sum
