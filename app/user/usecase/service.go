package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/user/domain/model"
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
	//err := uc.svc.IndentifyUser(ctx, uid)
	//if err != nil {
	//	return nil, err
	//}
	var url string
	url, err := uc.svc.UploadloadAvatar(avatar, fmt.Sprintf("%d", uid))
	if err != nil {
		return nil, fmt.Errorf("上传图片失败:%v", err)
	}
	image := &model.Image{
		Imageid: uc.svc.NewId(),
		Url:     url,
	}
	//上传图片
	err = uc.db.StoreImage(ctx, uid, image)
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, fmt.Errorf("上传图片失败:%v", err)
	}
	return image, nil
}
