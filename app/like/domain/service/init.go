package service

import (
	"sync/atomic"
	"wiliwili/app/like/domain/repository"
)

type LikeService struct {
	db    repository.LikeDB
}

var RedisAvailable atomic.Bool

func NewLikeService(db repository.LikeDB) *LikeService {
	if db == nil {
		panic("likeService`s db should not be nil")
	}
	svc := &LikeService{
		db:    db,
	}
	return svc
}