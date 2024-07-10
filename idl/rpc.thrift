namespace go rpc

include "bot.thrift"

struct Empty {}

enum CmdType {
    COMMAND,
    TEXT,
}

struct Cmd {
    1: string cmd
    2: string description
}

service RpcService {
    bool Call(1: bot.Message message),
    CmdType Type(1:Empty empty),
    Cmd Command(1:Empty empty),
}
