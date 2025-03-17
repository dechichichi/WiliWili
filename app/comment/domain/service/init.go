package service

import (
	"sync/atomic"
	"wiliwili/app/comment/domain/repository"
)

type CommentService struct {
	db    repository.CommentDB
	cache repository.CommentCache
}

var RedisAvailable atomic.Bool

func NewCommentService(db repository.CommentDB, cache repository.CommentCache) *CommentService {
	if cache == nil {
		panic("commentService`s cache should not be nil")
	}
	if db == nil {
		panic("commentService`s db should not be nil")
	}
	svc := &CommentService{
		db:    db,
		cache: cache,
	}
	return svc
}
