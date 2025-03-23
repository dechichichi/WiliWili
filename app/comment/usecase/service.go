package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/comment/domain/model"
	"wiliwili/pkg/constants"
)

func (s *useCase) CommentVideo(ctx context.Context, comment *model.Comment) (string, error) {
	uid, err := s.svc.CheckUser(ctx)
	if err != nil {
		return "", err
	}
	comment.UserID = fmt.Sprintf("%d", uid)
	comment.CommentType = constants.CommentTypeVideo
	comment.CommentID = s.svc.NewId()
	err = s.db.CreateComment(ctx, comment)
	if err != nil {
		return "", err
	}
	return comment.CommentID, nil
}
func (s *useCase) ReplyComment(ctx context.Context, comment *model.Comment) (string, error) {
	uid, err := s.svc.CheckUser(ctx)
	if err != nil {
		return "", err
	}
	comment.UserID = fmt.Sprintf("%d", uid)
	comment.CommentType = constants.CommentTypeReply
	comment.CommentID = s.svc.NewId()
	err = s.db.CreateComment(ctx, comment)
	if err != nil {
		return "", err
	}
	return comment.CommentID, nil
}
func (s *useCase) GetCommentList(ctx context.Context, videoID string, page, pageSize, commenttype int64) ([]*model.Comment, error) {
	models, err := s.db.GetCommentList(ctx, videoID, page, pageSize, commenttype)
	if err != nil {
		return nil, err
	}
	return models, nil
}
func (s *useCase) DeleteComment(ctx context.Context, commentID string) error {
	uid, err := s.svc.CheckUser(ctx)
	if err != nil {
		return err
	}
	err = s.db.DeleteComment(ctx, commentID, fmt.Sprintf("%d", uid))
	if err != nil {
		return err
	}
	return nil
}
