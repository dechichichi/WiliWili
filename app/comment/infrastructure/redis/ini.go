package redis

import (
	"context"
	"wiliwili/app/comment/domain/repository"

	"github.com/redis/go-redis/v9"
)

type CommentDBCache struct {
	client *redis.Client
}

func NewCommentCache(client *redis.Client) repository.CommentCache {
	return &CommentDBCache{client: client}
}

func (c *CommentDBCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}
