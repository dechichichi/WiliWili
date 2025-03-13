package constants

import "time"

const (
	ServerMaxRequestBodySize = 1 << 31

	CorsMaxAge = 12 * time.Hour

	SentinelThreshold        = 100
	SentinelStatIntervalInMs = 1000
	LoginDataKey             = "loginData"
)
