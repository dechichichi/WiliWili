package rpc

import (
	"context"
	"wiliwili/kitex_gen/comment"
	"wiliwili/pkg/base/client"
)

func InitCommentRPC() {
	c, err := client.InitCommentRPC()
	if err != nil {
		panic(err)
	}
	commentClient = *c
}

func CommentVideo(ctx context.Context, req *comment.CommentVideoReq) (response *comment.CommentVideoResp, err error) {
	resp, err := commentClient.CommentVideo(ctx, req)
	return resp, err
}

func ReplyComment(ctx context.Context, req *comment.ReplyCommentReq) (response *comment.ReplyCommentResp, err error) {
	resp, err := commentClient.ReplyComment(ctx, req)
	return resp, err
}

func GetCommentList(ctx context.Context, req *comment.GetCommentListReq) (response *comment.GetCommentListResp, err error) {
	resp, err := commentClient.GetCommentList(ctx, req)
	return resp, err
}

func DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (response *comment.DeleteCommentResp, err error) {
	resp, err := commentClient.DeleteComment(ctx, req)
	return resp, err
}
