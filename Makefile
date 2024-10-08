export ROOT_MOD=github.com/jizizr/LanMei
export APP_PATH=server/service
.PHONY: go-builder
go-builder:
	docker build -t my-go-builder .
	docker run -it -d --name go-builder -v .:/app my-go-builder sh
lanmei-base:
	docker build -f Base.Dockerfile -t lanmei-base:1.0 .
.PHONY: bot
bot:
	@cd ${APP_PATH}/bot && go mod init ${ROOT_MOD}/${APP_PATH}/bot && cwgo server --service bot --type HTTP --idl ../../../idl/bot.thrift --module ${ROOT_MOD}/${APP_PATH}/bot -I ../../../idl/ && rm go.mod
.PHONY: rpc
rpc:
	@cd server/rpc_gen && go mod init ${ROOT_MOD}/server/rpc_gen && cwgo client --service rpc --type RPC --module ${ROOT_MOD}/server/rpc_gen --I ../../idl/ --idl ../../idl/rpc.thrift && rm go.mod
.PHONY: hitokoto
hitokoto:
	$(eval NAME := hitokoto)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
.PHONY: history
history:
	$(eval NAME := history)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
cut:
	$(eval NAME := cut)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
wordcloud:
	$(eval NAME := wordcloud)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
limiter:
	$(eval NAME := limiter)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
music:
	$(eval NAME := music)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
reply:
	$(eval NAME := reply)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
sign:
	$(eval NAME := sign)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
rebirth:
	$(eval NAME := rebirth)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
code:
	$(eval NAME := code)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
gpt:
	$(eval NAME := gpt)
	@cd ${APP_PATH} && mkdir -p ${NAME}
	@cd ${APP_PATH}/${NAME} && go mod init ${ROOT_MOD}/${APP_PATH}/${NAME} && cwgo server --type RPC --service rpc --module ${ROOT_MOD}/${APP_PATH}/${NAME} --pass "-use ${ROOT_MOD}/server/rpc_gen/kitex_gen" -I ../../../idl/ --idl ../../../idl/rpc.thrift||true && rm go.mod
