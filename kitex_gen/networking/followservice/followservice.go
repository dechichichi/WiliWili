// Code generated by Kitex v0.12.1. DO NOT EDIT.

package followservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	networking "wiliwili/kitex_gen/networking"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"follow": kitex.NewMethodInfo(
		followHandler,
		newFollowServiceFollowArgs,
		newFollowServiceFollowResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"getFollowList": kitex.NewMethodInfo(
		getFollowListHandler,
		newFollowServiceGetFollowListArgs,
		newFollowServiceGetFollowListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"getFollowerList": kitex.NewMethodInfo(
		getFollowerListHandler,
		newFollowServiceGetFollowerListArgs,
		newFollowServiceGetFollowerListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"getFriendList": kitex.NewMethodInfo(
		getFriendListHandler,
		newFollowServiceGetFriendListArgs,
		newFollowServiceGetFriendListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	followServiceServiceInfo                = NewServiceInfo()
	followServiceServiceInfoForClient       = NewServiceInfoForClient()
	followServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return followServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return followServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return followServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FollowService"
	handlerType := (*networking.FollowService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "networking",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func followHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*networking.FollowServiceFollowArgs)
	realResult := result.(*networking.FollowServiceFollowResult)
	success, err := handler.(networking.FollowService).Follow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceFollowArgs() interface{} {
	return networking.NewFollowServiceFollowArgs()
}

func newFollowServiceFollowResult() interface{} {
	return networking.NewFollowServiceFollowResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*networking.FollowServiceGetFollowListArgs)
	realResult := result.(*networking.FollowServiceGetFollowListResult)
	success, err := handler.(networking.FollowService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceGetFollowListArgs() interface{} {
	return networking.NewFollowServiceGetFollowListArgs()
}

func newFollowServiceGetFollowListResult() interface{} {
	return networking.NewFollowServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*networking.FollowServiceGetFollowerListArgs)
	realResult := result.(*networking.FollowServiceGetFollowerListResult)
	success, err := handler.(networking.FollowService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceGetFollowerListArgs() interface{} {
	return networking.NewFollowServiceGetFollowerListArgs()
}

func newFollowServiceGetFollowerListResult() interface{} {
	return networking.NewFollowServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*networking.FollowServiceGetFriendListArgs)
	realResult := result.(*networking.FollowServiceGetFriendListResult)
	success, err := handler.(networking.FollowService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceGetFriendListArgs() interface{} {
	return networking.NewFollowServiceGetFriendListArgs()
}

func newFollowServiceGetFriendListResult() interface{} {
	return networking.NewFollowServiceGetFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Follow(ctx context.Context, req *networking.FollowOperationReq) (r *networking.FollowOperationResp, err error) {
	var _args networking.FollowServiceFollowArgs
	_args.Req = req
	var _result networking.FollowServiceFollowResult
	if err = p.c.Call(ctx, "follow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *networking.FollowListReq) (r *networking.FollowListResp, err error) {
	var _args networking.FollowServiceGetFollowListArgs
	_args.Req = req
	var _result networking.FollowServiceGetFollowListResult
	if err = p.c.Call(ctx, "getFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *networking.FollowerListReq) (r *networking.FollowerListResp, err error) {
	var _args networking.FollowServiceGetFollowerListArgs
	_args.Req = req
	var _result networking.FollowServiceGetFollowerListResult
	if err = p.c.Call(ctx, "getFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *networking.FriendListReq) (r *networking.FriendListResp, err error) {
	var _args networking.FollowServiceGetFriendListArgs
	_args.Req = req
	var _result networking.FollowServiceGetFriendListResult
	if err = p.c.Call(ctx, "getFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
