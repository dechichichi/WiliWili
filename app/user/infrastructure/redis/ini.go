package redis

import (
	"context"
	"wiliwili/app/user/domain/repository"

	"github.com/redis/go-redis/v9"
)

type UserDBCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) repository.UserCache {
	return &UserDBCache{client: client}
}

func (c *UserDBCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}
