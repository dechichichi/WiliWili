package usecase

import (
	"context"
	"wiliwili/app/like/domain/model"
	"wiliwili/app/like/domain/repository"
	"wiliwili/app/like/domain/service"
)

type LikeUseCase interface {
	LikeComment(ctx context.Context, comment_like *model.CommentLike, like_type int64) error
	LikeVideo(ctx context.Context, video_like *model.VideoLike, like_type int64) error
	CommentLikeNum(ctx context.Context, comment_id int64) (int64, error)
	VideoLikeNum(ctx context.Context, video_id int64) (int64, error)
}

type useCase struct {
	db    repository.LikeDB
	svc   *service.LikeService
	cache repository.LikeCache
}

func NewLikeUseCase(db repository.LikeDB, svc *service.LikeService, cache repository.LikeCache) LikeUseCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}
