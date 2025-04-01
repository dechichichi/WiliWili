package service

import (
	"context"
	logincontext "wiliwili/pkg/base/context"
	"wiliwili/pkg/errno"
)

func (svc *LikeService) CheckUser(ctx context.Context) (int64, error) {
	id, err := logincontext.GetLoginData(ctx)
	if err != nil {
		return 0, errno.Errorf(errno.ErrUserNotHavePermission, "用户未登录")
	}
	return id, nil
}
