package redis

import (
	"context"
	"encoding/json"
	"time"
	"wiliwili/app/user/domain/model"
)

func (c *UserDBCache) GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	val, err := c.client.Get(ctx, "user_profile_"+string(uid)).Result()
	if err != nil {
		return nil, err
	}
	var userProfile model.UserProfile
	err = json.Unmarshal([]byte(val), &userProfile)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (c *UserDBCache) StoreUserProFile(ctx context.Context, uid int64, profile *model.UserProfile) error {
	val, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, "user_profile_"+string(uid), val, 24*time.Hour).Err()
}
