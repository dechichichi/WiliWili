package base

import (
	"wiliwili/kitex_gen/model"
	"wiliwili/pkg/errno"
)

var SuccessBase = model.BaseResp{Code: errno.SuccessCode, Msg: errno.SuccessMsg}

func BuildBaseResp(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{Code: errno.SuccessCode, Msg: errno.Success.ErrorMsg}
	}
	Errno := errno.ConvertErr(err)
	return &model.BaseResp{Code: Errno.ErrorCode, Msg: Errno.ErrorMsg}
}

func buildSuccessBase() *model.BaseResp {
	return BuildBaseResp(nil)
}