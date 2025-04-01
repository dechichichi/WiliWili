namespace go chat

struct ChatReq {
   1: required string target_uid,
   2: required string content,
}

struct ChatResp {
}

service ChatService {
    ChatResp Chat(1: ChatReq req)
}