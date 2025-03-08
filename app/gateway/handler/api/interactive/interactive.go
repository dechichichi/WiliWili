// Code generated by hertz generator.

package interactive

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	interactive "wiliwili/app/gateway/model/api/interactive"
)

// LikePost .
// @router api/v1/interactive/likePost [POST]
func LikePost(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interactive.LikeOperationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interactive.LikeOperationResp)

	c.JSON(consts.StatusOK, resp)
}

// GetLikedPostList .
// @router api/v1/interactive/getLikedPostList [GET]
func GetLikedPostList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interactive.LikeListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interactive.LikeListResp)

	c.JSON(consts.StatusOK, resp)
}

// Comment .
// @router api/v1/interactive/comment [POST]
func Comment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interactive.CommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interactive.CommentResp)

	c.JSON(consts.StatusOK, resp)
}

// GetCommentList .
// @router api/v1/interactive/getCommentList [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interactive.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interactive.CommentListResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteComment .
// @router api/v1/interactive/deleteComment [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interactive.DeleteCommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interactive.DeleteCommentResp)

	c.JSON(consts.StatusOK, resp)
}
