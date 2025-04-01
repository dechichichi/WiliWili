package usecase

import (
	"context"
	"wiliwili/app/comment/domain/model"
	"wiliwili/app/comment/domain/repository"
	"wiliwili/app/comment/domain/service"
)

type CommentUseCase interface {
	CommentVideo(ctx context.Context, comment *model.Comment) (string, error)
	ReplyComment(ctx context.Context, comment *model.Comment) (string,error)
	GetCommentList(ctx context.Context, videoID string, page, pageSize ,commenttype int64) ([]*model.Comment, error)
	DeleteComment(ctx context.Context, commentID string) error
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
