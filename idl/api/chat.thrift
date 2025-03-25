namespace go api.chat



struct ChatMessage {
    1: required string target_uid,
    2: required string content,
}
struct ChatReq {
    1: ChatMessage message,
}

struct ChatResp {
    1: required string uid,
    2: required string content,
    3: required i64 timestamp,
}

service Chat {
    ChatResp Chat(1: ChatReq req)(api.get="api/v1/chat")
}