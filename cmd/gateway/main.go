package main

import (
	"context"
	"log"

	"wiliwili/app/gateway/router"
	"wiliwili/config"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/utils"

	"wiliwili/app/gateway/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

var serviceName = constants.GateWayServiceName

func init() {
	config.Init(serviceName)
	rpc.Init()
}

func main() {
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)
	listenAddr := ":8080"
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		panic(err)
	}

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
	h.Spin()
}
