package mysql

import (
	"context"
	"errors"
	"fmt"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/domain/repository"

	"gorm.io/gorm"
)

type userDB struct {
	client *gorm.DB
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) IsUserExist(ctx context.Context, username string) (bool, error) {
	var user model.User
	if err := db.client.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check user exist: %w", err)
	}
	return true, fmt.Errorf("user %s already exist", username)
}

func (db *userDB) GEtUserById(ctx context.Context, uid int64) (*model.User, error) {
	var user model.User
	if err := db.client.Where("uid = ?", uid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (db *userDB) CreateUser(ctx context.Context, user *model.User) error {
	if err := db.client.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
func (db *userDB) GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	panic("implement me")
}

func (db *userDB) StoreImage(ctx context.Context, uid int64, image []byte) (*model.Image, error) {
	panic("implement me")
}
