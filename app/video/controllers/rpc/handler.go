package rpc

import (
	"context"
	"wiliwili/app/video/controllers/rpc/pack"
	"wiliwili/app/video/domain/model"
	"wiliwili/app/video/usecase"
	"wiliwili/kitex_gen/video"
	"go.opentelemetry.io/otel/attribute"
	"wiliwili/app/video/controllers/rpc/otelmetrics"
)

type VideoHandler struct {
	useCase usecase.VideoUsecase
}

func (u *VideoHandler) VideoSubmission(ctx context.Context, req *video.VideoSubmissionReq) (r *video.VideoSubmissionResp, err error) {
	otelmetrics.VideoQPSCounter.Add(ctx, 1, attribute.String("method", "VideoSubmission"))
	r = new(video.VideoSubmissionResp)
	video := &model.Video{
		VideoName:     req.VideoName,
		VideoDuration: req.VideoDuration,
	}
	videoid, url, err := u.useCase.SubmitVideo(ctx, video, req.Video)
	r.VideoId = videoid
	r.VideoUrl = url
	return
}

func (u *VideoHandler) VideoGet(ctx context.Context, req *video.VideoGetReq) (r *video.VideoGetResp, err error) {
	otelmetrics.VideoQPSCounter.Add(ctx, 1, attribute.String("method", "VideoGet"))
	r = new(video.VideoGetResp)
	video, err := u.useCase.GetVideo(ctx, req.VideoId)
	if err != nil {
		return
	}
	r.Video = pack.ToVideoProfile(video)
	return
}

func (u *VideoHandler) VideoSearch(ctx context.Context, req *video.VideoSearchReq) (r *video.VideoSearchResp, err error) {
	otelmetrics.VideoQPSCounter.Add(ctx, 1, attribute.String("method", "VideoSearch"))
	r = new(video.VideoSearchResp)
	videos, err := u.useCase.SearchVideo(ctx, req.Keyword, req.PageNum, req.PageSize)
	r.Videos = pack.ToVideoProfileList(videos)
	return
}

func (u *VideoHandler) VideoTrending(ctx context.Context, req *video.VideoTrendingReq) (r *video.VideoTrendingResp, err error) {
	otelmetrics.VideoQPSCounter.Add(ctx, 1, attribute.String("method", "VideoTrending"))
	r = new(video.VideoTrendingResp)
	videos, err := u.useCase.VideoTrending(ctx, req.PageNum, req.PageSize)
	r.Videos = pack.ToVideoProfileList(videos)
	return
}

func NewVideoHandler(useCase usecase.VideoUsecase) *VideoHandler {
	return &VideoHandler{useCase}
}
