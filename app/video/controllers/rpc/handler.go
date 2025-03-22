package rpc

import (
	"context"
	"wiliwili/app/video/controllers/rpc/pack"
	"wiliwili/app/video/domain/model"
	"wiliwili/app/video/usecase"
	"wiliwili/kitex_gen/video"
	"wiliwili/pkg/base"
)

type VideoHandler struct {
	useCase usecase.VideoUsecase
}

func (u *VideoHandler) VideoSubmission(ctx context.Context, req *video.VideoSubmissionReq) (r *video.VideoSubmissionResp, err error) {
	r = new(video.VideoSubmissionResp)
	video := &model.Video{
		VideoName:     req.VideoName,
		VideoDuration: req.VideoDuration,
	}
	videoid, url, err := u.useCase.SubmitVideo(ctx, video, req.Video)
	r.Baseresp = base.BuildBaseResp(err)
	r.VideoId = videoid
	r.VideoUrl = url
	return
}

func (u *VideoHandler) VideoGet(ctx context.Context, req *video.VideoGetReq) (r *video.VideoGetResp, err error) {
	r = new(video.VideoGetResp)
	video, err := u.useCase.GetVideo(ctx, req.VideoId)
	r.Baseresp = base.BuildBaseResp(err)
	r.Video = pack.ToVideoProfile(video)
	return
}

func (u *VideoHandler) VideoSearch(ctx context.Context, req *video.VideoSearchReq) (r *video.VideoSearchResp, err error) {
	r = new(video.VideoSearchResp)
	videos, err := u.useCase.SearchVideo(ctx, req.Keyword, req.PageNum, req.PageSize)
	r.Baseresp = base.BuildBaseResp(err)
	r.Videos = pack.ToVideoProfileList(videos)
	return
}

func (u *VideoHandler) VideoTrending(ctx context.Context, req *video.VideoTrendingReq) (r *video.VideoTrendingResp, err error) {
	r = new(video.VideoTrendingResp)
	videos, err := u.useCase.VideoTrending(ctx, req.PageNum, req.PageSize)
	r.Baseresp = base.BuildBaseResp(err)
	r.Videos = pack.ToVideoProfileList(videos)
	return
}

func NewVideoHandler(useCase usecase.VideoUsecase) *VideoHandler {
	return &VideoHandler{useCase}
}
