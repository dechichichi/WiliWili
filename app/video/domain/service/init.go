package service

import (
	"sync/atomic"
	"wiliwili/app/video/domain/repository"
)

type VideoService struct {
	db    repository.VideoDB
	cache repository.VideoCache
}

var RedisAvailable atomic.Bool

func NewVideoService(db repository.VideoDB, cache repository.VideoCache) *VideoService {
	if cache == nil {
		panic("videoService`s cache should not be nil")
	}
	if db == nil {
		panic("videoService`s db should not be nil")
	}
	svc := &VideoService{
		db:    db,
		cache: cache,
	}
	return svc
}
