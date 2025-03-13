package context

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"google.golang.org/grpc/metadata"
)

func newContext(ctx context.Context, key string, value string) context.Context {
	return metainfo.WithPersistentValue(ctx, key, value)
}

func fromContext(ctx context.Context, key string) (string, bool) {
	return metainfo.GetPersistentValue(ctx, key)
}

// 流式传输ctx不能用传统方式传递，详情见https://www.cloudwego.io/zh/docs/kitex/tutorials/advanced-feature/metainfo/#kitex-grpc-metadata
// 返回的第一个值只去第一位是因为目前只传了uid，后续需要修改的话还要加以修正
func streamFromContext(ctx context.Context, key string) (string, bool) {
	md, success := metadata.FromIncomingContext(ctx)
	return md.Get(key)[0], success
}

func streamAppendContext(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
