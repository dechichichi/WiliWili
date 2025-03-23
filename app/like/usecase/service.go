package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/like/domain/model"
)

func (uc *useCase) LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool) error {
	uid, err := uc.svc.CheckUser(ctx)
	if err != nil {
		return err
	}
	err = uc.db.LikeComment(ctx, comment_like, islike, fmt.Sprintf("%d", uid))
	return nil
}
func (uc *useCase) LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool) error {
	uid, err := uc.svc.CheckUser(ctx)
	if err != nil {
		return err
	}
	err = uc.db.LikeVideo(ctx, video_like, islike, fmt.Sprintf("%d", uid))
	return nil
}
func (uc *useCase) CommentLikeNum(ctx context.Context, comment_id string) (int64, error) {
	count, err := uc.db.CommentLikeNum(ctx, comment_id)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (uc *useCase) VideoLikeNum(ctx context.Context, video_id string) (int64, error) {
	count, err := uc.db.VideoLikeNum(ctx, video_id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
