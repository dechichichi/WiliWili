package redis

import (
	"context"
	"fmt"
)

func (c *UserDBCache) NewUserId(ctx context.Context) (int64, error) {
	userID, err := c.client.Incr(ctx, "user_id").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get new user id: %w", err)
	}
	return userID, nil
}
