package main

import (
	"context"
	"log"
	"net/http"

	"wiliwili/app/gateway/router"
	"wiliwili/config"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func main() {
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)
	listenAddr := ":8080"
	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(constants.ServerMaxRequestBodySize),
	)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.Otel.Servicesname),
		provider.WithExportEndpoint(config.Otel.Addr),
		provider.WithInsecure(),
	)
	defer func() {
		if err := p.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown OpenTelemetry provider: %v", err)
		}
	}()
	router.GeneratedRegister(h)
	utils.InitSentinel()
	// 启动 HTTP 服务器
	go func() {
		log.Println("WebSocket 服务器启动，监听端口:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	h.Spin()
}
