package mw

import (
	"context"
	"wiliwili/app/gateway/pack"
	metainfoContext "wiliwili/pkg/base/context"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader(constants.AuthHeader))
		_, uid, err := utils.CheckToken(token)
		if err != nil {
			pack.RespError(c, err)
			c.Abort()
			return
		}
		access, refresh, err := utils.CreateAllToken(uid)
		if err != nil {
			pack.RespError(c, err)
			c.Abort()
			return
		}
		ctx = metainfoContext.WithLoginData(ctx, uid)
		ctx = metainfoContext.SetStreamLoginData(ctx, uid)
		c.Header(constants.AccessTokenHeader, access)
		c.Header(constants.RefreshTokenHeader, refresh)
		c.Next(ctx)
	}
}
