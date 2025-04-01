package client

import (
	"context"
	"errors"
	"fmt"
	"wiliwili/config"

	"wiliwili/pkg/errno"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient 传入dbName，具体参考 constants 包
func NewRedisClient(db int) (*redis.Client, error) {
	if config.Redis == nil {
		return nil, errors.New("redis config is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       db,
	})

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errno.Errorf(errno.RedisConnectFailed, fmt.Sprintf("redis connect failed, err: %v", err))
	}
	return client, nil
}
