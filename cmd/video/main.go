package main

import (
	"net"
	"wiliwili/app/video"
	"wiliwili/config"
	"wiliwili/kitex_gen/video/videoservice"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/middleware"
	"wiliwili/pkg/utils"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var serviceName = constants.VideoServiceName

func init() {
	config.Init(serviceName)
}

func main() {
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
	svr := videoservice.NewServer(
		video.InjectUserHandler(),
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
	go utils.InitVideoUploadTask()
	if err = svr.Run(); err != nil {
		logger.Fatalf("User: run server failed, err: %v", err)
	}
}
