package model

import (
	"time"
)

// VideoLike 表示用户对视频的点赞关系
type VideoLike struct {
	UserID    int64     `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	VideoID   int64     `json:"video_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// CommentLike 表示用户对评论的点赞关系
type CommentLike struct {
	UserID    int64     `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	CommentID int64     `json:"comment_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// LikeCount 表示对象（视频或评论）的点赞总数
type LikeCount struct {
	ObjectID   int64 `json:"object_id" gorm:"primaryKey;autoIncrement:false"`
	ObjectType int   `json:"object_type" gorm:"primaryKey;autoIncrement:false"` // 1: 视频, 2: 评论
	LikeNum    int   `json:"like_num" gorm:"default:0"`
}

func (VideoLike) TableName() string {
	return "video_likes"
}

func (CommentLike) TableName() string {
	return "comment_likes"
}

func (LikeCount) TableName() string {
	return "like_counts"
}
