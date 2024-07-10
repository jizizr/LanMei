namespace go bot

struct Sender {
  1: i64 user_id,
  2: string nickname,
  3: string card,
  4: optional string role  // 群聊角色，只有在群聊消息中才有
}

struct MessageData {
  1: string type,
  2: optional string text,
  3: optional string id  // 表情符号 ID
}

struct Message {
  1: i64 self_id,
  2: i64 user_id,
  3: i64 time,
  4: i64 message_id,
  5: i64 message_seq,
  6: i64 real_id,
  7: string message_type,
  8: Sender sender,
  9: string raw_message,
  10: i32 font,
  11: optional string sub_type,
  12: list<MessageData> message,
  13: string message_format,
  14: string post_type,
  15: optional i64 group_id,     // 群聊消息 ID
  16: optional string notice_type,  // 通知类型
  17: optional i64 operator_id    // 操作者 ID
}

struct Response{
    1: bool success
}

service BotService {
    Response GetMessage(1: Message message) (api.post="/bot")
}
