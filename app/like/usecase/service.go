package usecase

import (
	"context"
	"wiliwili/app/like/domain/model"
)

func (uc *useCase) LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool) error {
	//验证用户权限
	err:= uc.svc.IndentifyUser(ctx, comment_like.UserID)
	if err != nil {
		return err
	}
	//更新点赞数
	//更新缓存
	return nil
}
func (uc *useCase) LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool) error {
	panic("implement me")
}
func (uc *useCase) CommentLikeNum(ctx context.Context, comment_id int64) (int64, error) {
	panic("implement me")
}
func (uc *useCase) VideoLikeNum(ctx context.Context, video_id int64) (int64, error) {
	panic("implement me")
}
