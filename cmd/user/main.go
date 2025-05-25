package main

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"wiliwili/app/user"
	"wiliwili/config"
	"wiliwili/kitex_gen/user/userservice"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/middleware"
	"wiliwili/pkg/utils"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/export/otlp/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var serviceName = constants.UserServiceName

func init() {
	config.Init(serviceName)
}

func main() {
	// 日志输出 json 格式
	log.SetFlags(0)
	log.SetOutput(&jsonLogger{})

	// 初始化 OTel metrics
	ctx := context.Background()
	res, _ := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	client := otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithEndpoint(config.Otel.Addr),
		otlpmetricgrpc.WithInsecure(),
	)
	exp, _ := otlpmetricgrpc.New(ctx, client)
	mp := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(exp)),
	)
	global.SetMeterProvider(mp)
	defer func() {
		_ = mp.Shutdown(ctx)
	}()

	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("User: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("User: resolve tcp addr failed, err: %v", err)
	}
	err = utils.InitMinioClient(config.Minio.Addr, config.Minio.AccessKeyID, config.Minio.AccessKey)
	if err != nil {
		logger.Fatalf("User: new minio client failed, err: %v", err)
	}
	svr := userservice.NewServer(
		// 注入依赖
		user.InjectUserHandler(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
		server.WithMiddleware(middleware.Respond()),
	)
	go utils.InitAvatarUploadTask()
	if err = svr.Run(); err != nil {
		logger.Fatalf("User: run server failed, err: %v", err)
	}
}

// jsonLogger 实现 io.Writer，将 log.Print 输出为 json
// 可被 Loki/promtail 采集

type jsonLogger struct{}

func (j *jsonLogger) Write(p []byte) (n int, err error) {
	m := map[string]interface{}{
		"msg": string(p),
		"ts":  os.Getenv("TZ"),
	}
	b, _ := json.Marshal(m)
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
	return len(p), nil
}
