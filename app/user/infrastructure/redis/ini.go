package redis

import (
	"context"
	"wiliwili/app/user/domain/repository"

	"github.com/redis/go-redis/v9"
)

type commodityCache struct {
	client *redis.Client
}

func NewCommodityCache(client *redis.Client) repository.UserCache {
	return &commodityCache{client: client}
}

func (c *commodityCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}
