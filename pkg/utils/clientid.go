package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateClientID 基于时间戳和随机数生成唯一 ID
func GenerateClientID(userID string) string {
	rand.Seed(time.Now().UnixNano())
	randomPart := rand.Intn(1000000) // 生成一个随机数
	return fmt.Sprintf("client_%s_%d_%d", userID, time.Now().UnixNano(), randomPart)
}
