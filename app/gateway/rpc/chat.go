package rpc

import (
	"context"
	"wiliwili/kitex_gen/chat"
	"wiliwili/pkg/base/client"
)

func InitChatRPC() {
	c, err := client.InitChatRPC()
	if err != nil {
		panic(err)
	}
	chatClient = *c
}

func Chat(ctx context.Context, req *chat.ChatReq) (response *chat.ChatResp, err error) {
	resp, err := chatClient.Chat(ctx, req)
	return resp, err
}
