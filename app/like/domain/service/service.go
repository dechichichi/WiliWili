package service

import (
	"context"
	logincontext "wiliwili/pkg/base/context"
	"wiliwili/pkg/errno"
)

func (svc *LikeService) IndentifyUser(ctx context.Context, uid int64) error {
	id, err := logincontext.GetLoginData(ctx)
	if err != nil {
		return err
	}
	if uid != id {
		return errno.Errorf(errno.ErrUserNotHavePermission, "用户没有权限,id:%d,uid:%d", id, uid)
	}
	return nil
}
