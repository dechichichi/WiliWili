package rpc

import (
	"wiliwili/kitex_gen/interactive/interactive"
	"wiliwili/kitex_gen/networking/followservice"
	"wiliwili/kitex_gen/user/userservice"
	"wiliwili/kitex_gen/video/videoservice"
)

var (
	userClient        userservice.Client
	videoClient       videoservice.Client
	networkingClient  followservice.Client
	interactiveClient interactive.Client
)
