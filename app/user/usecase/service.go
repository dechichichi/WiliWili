package usecase

import (
	"context"
	"wiliwili/app/user/domain/model"
)

func (uc *useCase) UserRegister(ctx context.Context, user *model.User) (int64, error) {
	panic("implement me")
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
