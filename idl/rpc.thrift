namespace go hitokoto

include "bot.thrift"

service RpcService {
    bool Call(1: bot.Message message)
}
