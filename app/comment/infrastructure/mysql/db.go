package mysql

import (
	"context"
	"wiliwili/app/comment/domain/model"
	"wiliwili/app/comment/domain/repository"

	"gorm.io/gorm"
)

type commentDB struct {
	client *gorm.DB
}

func NewCommentDB(client *gorm.DB) repository.CommentDB {
	return &commentDB{
		client: client,
	}
}

func (db *commentDB) CreateComment(ctx context.Context, comment *model.Comment) error {
	err := db.client.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *commentDB) GetCommentList(ctx context.Context, videoID string, page, pageSize, commenttype int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := db.client.Table(model.Comment{}.TableName()).Where("be_comment_id = ? and comment_type = ?", videoID, commenttype).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (db *commentDB) DeleteComment(ctx context.Context, commentID string, uid string) error {
	// 查询评论是否存在，并且 UserID 是否与传入的 uid 相等
	var comment model.Comment
	if err := db.client.Table(model.Comment{}.TableName()).
		Where("id = ? AND user_id = ?", commentID, uid).
		First(&comment).Error; err != nil {
		// 如果查询不到，说明评论不存在或者用户ID不匹配，返回错误
		return err
	}

	// 如果查询到评论，且用户ID匹配，则进行删除操作
	if err := db.client.Table(model.Comment{}.TableName()).
		Where("id = ?", commentID).
		Delete(&model.Comment{}).Error; err != nil {
		return err
	}

	return nil
}