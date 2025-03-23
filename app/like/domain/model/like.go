package model

import (
	"time"
)

type VideoLike struct {
	UserID    string    `json:"user_id" gorm:"primaryKey"`
	VideoID   string    `json:"video_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (VideoLike) TableName() string {
	return "video_like"
}

type CommentLike struct {
	UserID    string    `json:"user_id" gorm:"primaryKey"`
	CommentID string    `json:"comment_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (CommentLike) TableName() string {
	return "comment_like"
}
