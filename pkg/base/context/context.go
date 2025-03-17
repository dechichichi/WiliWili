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

func streamFromContext(ctx context.Context, key string) (string, bool) {
	md, success := metadata.FromIncomingContext(ctx)
	return md.Get(key)[0], success
}

func streamAppendContext(ctx context.Context, key string, value string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, key, value)
}
