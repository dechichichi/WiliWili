package client

import (
	"errors"
	"fmt"
	"wiliwili/config"
	"wiliwili/kitex_gen/interactive/interactive"
	"wiliwili/kitex_gen/user/userservice"
	"wiliwili/kitex_gen/video/videoservice"
	"wiliwili/pkg/constants"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func initRPCClient[T any](serviceName string, newClientFunc func(string, ...client.Option) (T, error)) (*T, error) {
	if config.Etcd == nil || config.Etcd.Addr == "" {
		return nil, errors.New("config.Etcd.Addr is nil")
	}
	// 初始化Etcd Resolver
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		return nil, fmt.Errorf("initRPCClient etcd.NewEtcdResolver failed: %w", err)
	}
	// 初始化具体的RPC客户端
	client, err := newClientFunc(serviceName,
		client.WithResolver(r),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fmt.Sprintf(constants.KitexClientEndpointInfoFormat, serviceName)}),
	)
	if err != nil {
		return nil, fmt.Errorf("initRPCClient NewClient failed: %w", err)
	}
	return &client, nil
}

func InitUserRPC() (*userservice.Client, error) {
	return initRPCClient("user", userservice.NewClient)
}

func InitVideoRPC() (*videoservice.Client, error) {
	return initRPCClient("video", videoservice.NewClient)
}

func InitInteractiveRPC() (*interactive.Client, error) {
	return initRPCClient("interactive", interactive.NewClient)
}
