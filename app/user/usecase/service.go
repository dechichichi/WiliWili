package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/user/domain/model"
	"wiliwili/pkg/constants"
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
	uid := ctx.Value(constants.LoginDataKey)
	theuser, err := uc.db.GEtUserById(ctx, uid.(int64))
	if theuser == nil {
		return nil, fmt.Errorf("user not exist")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	// 检测图片格式
	imageFormat, err := uc.svc.DetectImageFormat(avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to detect image format: %w", err)
	}

	// 检查是否支持该格式
	supportedFormats := []string{"JPEG", "PNG"}
	if !contains(supportedFormats, imageFormat) {
		return nil, fmt.Errorf("unsupported image format: %s", imageFormat)
	}

	// 检查图片大小
	if len(avatar) > 10*1024*1024 { // 限制为 10MB
		return nil, fmt.Errorf("image size exceeds limit")
	}
	image := &model.Image{}
	image.Imageid = uc.svc.NewId()
	image.Url, err = uc.svc.DownloadAvatar(avatar, string(image.Imageid), imageFormat)
	if err != nil {
		return nil, fmt.Errorf("failed to download avatar: %w", err)
	}
	//上传图片
	err = uc.db.StoreImage(ctx, uid.(int64), image)
	if err != nil {
		return nil, fmt.Errorf("failed to store image: %w", err)
	}
	return image, nil
}

// 辅助函数：检查切片中是否包含某个值
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
