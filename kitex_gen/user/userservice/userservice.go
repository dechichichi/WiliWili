// Code generated by Kitex v0.12.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "wiliwili/kitex_gen/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"userRegister": kitex.NewMethodInfo(
		userRegisterHandler,
		newUserServiceUserRegisterArgs,
		newUserServiceUserRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"userLogin": kitex.NewMethodInfo(
		userLoginHandler,
		newUserServiceUserLoginArgs,
		newUserServiceUserLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"userProfile": kitex.NewMethodInfo(
		userProfileHandler,
		newUserServiceUserProfileArgs,
		newUserServiceUserProfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"userAvatarUpload": kitex.NewMethodInfo(
		userAvatarUploadHandler,
		newUserServiceUserAvatarUploadArgs,
		newUserServiceUserAvatarUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
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

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserRegisterArgs)
	realResult := result.(*user.UserServiceUserRegisterResult)
	success, err := handler.(user.UserService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegisterArgs() interface{} {
	return user.NewUserServiceUserRegisterArgs()
}

func newUserServiceUserRegisterResult() interface{} {
	return user.NewUserServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserLoginArgs)
	realResult := result.(*user.UserServiceUserLoginResult)
	success, err := handler.(user.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return user.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return user.NewUserServiceUserLoginResult()
}

func userProfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserProfileArgs)
	realResult := result.(*user.UserServiceUserProfileResult)
	success, err := handler.(user.UserService).UserProfile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserProfileArgs() interface{} {
	return user.NewUserServiceUserProfileArgs()
}

func newUserServiceUserProfileResult() interface{} {
	return user.NewUserServiceUserProfileResult()
}

func userAvatarUploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserAvatarUploadArgs)
	realResult := result.(*user.UserServiceUserAvatarUploadResult)
	success, err := handler.(user.UserService).UserAvatarUpload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserAvatarUploadArgs() interface{} {
	return user.NewUserServiceUserAvatarUploadArgs()
}

func newUserServiceUserAvatarUploadResult() interface{} {
	return user.NewUserServiceUserAvatarUploadResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *user.UserRegisterReq) (r *user.UserRegisterResp, err error) {
	var _args user.UserServiceUserRegisterArgs
	_args.Req = req
	var _result user.UserServiceUserRegisterResult
	if err = p.c.Call(ctx, "userRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *user.UserLoginReq) (r *user.UserLoginResp, err error) {
	var _args user.UserServiceUserLoginArgs
	_args.Req = req
	var _result user.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "userLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserProfile(ctx context.Context, req *user.UserProfileReq) (r *user.UserProfileResp, err error) {
	var _args user.UserServiceUserProfileArgs
	_args.Req = req
	var _result user.UserServiceUserProfileResult
	if err = p.c.Call(ctx, "userProfile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserAvatarUpload(ctx context.Context, req *user.UserAvatarUploadReq) (r *user.UserAvatarUploadResp, err error) {
	var _args user.UserServiceUserAvatarUploadArgs
	_args.Req = req
	var _result user.UserServiceUserAvatarUploadResult
	if err = p.c.Call(ctx, "userAvatarUpload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
