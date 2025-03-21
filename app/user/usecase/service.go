package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/user/domain/model"
	"wiliwili/config"
	"wiliwili/pkg/constants"
)

func (uc *useCase) UserRegister(ctx context.Context, user *model.User) (int64, error) {
	err := uc.db.IsUserExist(ctx, user.Username)
	if err != nil {
		return 0, err
	}
	user.Uid = uc.svc.NewId()
	err = uc.db.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return user.Uid, nil
}

func (uc *useCase) UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error) {
	tuser, err := uc.db.GEtUserById(ctx, user.Uid)
	if err != nil {
		return nil, err
	}
	userInfo, err := uc.svc.CheckPassword(ctx, tuser, user.Password)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (uc *useCase) UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	user, err := uc.db.GEtUserById(ctx, uid)
	if err != nil {
		return nil, err
	}
	userProfile := &model.UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Signature: user.Signature,
	}
	return userProfile, nil
}

func (uc *useCase) UserAvatarUpload(ctx context.Context, uid int64, avatar []byte) (*model.Image, error) {
	//检查用户是否有权限
	err := uc.svc.IndentifyUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	var url string
	err = uc.svc.UploadloadAvatar(avatar, fmt.Sprintf("%d", uid))
	if err != nil {
		return nil, fmt.Errorf("upload avatar failed: %v", err)
	}
	url = fmt.Sprintf("%s/%s/%d", config.Minio.Addr,constants.ImageBucket, uid)
	image := &model.Image{
		Uid: uid,
		Url: url,
	}
	//上传图片
	err = uc.db.StoreImage(ctx, image)
	if err != nil {
		return nil, err
	}
	return image, nil
}
