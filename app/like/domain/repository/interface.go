package repository

import (
	"context"
	"wiliwili/app/like/domain/model"
)

type LikeDB interface {
	LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool, uid string) error
	LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool, uid string) error
	CommentLikeNum(ctx context.Context, comment_id string) (int64, error)
	VideoLikeNum(ctx context.Context, video_id string) (int64, error)
}

type LikeCache interface {
}
