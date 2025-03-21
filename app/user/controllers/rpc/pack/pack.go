package pack

import (
	"wiliwili/kitex_gen/model"

	domainmodel "wiliwili/app/user/domain/model"
)

func BuildUser(user *domainmodel.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Username: user.Username,
		Uid:      user.Id,
	}
}

func BuildUserProfile(user *domainmodel.UserProfile) *model.UserProfile {
	return &model.UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Signature: user.Signature,
	}
}

func BuildImage(image *domainmodel.Image) *model.Image {
	return &model.Image{
		ImageId:  image.ImageID,
		ImageUrl: image.Url,
	}
}
