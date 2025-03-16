package context

import (
	"context"
	"fmt"
	"strconv"
	"wiliwili/pkg/constants"
)

// WithLoginData 将LoginData加入到context中，通过metainfo传递到RPC server
func WithLoginData(ctx context.Context, uid int64) context.Context {
	return newContext(ctx, constants.LoginDataKey, strconv.FormatInt(uid, 10))
}

// GetStreamLoginData 流式传输传递ctx, 获取loginData
func GetStreamLoginData(ctx context.Context) (int64, error) {
	uid, success := streamFromContext(ctx, constants.LoginDataKey)
	if !success {
		return -1, fmt.Errorf("Failed to get info in context when get loginData")
	}

	value, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("Failed to parse loginData to int64")
	}
	return value, nil
}

// SetStreamLoginData 流式传输传递ctx, 设置ctx值
func SetStreamLoginData(ctx context.Context, uid int64) context.Context {
	value := strconv.FormatInt(uid, 10)
	return streamAppendContext(ctx, constants.LoginDataKey, value)
}

func GetLoginData(ctx context.Context) (int64, error) {
	uid, success := fromContext(ctx, constants.LoginDataKey)
	if !success {
		return -1, fmt.Errorf("Failed to get info in context when get loginData")
	}

	value, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("Failed to parse loginData to int64")
	}
	return value, nil
}
