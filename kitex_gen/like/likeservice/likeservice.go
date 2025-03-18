// Code generated by Kitex v0.12.1. DO NOT EDIT.

package likeservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	like "wiliwili/kitex_gen/like"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"LikeComment": kitex.NewMethodInfo(
		likeCommentHandler,
		newLikeServiceLikeCommentArgs,
		newLikeServiceLikeCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"LikeVideo": kitex.NewMethodInfo(
		likeVideoHandler,
		newLikeServiceLikeVideoArgs,
		newLikeServiceLikeVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CommentLikeNum": kitex.NewMethodInfo(
		commentLikeNumHandler,
		newLikeServiceCommentLikeNumArgs,
		newLikeServiceCommentLikeNumResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoLikeNum": kitex.NewMethodInfo(
		videoLikeNumHandler,
		newLikeServiceVideoLikeNumArgs,
		newLikeServiceVideoLikeNumResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	likeServiceServiceInfo                = NewServiceInfo()
	likeServiceServiceInfoForClient       = NewServiceInfoForClient()
	likeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return likeServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return likeServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return likeServiceServiceInfoForClient
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
	serviceName := "LikeService"
	handlerType := (*like.LikeService)(nil)
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
		"PackageName": "like",
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

func likeCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceLikeCommentArgs)
	realResult := result.(*like.LikeServiceLikeCommentResult)
	success, err := handler.(like.LikeService).LikeComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceLikeCommentArgs() interface{} {
	return like.NewLikeServiceLikeCommentArgs()
}

func newLikeServiceLikeCommentResult() interface{} {
	return like.NewLikeServiceLikeCommentResult()
}

func likeVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceLikeVideoArgs)
	realResult := result.(*like.LikeServiceLikeVideoResult)
	success, err := handler.(like.LikeService).LikeVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceLikeVideoArgs() interface{} {
	return like.NewLikeServiceLikeVideoArgs()
}

func newLikeServiceLikeVideoResult() interface{} {
	return like.NewLikeServiceLikeVideoResult()
}

func commentLikeNumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceCommentLikeNumArgs)
	realResult := result.(*like.LikeServiceCommentLikeNumResult)
	success, err := handler.(like.LikeService).CommentLikeNum(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceCommentLikeNumArgs() interface{} {
	return like.NewLikeServiceCommentLikeNumArgs()
}

func newLikeServiceCommentLikeNumResult() interface{} {
	return like.NewLikeServiceCommentLikeNumResult()
}

func videoLikeNumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*like.LikeServiceVideoLikeNumArgs)
	realResult := result.(*like.LikeServiceVideoLikeNumResult)
	success, err := handler.(like.LikeService).VideoLikeNum(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLikeServiceVideoLikeNumArgs() interface{} {
	return like.NewLikeServiceVideoLikeNumArgs()
}

func newLikeServiceVideoLikeNumResult() interface{} {
	return like.NewLikeServiceVideoLikeNumResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) LikeComment(ctx context.Context, req *like.LikeCommentReq) (r *like.LikeCommentResp, err error) {
	var _args like.LikeServiceLikeCommentArgs
	_args.Req = req
	var _result like.LikeServiceLikeCommentResult
	if err = p.c.Call(ctx, "LikeComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LikeVideo(ctx context.Context, req *like.LikeVideoReq) (r *like.LikeVideoResp, err error) {
	var _args like.LikeServiceLikeVideoArgs
	_args.Req = req
	var _result like.LikeServiceLikeVideoResult
	if err = p.c.Call(ctx, "LikeVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentLikeNum(ctx context.Context, req *like.CommentLikeNumReq) (r *like.CommentLikeNumResp, err error) {
	var _args like.LikeServiceCommentLikeNumArgs
	_args.Req = req
	var _result like.LikeServiceCommentLikeNumResult
	if err = p.c.Call(ctx, "CommentLikeNum", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoLikeNum(ctx context.Context, req *like.VideoLikeNumReq) (r *like.VideoLikeNumResp, err error) {
	var _args like.LikeServiceVideoLikeNumArgs
	_args.Req = req
	var _result like.LikeServiceVideoLikeNumResult
	if err = p.c.Call(ctx, "VideoLikeNum", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
