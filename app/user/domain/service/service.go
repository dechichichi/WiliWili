package service

import (
	"context"
	"fmt"
	"time"
	"wiliwili/app/user/domain/model"
	"wiliwili/pkg/constants"
	logincontext "wiliwili/pkg/base/context"
	"wiliwili/pkg/errno"
	"wiliwili/pkg/utils"

	"github.com/bwmarrin/snowflake"
)

func (svc *UserService) CheckPassword(ctx context.Context, user *model.User, password string) (*model.UserInfo, error) {
	if user.Password != password {
		return nil, errno.Errorf(errno.ErrCodePasswordIncorrect, "password incorrect")
	}
	userInfo := &model.UserInfo{
		Username: user.Username,
		Id:       user.Uid,
	}
	return userInfo, nil
}

// NewId 生成新id
func (svc *UserService) NewId() int64 {
	node, _ := snowflake.NewNode(1)
	return node.Generate().Int64()
}

func (svc *UserService) UploadloadAvatar(avatar []byte, imageID string) (string, error) {
	if err := utils.MinioClientGlobal.UploadFile(constants.ImageBucket, imageID, avatar); err != nil {
		return "", err
	}
	url, err := utils.MinioClientGlobal.Client.PresignedGetObject(constants.ImageBucket, imageID, 24*time.Hour, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (svc *UserService) IndentifyUser(ctx context.Context, uid int64) (error) {
	id, err := logincontext.GetLoginData(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get info in context when get loginData")
	}
	if uid != id {
		return errno.Errorf(errno.ErrUserNotHavePermission, "用户没有权限,id:%d,uid:%d", id, uid)
	}
	return nil
}
