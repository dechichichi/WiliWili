package client

import (
	"context"
	"errors"
	"fmt"
	"wiliwili/config"

	"github.com/redis/go-redis/v9"
	"github.com/west2-online/DomTok/pkg/errno"
	"github.com/west2-online/DomTok/pkg/logger"
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
	l := logger.GetRedisLogger()
	redis.SetLogger(l)
	client.AddHook(l)
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("client.NewRedisClient: ping redis failed: %v", err))
	}
	return client, nil
}
