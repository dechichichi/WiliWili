package redis

import (
	"context"
	"encoding/json"
	"time"
	"wiliwili/app/user/domain/model"
)

func (c *UserDBCache) GetUser(ctx context.Context, uid int64) (*model.User, error) {
	val, err := c.client.Get(ctx, "user"+string(uid)).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserDBCache) StoreUser(ctx context.Context, uid int64, user *model.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, "user"+string(uid), val, 24*time.Hour).Err()
}
