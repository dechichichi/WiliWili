package main

import (
	"wiliwili/app/gateway/router"
	"wiliwili/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	listenAddr := ":8080"
	h := server.New(
		server.WithHostPorts(listenAddr),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(constants.ServerMaxRequestBodySize),
	)
	router.GeneratedRegister(h)
	h.Spin()
}
