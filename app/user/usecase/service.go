package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/user/domain/model"
)

func (uc *useCase) UserRegister(ctx context.Context, user *model.User) (int64, error) {
	exist, err := uc.db.IsUserExist(ctx, user.Username)
	if err != nil {
		return 0, fmt.Errorf("failed to check user exist: %w", err)
	}
	if exist {
		return 0, fmt.Errorf("user already exist,%w", err)
	}
	uid, err := uc.cache.NewUserId(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get new user id: %w", err)
	}
	user.UId = uid
	err = uc.db.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return uid, nil
}

func (uc *useCase) UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error) {
	panic("implement me")
}

func (uc *useCase) UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	panic("implement me")
}

func (uc *useCase) UserAvatarUpload(ctx context.Context, avatar []byte) (*model.Image, error) {
	panic("implement me")
}
