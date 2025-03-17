// Code generated by Kitex v0.12.1. DO NOT EDIT.
package likeservice

import (
	server "github.com/cloudwego/kitex/server"
	like "wiliwili/kitex_gen/like"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler like.LikeService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler like.LikeService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
