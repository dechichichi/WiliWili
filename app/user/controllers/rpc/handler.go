package rpc

import (
	"context"
	"wiliwili/app/user/controllers/rpc/pack"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/usecase"
	"wiliwili/kitex_gen/user"
)

type UserHandler struct {
	useCase usecase.UserUsecase
}

func (u *UserHandler) UserRegister(ctx context.Context, req *user.UserRegisterReq) (r *user.UserRegisterResp, err error) {
	r = new(user.UserRegisterResp)
	user := &model.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Gender:    req.Gender,
		Signature: req.Signature,
	}
	uid, err := u.useCase.UserRegister(ctx, user)
	r.Uid = uid
	return
}

func (u *UserHandler) UserLogin(ctx context.Context, req *user.UserLoginReq) (r *user.UserLoginResp, err error) {
	r = new(user.UserLoginResp)
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
	}
	userInfo, err := u.useCase.UserLogin(ctx, user)
	r.UserInfo = pack.BuildUser(userInfo)
	return
}

func (u *UserHandler) UserProfile(ctx context.Context, req *user.UserProfileReq) (r *user.UserProfileResp, err error) {
	r = new(user.UserProfileResp)
	UserProfile, err := u.useCase.UserProfile(ctx, req.Uid)
	r.UserProfile = pack.BuildUserProfile(UserProfile)
	return

}

func (u *UserHandler) UserAvatarUpload(ctx context.Context, req *user.UserAvatarUploadReq) (r *user.UserAvatarUploadResp, err error) {
	r = new(user.UserAvatarUploadResp)
	image, err := u.useCase.UserAvatarUpload(ctx, req.Avatar)
	r.Image = pack.BuildImage(image)
	return
}

func NewUserHandler(useCase usecase.UserUsecase) *UserHandler {
	return &UserHandler{useCase}
}
