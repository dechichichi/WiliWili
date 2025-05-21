package redis

import (
	"context"
	"fmt"
	"wiliwili/app/video/domain/model"
)

func (c *VideoDBCache) GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error) {
	key := fmt.Sprintf("video:%s", videoid)
	var video model.VideoProfile
	if err := c.client.Get(ctx, key).Scan(&video); err != nil {
		return nil, err
	}
	return &video, nil
}

func (c *VideoDBCache) SetVideo(ctx context.Context, video *model.VideoProfile) error {
	key := fmt.Sprintf("video:%s", video.VideoID)
	return c.client.Set(ctx, key, video, 0).Err()
}
