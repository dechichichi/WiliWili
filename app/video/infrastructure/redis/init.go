package redis

import (
	"context"
	"wiliwili/app/video/domain/repository"

	"github.com/redis/go-redis/v9"
)

type VideoDBCache struct {
	client *redis.Client
}

func NewVideoCache(client *redis.Client) repository.VideoCache {
	return &VideoDBCache{client: client}
}

func (c *VideoDBCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}
