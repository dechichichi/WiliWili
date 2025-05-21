package rpc

import (
	"context"
	"wiliwili/app/like/domain/model"
	"wiliwili/app/like/usecase"
	"wiliwili/kitex_gen/like"
)

type LikeHandler struct {
	useCase usecase.LikeUseCase
}

func (l *LikeHandler) LikeComment(ctx context.Context, req *like.LikeCommentReq) (r *like.LikeCommentResp, err error) {
	r = new(like.LikeCommentResp)
	comment_like := &model.CommentLike{
		CommentID: req.CommentId,
	}

	if err = l.useCase.LikeComment(ctx, comment_like, req.IsLike); err != nil {
		return
	}
	return
}

func (l *LikeHandler) LikeVideo(ctx context.Context, req *like.LikeVideoReq) (r *like.LikeVideoResp, err error) {
	r = new(like.LikeVideoResp)
	video_like := &model.VideoLike{
		VideoID: req.VideoId,
	}
	if err = l.useCase.LikeVideo(ctx, video_like, req.IsLike); err != nil {
		return
	}
	return
}

func (l *LikeHandler) CommentLikeNum(ctx context.Context, req *like.CommentLikeNumReq) (r *like.CommentLikeNumResp, err error) {
	r = new(like.CommentLikeNumResp)
	var num int64
	if num, err = l.useCase.CommentLikeNum(ctx, req.CommentId); err != nil {
		return
	}
	r.TotalCount = num
	return
}

func (l *LikeHandler) VideoLikeNum(ctx context.Context, req *like.VideoLikeNumReq) (r *like.VideoLikeNumResp, err error) {
	r = new(like.VideoLikeNumResp)
	var num int64
	if num, err = l.useCase.VideoLikeNum(ctx, req.VideoId); err != nil {
		return
	}
	r.TotalCount = num
	return
}

func NewLikeHandler(useCase usecase.LikeUseCase) *LikeHandler {
	return &LikeHandler{useCase}
}
