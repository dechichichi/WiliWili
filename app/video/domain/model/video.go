package model

import "time"

// Video 定义了视频的基本信息
type Video struct {
	VideoID      string    `json:"video_id" gorm:"table:video;primaryKey;column:video_id"`
	VideoName    string    `json:"video_name" gorm:"column:video_name"`
	VideoURL     string    `json:"video_url" gorm:"column:video_url"`
	VideoDuration string   `json:"video_duration" gorm:"column:video_duration"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// VideoInfo 定义了视频的简化信息，用于返回给客户端
type VideoInfo struct {
	VideoID   string `json:"video_id"`
	VideoName string `json:"video_name"`
}
// VideoProfile 定义了视频的详细信息，用于展示视频的完整信息
type VideoProfile struct {
	VideoID      string `json:"video_id"`
	VideoName    string `json:"video_name"`
	VideoURL     string `json:"video_url"`
	VideoDuration string `json:"video_duration"`
}