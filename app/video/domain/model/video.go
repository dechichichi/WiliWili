package model

import "time"

// Video 定义了视频的基本信息
type Video struct {
	OwnerID       string    `json:"owner_id" gorm:"column:owner_id"`
	VideoID       string    `json:"video_id" gorm:"primaryKey;column:video_id"`
	VideoName     string    `json:"video_name" gorm:"column:video_name"`
	VideoURL      string    `json:"video_url" gorm:"column:video_url"`
	VideoDuration string    `json:"video_duration" gorm:"column:video_duration"`
	LikesCount    int       `json:"likes_count" gorm:"column:likes_count"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// VideoProfile 定义了视频的详细信息，用于展示视频的完整信息
type VideoProfile struct {
	VideoID       string `json:"video_id"`
	VideoName     string `json:"video_name"`
	VideoDuration string `json:"video_duration"`
	LikesCount    int64  `json:"likes_count"`
	VideoURL      string `json:"video_url"`
}

func (Video) TableName() string {
	return "video"
}

func (VideoProfile) TableName() string {
	return "video"
}
