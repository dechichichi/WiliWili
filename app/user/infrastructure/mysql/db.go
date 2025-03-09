package mysql

import (
	"context"
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
	panic("implement me")
}
func (db *userDB) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	panic("implement me")
}
func (db *userDB) CheckPassword(ctx context.Context, username string, password string) (*model.UserInfo, error) {
	panic("implement me")
}
func (db *userDB) GetUserProFile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	panic("implement me")
}
