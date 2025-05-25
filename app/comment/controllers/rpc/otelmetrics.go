package otelmetrics

import (
	"sync"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	CommentQPSCounter metric.Int64Counter
	once sync.Once
)

func init() {
	once.Do(func() {
		meter := global.Meter("wiliwili-comment-service")
		CommentQPSCounter, _ = meter.Int64Counter(
			"comment_handler_qps",
			metric.WithDescription("QPS of comment handler methods"),
		)
	})
} 