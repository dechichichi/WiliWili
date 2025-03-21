package model

import (
	"time"
)

// Comment 表示评论信息
type Comment struct {
	CommentID      int64     `json:"comment_id" gorm:"primaryKey;autoIncrement:true"` // 评论ID
	CommentType    int       `json:"comment_type" gorm:"notNull"`                     // 评论类型 (1: 评论视频, 2: 评论评论)
	UserID         int64     `json:"user_id" gorm:"notNull"`                          // 用户ID
	BeCommentID    int64     `json:"be_comment_id" gorm:"notNull"`                    // 被评论对象的ID (视频ID或评论ID)
	CommentContent string    `json:"comment" gorm:"type:TEXT;notNull"`                // 评论内容
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`                // 评论时间
}

func (Comment) TableName() string {
	return "comments"
}
