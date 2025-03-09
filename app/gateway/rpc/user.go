package rpc

import (
	"context"
	"wiliwili/kitex_gen/user"
)

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
