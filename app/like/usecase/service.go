package usecase

import (
	"context"
	"wiliwili/app/like/domain/model"
)

func (uc *useCase) LikeComment(ctx context.Context, comment_like *model.CommentLike, like_type int64) error {
	panic("implement me")
}
func (uc *useCase) LikeVideo(ctx context.Context, video_like *model.VideoLike, like_type int64) error {
	panic("implement me")
}
func (uc *useCase) CommentLikeNum(ctx context.Context, comment_id int64) (int64, error) {
	panic("implement me")
}
func (uc *useCase) VideoLikeNum(ctx context.Context, video_id int64) (int64, error) {
	panic("implement me")
}
