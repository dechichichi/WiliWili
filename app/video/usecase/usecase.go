package usecase

import (
	"wiliwili/app/video/domain/repository"
	"wiliwili/app/video/domain/service"
)

type VideoUsecase interface {
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
