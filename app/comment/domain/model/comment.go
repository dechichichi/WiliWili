package model

import (
	"time"
)

// Comment 表示评论信息
type Comment struct {
    CommentID      string    `json:"comment_id" gorm:"primaryKey"`
    CommentType    int       `json:"comment_type" gorm:"notNull"`
    UserID         string    `json:"user_id" gorm:"notNull"`
    BeCommentID    string    `json:"be_comment_id" gorm:"notNull"`
    CommentContent string    `json:"comment" gorm:"column:comment;type:TEXT;notNull"` 
    CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
    LikesCount     int       `json:"likes_count" gorm:"notNull;default:0"`
}

func (Comment) TableName() string {
	return "comments"
}
