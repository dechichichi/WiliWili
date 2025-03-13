package service

import (
	"context"
	"fmt"
	"wiliwili/app/user/domain/model"
)

func (svc *UserService) CheckPassword(ctx context.Context, user *model.User, password string) (*model.UserInfo, error) {
	if user.Password != password {
		return nil, fmt.Errorf("password not correct")
	}
	userInfo := &model.UserInfo{
		Username: user.Username,
		Id:       user.Uid,
	}
	return userInfo, nil
}
