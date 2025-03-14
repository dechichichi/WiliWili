package repository

import (
	"context"
	"wiliwili/app/user/domain/model"
)

type UserDB interface {
	IsUserExist(ctx context.Context, username string) (bool, error)
	GEtUserById(ctx context.Context, uid int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error)
	StoreImage(ctx context.Context, uid int64,image  *model.Image) error
}

type UserCache interface {
	NewUserId(ctx context.Context) (int64, error)
	GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error)
	StoreUserProFile(ctx context.Context, uid int64, profile *model.UserProfile) error
	GetImage(ctx context.Context, imageid int64) (*model.Image, error)
	StoreImage(ctx context.Context, imageid int64, image *model.Image) error
}
