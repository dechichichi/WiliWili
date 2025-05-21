package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"wiliwili/app/user/domain/model"
)

// GetUser 获取用户信息
func (c *UserDBCache) GetUser(ctx context.Context, uid int64) (*model.User, error) {
	key := "user:" + strconv.FormatInt(uid, 10)

	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// StoreUser 存储用户信息
func (c *UserDBCache) StoreUser(ctx context.Context, uid int64, user *model.User) error {
	key := "user:" + strconv.FormatInt(uid, 10)

	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, val, 24*time.Hour).Err()
}
