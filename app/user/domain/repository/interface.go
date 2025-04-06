package repository

import (
	"context"
	"wiliwili/app/user/domain/model"
)

type UserDB interface {
	IsUserExist(ctx context.Context, username string) error
	GEtUserById(ctx context.Context, uid int64) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	StoreImage(ctx context.Context, image *model.Image) error
	GetImage(ctx context.Context, imageid int64) (*model.Image, error)
}

type UserCache interface {
	GetUser(ctx context.Context, uid int64) (*model.User, error)
	StoreUser(ctx context.Context, uid int64, user *model.User) error
	GetImage(ctx context.Context, imageID int64) (string, error)
	StoreImage(ctx context.Context, image *model.Image) error
}
