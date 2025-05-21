package mw

import (
	"context"
	"fmt"
	"wiliwili/app/gateway/pack"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/cloudwego/hertz/pkg/app"
)

func SentinelServer() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resourceName := fmt.Sprintf("%s:%s", c.Method(), c.URI().Path())

		entry, blockErr := api.Entry(resourceName, api.WithTrafficType(base.Inbound))
		if blockErr != nil {
			pack.RespError(c, blockErr)
			c.Abort()
			return
		}
		defer entry.Exit()

		c.Next(ctx)
	}
}