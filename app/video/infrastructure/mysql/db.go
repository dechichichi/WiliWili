package mysql

import (
	"context"
	"wiliwili/app/video/domain/model"
	"wiliwili/app/video/domain/repository"

	"gorm.io/gorm"
)

type videoDB struct {
	client *gorm.DB
}

func NewVideoDB(client *gorm.DB) repository.VideoDB {
	return &videoDB{client: client}
}

func (v *videoDB) StoreVideo(ctx context.Context, video *model.Video) error {
	err := v.client.Create(video).Error
	if err != nil {
		return err
	}
	return nil
}

func (v *videoDB) GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error) {
	var video model.VideoProfile
	err := v.client.WithContext(ctx).Where("video_id = ?", videoid).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (v *videoDB) SearchVideo(ctx context.Context, keyword string, pageNum, pageSize int64) ([]*model.VideoProfile, error) {
	var videos []*model.VideoProfile
	offset := int((pageNum - 1) * pageSize)
	err := v.client.Where("video_name LIKE ?", "%"+keyword+"%").Offset(offset).Limit(int(pageSize)).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *videoDB) VideoTrending(ctx context.Context, pageNum, pageSize int64) ([]*model.VideoProfile, error) {
	var videos []*model.VideoProfile
	offset := int((pageNum - 1) * pageSize)
	err := v.client.Order("likes_count DESC, video_id DESC").Offset(offset).Limit(int(pageSize)).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
