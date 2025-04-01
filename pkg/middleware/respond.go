package middleware

import (
	"context"
	"errors"
	"wiliwili/kitex_gen/model"
	"wiliwili/pkg/errno"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
)

type response interface {
	GetResult() interface{}
	IsSetSuccess() bool
}

type baser interface {
	IsGetBase() bool
	GetBase() *model.BaseResp
	SetBase(base *model.BaseResp)
}

//拦截所有响应和操作
//为response中的BaseResponse加上code和msg

func Respond() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			err = next(ctx, req, resp)
			packResp, ok := resp.(response)
			if !ok || !packResp.IsSetSuccess() {
				return err
			}
			var res baser
			res, ok = packResp.GetResult().(baser)
			if !ok {
				server, _ := kitexutil.GetCaller(ctx)
				method, _ := kitexutil.GetMethod(ctx)
				logger.Errorf("response not implement baser, server: %s, method: %s", server, method)
				return err
			}
			if !res.IsGetBase() {
				res.SetBase(&model.BaseResp{})
			}
			base := res.GetBase()
			var e errno.ErrNo
			if err != nil&&base.Code==0 {
				base.Code=errno.SuccessCode
				base.Msg=errno.SuccessMsg
				return nil
			}else if errors.As(err, &e) {
				if base.Code == 0 {
					base.Code = e.ErrorCode
					base.Msg = e.ErrorMsg
				}
				return nil
			}
			return err
		}
	}
}
