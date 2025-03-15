// Code generated by hertz generator.

package user

import (
	"context"
	"wiliwili/app/gateway/pack"
	"wiliwili/app/gateway/rpc"

	api "wiliwili/app/gateway/model/api/user"
	"wiliwili/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterUser .
// @router api/v1/user/register [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegiterUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	resp, err := rpc.RegisterUser(ctx, &user.UserRegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Gender:   req.Gender,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp)
}

// Login .
// @router api/v1/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	resp, err := rpc.Login(ctx, &user.UserLoginReq{
		Uid:      req.ID,
		Password: req.Password,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp)
}

// GetProfile .

// @router api/v1/user/profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ProfileReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.GetProfile(ctx, &user.UserProfileReq{
		Uid: req.UserId,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp)
}

// UploadAvatar .
// @router api/v1/user/avatar [POST]
func UploadAvatar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserAvatarUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.UploadAvatar(ctx, &user.UserAvatarUploadReq{
		Avatar: req.Avatar,
	})
	if err != nil {
		pack.RespError(c, err)
		return
	}
	pack.RespData(c, resp)
}
