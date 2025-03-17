package main

import (
	"wiliwili/app/gateway/router"
	"wiliwili/config"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/utils"

	"wiliwili/app/gateway/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

var serviceName = constants.GateWayServiceName

func init() {
	config.Init(serviceName)
	rpc.Init()
}

func main() {
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		panic(err)
	}
	utils.NewMinioClient(config.Minio.Addr, config.Minio.AccessKeyID, config.Minio.AccessKey, false)
	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(constants.ServerMaxRequestBodySize),
	)
	router.GeneratedRegister(h)
	h.Spin()
}
