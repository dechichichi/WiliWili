package otelmetrics

import (
	"sync"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	VideoQPSCounter metric.Int64Counter
	once sync.Once
)

func init() {
	once.Do(func() {
		meter := global.Meter("wiliwili-video-service")
		VideoQPSCounter, _ = meter.Int64Counter(
			"video_handler_qps",
			metric.WithDescription("QPS of video handler methods"),
		)
	})
} 