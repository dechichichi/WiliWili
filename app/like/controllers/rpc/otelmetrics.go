package otelmetrics

import (
	"sync"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	LikeQPSCounter metric.Int64Counter
	once sync.Once
)

func init() {
	once.Do(func() {
		meter := global.Meter("wiliwili-like-service")
		LikeQPSCounter, _ = meter.Int64Counter(
			"like_handler_qps",
			metric.WithDescription("QPS of like handler methods"),
		)
	})
} 