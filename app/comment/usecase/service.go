package usecase

import (
	"context"
	"wiliwili/app/comment/domain/model"
)

func (s *useCase) CommentVideo(ctx context.Context, comment *model.Comment) (int64, error) {
	panic("implement me")
}
func (s *useCase) ReplyComment(ctx context.Context, comment *model.Comment) (int64, error) {
	panic("implement me")
}
func (s *useCase) GetCommentList(ctx context.Context, videoID int64, page, pageSize, commenttype int64) ([]*model.Comment, error) {
	panic("implement me")
}
func (s *useCase) DeleteComment(ctx context.Context, commentID int64, userID int64) error {
	panic("implement me")
}
