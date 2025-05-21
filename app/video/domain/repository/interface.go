package repository

import (
	"context"
	"wiliwili/app/video/domain/model"
)

type VideoDB interface {
	StoreVideo(ctx context.Context, video *model.Video) error
	GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error)
	SearchVideo(ctx context.Context, keyword string, pageNum, pageSize int64) ([]*model.VideoProfile, error)
	VideoTrending(ctx context.Context, pageNum, pageSize int64) ([]*model.VideoProfile, error)
}

type VideoCache interface {
	GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error)
	SetVideo(ctx context.Context, video *model.VideoProfile) error
}
