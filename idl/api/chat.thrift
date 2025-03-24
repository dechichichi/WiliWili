namespace go api.chat

struct ChatReq {
    1: required string uid,
    2: required string msg,
}

struct ChatResp {
    1: required string uid,
    2: required string msg,
    3: required i64 timestamp,
}

service Chat {
    ChatResp Chat(1: ChatReq req)(api.get="api/v1/chat")
}