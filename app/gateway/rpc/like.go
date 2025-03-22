package rpc

import (
	"context"
	"wiliwili/kitex_gen/like"
	"wiliwili/pkg/base/client"
)

func InitLikeRPC() {
	c, err := client.InitLikeRPC()
	if err != nil {
		panic(err)
	}
	likeClient = *c
}

func LikeComment(ctx context.Context, req *like.LikeCommentReq) (response *like.LikeCommentResp, err error) {
	resp, err := likeClient.LikeComment(ctx, req)
	return resp, err
}

func LikeVideo(ctx context.Context, req *like.LikeVideoReq) (response *like.LikeVideoResp, err error) {
	resp, err := likeClient.LikeVideo(ctx, req)
	return resp, err
}

func GetCommentLikeNum(ctx context.Context, req *like.CommentLikeNumReq) (response *like.CommentLikeNumResp, err error) {
	resp, err := likeClient.CommentLikeNum(ctx, req)
	return resp, err
}

func GetVideoLikeNum(ctx context.Context, req *like.VideoLikeNumReq) (response *like.VideoLikeNumResp, err error) {
	resp, err := likeClient.VideoLikeNum(ctx, req)
	return resp, err
}
