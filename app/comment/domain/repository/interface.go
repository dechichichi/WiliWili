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
	DelCommentCache(ctx context.Context, id string) error
	GetCommentList(ctx context.Context, id string, page, pageSize, commenttype int64) ([]*model.Comment, error)
	SetCommentList(ctx context.Context, id string, commenttype int64, models []*model.Comment) error
}
