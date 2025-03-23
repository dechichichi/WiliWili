package repository

import (
	"context"
	"wiliwili/app/comment/domain/model"
)

type CommentDB interface {
	CreateComment(ctx context.Context, comment *model.Comment) error
	GetCommentList(ctx context.Context, videoID string, page, pageSize, commenttype int64) ([]*model.Comment, error)
	DeleteComment(ctx context.Context, commentID string, uid string) error
}

type CommentCache interface {
}
