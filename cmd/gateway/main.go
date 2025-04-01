package main

import (
	"log"
	"net/http"

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

	// 启动 HTTP 服务器
	go func() {
		log.Println("WebSocket 服务器启动，监听端口:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	h.Spin()
}
