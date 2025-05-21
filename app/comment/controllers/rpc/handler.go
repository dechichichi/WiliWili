package rpc

import (
	"context"
	"wiliwili/app/comment/controllers/rpc/pack"
	"wiliwili/app/comment/domain/model"
	"wiliwili/app/comment/usecase"
	"wiliwili/kitex_gen/comment"
	"wiliwili/pkg/constants"
)

type CommentHandler struct {
	useCase usecase.CommentUseCase
}

func (c *CommentHandler) CommentVideo(ctx context.Context, req *comment.CommentVideoReq) (r *comment.CommentVideoResp, err error) {
	r = new(comment.CommentVideoResp)
	comment := &model.Comment{
		BeCommentID:    req.VideoId,
		CommentType:    constants.CommentTypeVideo,
		CommentContent: req.Content,
	}
	commentid, err := c.useCase.CommentVideo(ctx, comment)
	if err != nil {
		return
	}
	r.CommentId = commentid
	return
}

func (c *CommentHandler) ReplyComment(ctx context.Context, req *comment.ReplyCommentReq) (r *comment.ReplyCommentResp, err error) {
	r = new(comment.ReplyCommentResp)
	comment := &model.Comment{
		BeCommentID:    req.CommentId,
		CommentType:    constants.CommentTypeReply,
		CommentContent: req.Content,
	}
	commentid, err := c.useCase.ReplyComment(ctx, comment)
	if err != nil {
		return
	}
	r.CommentId = commentid
	return
}

func (c *CommentHandler) GetCommentList(ctx context.Context, req *comment.GetCommentListReq) (r *comment.GetCommentListResp, err error) {
	r = new(comment.GetCommentListResp)
	commentList, err := c.useCase.GetCommentList(ctx, req.Id, req.Page, req.PageSize, req.CommentTpye)
	if err != nil {
		return
	}
	r.CommentList = *pack.BuildCommentList(commentList)
	return
}

func (c *CommentHandler) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (r *comment.DeleteCommentResp, err error) {
	r = new(comment.DeleteCommentResp)
	err = c.useCase.DeleteComment(ctx, req.CommentId)
	if err != nil {
	}
	return
}

func NewCommentHandler(useCase usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{useCase}
}
