package rpc

import (
	"wiliwili/kitex_gen/comment/commentservice"
	"wiliwili/kitex_gen/like/likeservice"
	"wiliwili/kitex_gen/user/userservice"
	"wiliwili/kitex_gen/video/videoservice"
)

var (
	userClient    userservice.Client
	videoClient   videoservice.Client
	likeClient    likeservice.Client
	commentClient commentservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitCommentRPC()
	InitLikeRPC()
}
