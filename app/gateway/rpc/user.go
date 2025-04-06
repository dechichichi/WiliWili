package rpc

import (
	"context"
	"wiliwili/kitex_gen/user"
	"wiliwili/pkg/base/client"
)

func InitUserRPC() {
	c, err := client.InitUserRPC()
	if err != nil {
		panic(err)
	}
	userClient = *c
}

func RegisterUser(ctx context.Context, req *user.UserRegisterReq) (reponse *user.UserRegisterResp, err error) {
	resp, err := userClient.UserRegister(ctx, req)
	return resp, err
}

func Login(ctx context.Context, req *user.UserLoginReq) (reponse *user.UserLoginResp, err error) {
	resp, err := userClient.UserLogin(ctx, req)
	return resp, err
}

func GetProfile(ctx context.Context, req *user.UserProfileReq) (reponse *user.UserProfileResp, err error) {
	resp, err := userClient.UserProfile(ctx, req)
	return resp, err
}
func UploadAvatar(ctx context.Context, req *user.UserAvatarUploadReq) (reponse *user.UserAvatarUploadResp, err error) {
	resp, err := userClient.UserAvatarUpload(ctx, req)
	return resp, err
}
func GetAvatar(ctx context.Context, req *user.UserAvatarGetReq) (reponse *user.UserAvatarGetResp, err error) {
	resp, err := userClient.UserAvatarGet(ctx, req)
	return resp, err
}
