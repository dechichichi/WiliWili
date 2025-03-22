package usecase

import (
	"context"
	"wiliwili/app/video/domain/model"
	"wiliwili/app/video/domain/repository"
	"wiliwili/app/video/domain/service"
)

type VideoUsecase interface {
	SubmitVideo(ctx context.Context, video *model.Video, data []byte) (string, string, error)
	GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error)
	SearchVideo(ctx context.Context, keyword string, pageNum, pageSize int64) ([]*model.VideoProfile, error)
	VideoTrending(ctx context.Context, pageNum, pageSize int64) ([]*model.VideoProfile, error)
}

type useCase struct {
	db    repository.VideoDB
	svc   *service.VideoService
	cache repository.VideoCache
}

func NewVideoUsecase(db repository.VideoDB, svc *service.VideoService, cache repository.VideoCache) VideoUsecase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}
