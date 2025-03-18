package rpc

import (
	"context"
	"wiliwili/app/video/usecase"
	"wiliwili/kitex_gen/video"
)

type VideoHandler struct {
	useCase usecase.VideoUsecase
}

func (u *VideoHandler) VideoSubmission(ctx context.Context, req *video.VideoSubmissionReq) (r *video.VideoSubmissionResp, err error) {
	r = new(video.VideoSubmissionResp)
	return
}

func (u *VideoHandler) VideoGet(ctx context.Context, req *video.VideoGetReq) (r *video.VideoGetResp, err error) {
	r = new(video.VideoGetResp)
	return
}

func (u *VideoHandler) VideoSearch(ctx context.Context, req *video.VideoSearchReq) (r *video.VideoSearchResp, err error) {
	r = new(video.VideoSearchResp)
	return
}

func (u *VideoHandler) VideoTrending(ctx context.Context, req *video.VideoTrendingReq) (r *video.VideoTrendingResp, err error) {
	r = new(video.VideoTrendingResp)
	return
}

func NewVideoHandler(useCase usecase.VideoUsecase) *VideoHandler {
	return &VideoHandler{useCase}
}
