package rpc

import (
	"context"
	"wiliwili/kitex_gen/video"
	"wiliwili/pkg/base/client"
)

func InitVideoRPC() {
	c, err := client.InitVideoRPC()
	if err != nil {
		panic(err)
	}
	videoClient = *c
}

func SubmitVideo(ctx context.Context, req *video.VideoSubmissionReq) (reponse *video.VideoSubmissionResp, err error) {
	resp, err := videoClient.VideoSubmission(ctx, req)
	return resp, err
}

func GetVideo(ctx context.Context, req *video.VideoGetReq) (reponse *video.VideoGetResp, err error) {
	resp, err := videoClient.VideoGet(ctx, req)
	return resp, err
}

func SearchVideo(ctx context.Context, req *video.VideoSearchReq) (reponse *video.VideoSearchResp, err error) {
	resp, err := videoClient.VideoSearch(ctx, req)
	return resp, err
}

func VideoTrending(ctx context.Context, req *video.VideoTrendingReq) (reponse *video.VideoTrendingResp, err error) {
	resp, err := videoClient.VideoTrending(ctx, req)
	return resp, err
}
