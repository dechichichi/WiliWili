package mysql

import (
	"context"
	"errors"
	"fmt"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/domain/repository"
	"wiliwili/pkg/errno"

	"gorm.io/gorm"
)

type userDB struct {
	client *gorm.DB
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) IsUserExist(ctx context.Context, username string) (error) {
	var user model.User
	if err := db.client.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errno.Errorf(errno.ErrCodeDBerror, "failed to get user by username: %w", err)
	}
	return errno.Errorf(errno.ErrCodeUserExist, "user already exist")
}

func (db *userDB) GEtUserById(ctx context.Context, uid int64) (*model.User, error) {
	var user model.User
	if err := db.client.Where("uid = ?", uid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.Errorf(errno.ErrCodeUserNotExist, "user not exist")
		}
		return nil, errno.Errorf(errno.ErrCodeDBerror, "failed to get user by id: %w", err)
	}
	return &user, nil
}

func (db *userDB) CreateUser(ctx context.Context, user *model.User) error {
	if err := db.client.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (db *userDB) StoreImage(ctx context.Context, uid int64, image *model.Image) error {
	err := db.client.Where("uid = ?", uid).Create(image).Error
	if err != nil {
		return fmt.Errorf("failed to store image: %w", err)
	}
	return nil
}
