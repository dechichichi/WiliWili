package service

import (
	"sync/atomic"
	"wiliwili/app/like/domain/repository"
)

type LikeService struct {
	db    repository.LikeDB
	cache repository.LikeCache
}

var RedisAvailable atomic.Bool

func NewLikeService(db repository.LikeDB, cache repository.LikeCache) *LikeService {
	if cache == nil {
		panic("likeService`s cache should not be nil")
	}
	if db == nil {
		panic("likeService`s db should not be nil")
	}
	svc := &LikeService{
		db:    db,
		cache: cache,
	}
	return svc
}