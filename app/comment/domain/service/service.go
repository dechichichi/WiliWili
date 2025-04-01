package service

import (
	"context"
	logincontext "wiliwili/pkg/base/context"
	"wiliwili/pkg/errno"

	"github.com/teris-io/shortid"
)

func (svc *CommentService) CheckUser(ctx context.Context) (int64, error) {
	id, err := logincontext.GetLoginData(ctx)
	if err != nil {
		return 0, errno.Errorf(errno.ErrUserNotHavePermission, "用户未登录")
	}
	return id, nil
}
func (svc *CommentService) NewId() string {
	id, _ := shortid.Generate()
	return id
}
