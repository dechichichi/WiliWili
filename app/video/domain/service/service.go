package service

import (
	"context"
	logincontext "wiliwili/pkg/base/context"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/errno"
	"wiliwili/pkg/utils"

	"github.com/teris-io/shortid"
)

func (svc *VideoService) CheckUser(ctx context.Context) (int64, error) {
	id, err := logincontext.GetLoginData(ctx)
	if err != nil {
		return 0, errno.Errorf(errno.ErrUserNotHavePermission, "用户未登录")
	}
	return id, nil
}

func (svc *VideoService) UploadloadAvatar(data []byte, videoID string) error {
	err := utils.MinioClientGlobal.UploadFile(constants.VideoBucket, videoID, constants.Location, constants.VideoType, data)
	if err != nil {
		return err
	}
	return nil
}

func (svc *VideoService) NewId() string {
	id, _ := shortid.Generate()
	return id
}
