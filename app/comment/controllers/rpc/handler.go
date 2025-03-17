package rpc

import (
	"context"
	"wiliwili/app/comment/usecase"
	"wiliwili/kitex_gen/comment"
)

type CommentHandler struct {
	useCase usecase.CommentUseCase
}

func (c *CommentHandler) CommentVideo(ctx context.Context, req *comment.CommentVideoReq) (r *comment.CommentVideoResp, err error) {
	return
}

func (c *CommentHandler) ReplyComment(ctx context.Context, req *comment.ReplyCommentReq) (r *comment.ReplyCommentResp, err error) {
	return
}

func (c *CommentHandler) GetVideoCommentList(ctx context.Context, req *comment.GetVideoCommentListReq) (r *comment.GetVideoCommentListResp, err error) {
	return
}

func (c *CommentHandler) GetCommentReplyList(ctx context.Context, req *comment.GetCommentReplyListReq) (r *comment.GetCommentReplyListResp, err error) {
	return
}

func (c *CommentHandler) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (r *comment.DeleteCommentResp, err error) {
	return
}

func NewCommentHandler(useCase usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{useCase}
}
