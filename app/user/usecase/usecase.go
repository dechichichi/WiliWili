package usecase

import (
	"context"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/domain/repository"
	"wiliwili/app/user/domain/service"
)

type UserUsecase interface {
	UserRegister(ctx context.Context, user *model.User) (int64, error)
	UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error)
	UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error)
	UserAvatarUpload(ctx context.Context, uid int64, avatar []byte) (*model.Image, error)
}

type useCase struct {
	db    repository.UserDB
	svc   *service.UserService
	cache repository.UserCache
}

func NewUserUsecase(db repository.UserDB, svc *service.UserService, cache repository.UserCache) UserUsecase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}
