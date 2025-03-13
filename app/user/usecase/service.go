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
	user.Uid = uid
	err = uc.db.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return uid, nil
}

func (uc *useCase) UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error) {
	theuser, err := uc.db.GEtUserById(ctx, user.Uid)
	if user == nil {
		return nil, fmt.Errorf("user not exist")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	userInfo, err := uc.svc.CheckPassword(ctx, theuser, user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to check password: %w", err)
	}
	return userInfo, nil
}

func (uc *useCase) UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	userprofile, err := uc.cache.GetUserProFile(ctx, uid)
	if err == nil {
		userprofile, err = uc.db.GetUserProFile(ctx, uid)
		if err != nil {
			return nil, fmt.Errorf("failed to get user profile: %w", err)
		}
	}
	err = uc.cache.StoreUserProFile(ctx, uid, userprofile)
	if err != nil {
		return nil, fmt.Errorf("failed to store user profile: %w", err)
	}
	return userprofile, nil
}

func (uc *useCase) UserAvatarUpload(ctx context.Context, avatar []byte) (*model.Image, error) {
	//检查用户是否有权限

	panic("implement me")
}
