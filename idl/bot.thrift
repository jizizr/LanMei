namespace go bot

struct User {
  1: i64 self_id
  2: i64 user_id
  3: i64 time
  4: i32 message_id
  5: i32 message_seq
  6: i32 real_id
}

struct Sender {
  1: i64 user_id
  2: string nickname
  3: string card
}

struct MessageData {
  1: string text
}

struct Message {
  1: User user
  2: string message_type
  3: Sender sender
  4: string raw_message
  5: i32 font
  6: string sub_type
  7: list<MessageData> message
  8: string message_format
  9: string post_type
}

struct Response{
    1: bool success
}

service BotService {
    Response GetMessage(1: Message message) (api.post="/bot")
}
