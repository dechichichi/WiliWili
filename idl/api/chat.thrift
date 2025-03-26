namespace go api.chat

struct ChatReq {
   1: required string target_uid,
   2: required string content,
}

struct ChatResp {
}

service Chat {
    ChatResp Chat(1: ChatReq req)(api.get="api/v1/chat")
}