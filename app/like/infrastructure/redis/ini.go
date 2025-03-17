package redis

import (
	"context"
	"wiliwili/app/like/domain/repository"

	"github.com/redis/go-redis/v9"
)

type LikeDBCache struct {
	client *redis.Client
}

func NewLikeCache(client *redis.Client) repository.LikeCache {
	return &LikeDBCache{client: client}
}

func (c *LikeDBCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}
