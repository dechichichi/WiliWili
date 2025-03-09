package repository

import (
	"context"
	"wiliwili/app/user/domain/model"
)

type UserDB interface {
	IsUserExist(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	CheckPassword(ctx context.Context, username string, password string) (*model.UserInfo, error)
	GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error)
}

type UserCache interface {
	
}