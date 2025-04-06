package usecase

import (
	"context"
	"wiliwili/app/like/domain/model"
	"wiliwili/app/like/domain/repository"
	"wiliwili/app/like/domain/service"
)

type LikeUseCase interface {
	LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool) error
	LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool) error
	CommentLikeNum(ctx context.Context, comment_id string) (int64, error)
	VideoLikeNum(ctx context.Context, video_id string) (int64, error)
}

type useCase struct {
	db    repository.LikeDB
	svc   *service.LikeService
}

func NewLikeUseCase(db repository.LikeDB, svc *service.LikeService) LikeUseCase {
	return &useCase{
		db:    db,
		svc:   svc,
	}
}
