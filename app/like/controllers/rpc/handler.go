package rpc

import (
	"context"
	"wiliwili/app/like/usecase"
	"wiliwili/kitex_gen/like"
)

type LikeHandler struct {
	useCase usecase.LikeUseCase
}

func (l *LikeHandler) LikeComment(ctx context.Context, req *like.LikeCommentReq) (r *like.LikeCommentResp, err error) {
	return
}

func (l *LikeHandler) LikeVideo(ctx context.Context, req *like.LikeVideoReq) (r *like.LikeVideoResp, err error) {
	return
}

func (l *LikeHandler) CommentLikeList(ctx context.Context, req *like.CommentLikeNumReq) (r *like.CommentLikeNumResp, err error) {
	return
}

func (l *LikeHandler) VideoLikeList(ctx context.Context, req *like.VideoLikeNumReq) (r *like.VideoLikeNumResp, err error) {
	return
}

func NewLikeHandler(useCase usecase.LikeUseCase) *LikeHandler {
	return &LikeHandler{useCase}
}
