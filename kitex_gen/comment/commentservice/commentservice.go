// Code generated by Kitex v0.12.1. DO NOT EDIT.

package commentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	comment "wiliwili/kitex_gen/comment"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CommentVideo": kitex.NewMethodInfo(
		commentVideoHandler,
		newCommentServiceCommentVideoArgs,
		newCommentServiceCommentVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ReplyComment": kitex.NewMethodInfo(
		replyCommentHandler,
		newCommentServiceReplyCommentArgs,
		newCommentServiceReplyCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetCommentList": kitex.NewMethodInfo(
		getCommentListHandler,
		newCommentServiceGetCommentListArgs,
		newCommentServiceGetCommentListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteComment": kitex.NewMethodInfo(
		deleteCommentHandler,
		newCommentServiceDeleteCommentArgs,
		newCommentServiceDeleteCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	commentServiceServiceInfo                = NewServiceInfo()
	commentServiceServiceInfoForClient       = NewServiceInfoForClient()
	commentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForClient
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
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
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
		"PackageName": "comment",
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

func commentVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentVideoArgs)
	realResult := result.(*comment.CommentServiceCommentVideoResult)
	success, err := handler.(comment.CommentService).CommentVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentVideoArgs() interface{} {
	return comment.NewCommentServiceCommentVideoArgs()
}

func newCommentServiceCommentVideoResult() interface{} {
	return comment.NewCommentServiceCommentVideoResult()
}

func replyCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceReplyCommentArgs)
	realResult := result.(*comment.CommentServiceReplyCommentResult)
	success, err := handler.(comment.CommentService).ReplyComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceReplyCommentArgs() interface{} {
	return comment.NewCommentServiceReplyCommentArgs()
}

func newCommentServiceReplyCommentResult() interface{} {
	return comment.NewCommentServiceReplyCommentResult()
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceGetCommentListArgs)
	realResult := result.(*comment.CommentServiceGetCommentListResult)
	success, err := handler.(comment.CommentService).GetCommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentListArgs() interface{} {
	return comment.NewCommentServiceGetCommentListArgs()
}

func newCommentServiceGetCommentListResult() interface{} {
	return comment.NewCommentServiceGetCommentListResult()
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceDeleteCommentArgs)
	realResult := result.(*comment.CommentServiceDeleteCommentResult)
	success, err := handler.(comment.CommentService).DeleteComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceDeleteCommentArgs() interface{} {
	return comment.NewCommentServiceDeleteCommentArgs()
}

func newCommentServiceDeleteCommentResult() interface{} {
	return comment.NewCommentServiceDeleteCommentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentVideo(ctx context.Context, req *comment.CommentVideoReq) (r *comment.CommentVideoResp, err error) {
	var _args comment.CommentServiceCommentVideoArgs
	_args.Req = req
	var _result comment.CommentServiceCommentVideoResult
	if err = p.c.Call(ctx, "CommentVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ReplyComment(ctx context.Context, req *comment.ReplyCommentReq) (r *comment.ReplyCommentResp, err error) {
	var _args comment.CommentServiceReplyCommentArgs
	_args.Req = req
	var _result comment.CommentServiceReplyCommentResult
	if err = p.c.Call(ctx, "ReplyComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, req *comment.GetCommentListReq) (r *comment.GetCommentListResp, err error) {
	var _args comment.CommentServiceGetCommentListArgs
	_args.Req = req
	var _result comment.CommentServiceGetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, req *comment.DeleteCommentReq) (r *comment.DeleteCommentResp, err error) {
	var _args comment.CommentServiceDeleteCommentArgs
	_args.Req = req
	var _result comment.CommentServiceDeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
