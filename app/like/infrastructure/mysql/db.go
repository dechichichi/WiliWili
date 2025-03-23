package mysql

import (
	"context"
	"fmt"
	"wiliwili/app/like/domain/model"
	"wiliwili/app/like/domain/repository"

	"gorm.io/gorm"
)

type likeDB struct {
	client *gorm.DB
}

func NewLikeDB(client *gorm.DB) repository.LikeDB {
	return &likeDB{client: client}
}

func (db *likeDB) LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool, uid string) error {
	if islike {
		// 检查是否存在点赞关系
		var existingLike model.VideoLike
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Where("user_id = ? AND video_id = ?", uid, video_like.VideoID).
			First(&existingLike).Error; err == nil {
			// 如果已经存在点赞关系，返回错误或忽略
			return fmt.Errorf("already liked this video")
		}
		// 如果不存在点赞关系，创建新的点赞记录
		video_like.UserID = uid
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Create(video_like).Error; err != nil {
			return err
		}
		// 更新评论表中的点赞数
		if err := db.client.Table("video").
			Where("video_id = ?", video_like.VideoID).
			Update("likes_count", gorm.Expr("likes_count + 1")).Error; err != nil {
			return err
		}
	} else {
		// 如果 islike 为 false，表示取消点赞
		// 删除点赞记录
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Where("user_id = ? AND video_id = ?", uid, video_like.VideoID).
			Delete(&model.VideoLike{}).Error; err != nil {
			return err
		}

		// 更新评论表中的点赞数
		if err := db.client.Table("video").
			Where("video_id = ?", video_like.VideoID).
			Update("likes_count", gorm.Expr("likes_count - 1")).Error; err != nil {
			return err
		}
	}
	return nil
}
func (db *likeDB) LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool, uid string) error {
	if islike {
		// 检查是否存在点赞关系
		var existingLike model.VideoLike
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Where("user_id = ? AND comment_id = ?", uid, comment_like.CommentID).
			First(&existingLike).Error; err == nil {
			// 如果已经存在点赞关系，返回错误或忽略
			return fmt.Errorf("already liked this comment")
		}
		// 如果不存在点赞关系，创建新的点赞记录
		comment_like.UserID = uid
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Create(comment_like).Error; err != nil {
			return err
		}
		// 更新评论表中的点赞数
		if err := db.client.Table("comments").
			Where("comment_id = ? ", comment_like.CommentID).
			Update("likes_count", gorm.Expr("likes_count + 1")).Error; err != nil {
			return err
		}
	} else {
		// 如果 islike 为 false，表示取消点赞
		// 删除点赞记录
		if err := db.client.Table(model.VideoLike{}.TableName()).
			Where("user_id = ? AND comment_id = ?", uid, comment_like.CommentID).
			Delete(&model.VideoLike{}).Error; err != nil {
			return err
		}

		// 更新评论表中的点赞数
		if err := db.client.Table("comments").
			Where("comment_id = ? ", comment_like.CommentID).
			Update("likes_count", gorm.Expr("likes_count - 1")).Error; err != nil {
			return err
		}
	}
	return nil
}
func (db *likeDB) CommentLikeNum(ctx context.Context, comment_id string) (int64, error) {
	var count int64
	if err := db.client.Table("comments").Where("comment_id = ?", comment_id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (db *likeDB) VideoLikeNum(ctx context.Context, video_id string) (int64, error) {
	var count int64
	if err := db.client.Table("video").Where("video_id = ?", video_id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
