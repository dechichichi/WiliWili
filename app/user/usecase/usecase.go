package usecase

import (
	"context"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/domain/respository"
	"wiliwili/app/user/domain/service"
)

type UserUsecase interface {
	UserRegister(ctx context.Context, user *model.User) (int64, error)
	UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error)
	UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error)
	UserAvatarUpload(ctx context.Context, avatar []byte) (*model.Image, error)
}

type useCase struct {
	db  respository.UserDB
	svc *service.UserService
}

func NewUserUsecase(db respository.UserDB, svc *service.UserService) UserUsecase {
	return &useCase{
		db:  db,
		svc: svc,
	}
}
