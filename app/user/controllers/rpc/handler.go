package rpc

import (
	"context"
	"wiliwili/app/user/controllers/rpc/pack"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/usecase"
	"wiliwili/kitex_gen/user"
	"go.opentelemetry.io/otel/attribute"
	"wiliwili/app/user/controllers/rpc/otelmetrics"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func (u *UserHandler) UserRegister(ctx context.Context, req *user.UserRegisterReq) (r *user.UserRegisterResp, err error) {
	otelmetrics.UserQPSCounter.Add(ctx, 1, attribute.String("method", "UserRegister"))
	r = new(user.UserRegisterResp)
	user := &model.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Gender:    req.Gender,
		Signature: req.Signature,
	}
	var uid int64
	if uid, err = u.useCase.UserRegister(ctx, user); err != nil {
		return
	}
	r.Uid = uid
	return
}

func (u *UserHandler) UserLogin(ctx context.Context, req *user.UserLoginReq) (r *user.UserLoginResp, err error) {
	otelmetrics.UserQPSCounter.Add(ctx, 1, attribute.String("method", "UserLogin"))
	r = new(user.UserLoginResp)
	user := &model.User{
		Uid:      req.Uid,
		Password: req.Password,
	}
	userInfo, err := u.useCase.UserLogin(ctx, user)
	if err != nil {
		return
	}
	r.UserInfo = pack.BuildUser(userInfo)
	return
}

func (u *UserHandler) UserProfile(ctx context.Context, req *user.UserProfileReq) (r *user.UserProfileResp, err error) {
	otelmetrics.UserQPSCounter.Add(ctx, 1, attribute.String("method", "UserProfile"))
	r = new(user.UserProfileResp)
	UserProfile, err := u.useCase.UserProfile(ctx, req.Uid)
	if err != nil {
		return
	}
	r.UserProfile = pack.BuildUserProfile(UserProfile)
	return
}

func (u *UserHandler) UserAvatarUpload(ctx context.Context, req *user.UserAvatarUploadReq) (r *user.UserAvatarUploadResp, err error) {
	otelmetrics.UserQPSCounter.Add(ctx, 1, attribute.String("method", "UserAvatarUpload"))
	r = new(user.UserAvatarUploadResp)
	image, err := u.useCase.UserAvatarUpload(ctx, req.Uid, req.Avatar)
	if err != nil {
		return
	}
	if image == nil {
		return
	}
	r.Image = pack.BuildImage(image)
	return
}

func (u *UserHandler) UserAvatarGet(ctx context.Context, req *user.UserAvatarGetReq) (r *user.UserAvatarGetResp, err error) {
	otelmetrics.UserQPSCounter.Add(ctx, 1, attribute.String("method", "UserAvatarGet"))
	r = new(user.UserAvatarGetResp)
	url, err := u.useCase.UserAvatarGet(ctx, req.Uid)
	if err != nil {
		return
	}
	if url == "" {
		return
	}
	r.Url = url
	return
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase}
}
