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
	StoreImage(ctx context.Context, uid int64, image []byte) (*model.Image, error)
}

type UserCache interface {
	NewUserId(ctx context.Context) (int64, error)
}
