namespace go chat


struct ChatReq{
    1: required string uid,
    2: required string msg,
}


struct ChatResp{
    1: required string uid,
    2: required string msg,
    3: required string timestamp,
}

service ChatService {
    ChatResp Chat(1:ChatReq req)
}