package usecase

import (
	"wiliwili/app/comment/domain/repository"
	"wiliwili/app/comment/domain/service"
)

type CommentUseCase interface {
}

type useCase struct {
	db    repository.CommentDB
	svc   *service.CommentService
	cache repository.CommentCache
}

func NewcommentUseCase(db repository.CommentDB, svc *service.CommentService, cache repository.CommentCache) CommentUseCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}
