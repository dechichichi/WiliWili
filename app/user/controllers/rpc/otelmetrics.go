package otelmetrics

import (
	"context"
	"sync"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	UserQPSCounter metric.Int64Counter
	once sync.Once
)

func init() {
	once.Do(func() {
		meter := global.Meter("wiliwili-user-service")
		UserQPSCounter, _ = meter.Int64Counter(
			"user_handler_qps",
			metric.WithDescription("QPS of user handler methods"),
		)
	})
}

// 预留：如需自定义注册更多指标，可在此扩展 