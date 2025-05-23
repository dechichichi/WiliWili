// Code generated by Kitex v0.12.1. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	video "wiliwili/kitex_gen/video"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	VideoSubmission(ctx context.Context, req *video.VideoSubmissionReq, callOptions ...callopt.Option) (r *video.VideoSubmissionResp, err error)
	VideoGet(ctx context.Context, req *video.VideoGetReq, callOptions ...callopt.Option) (r *video.VideoGetResp, err error)
	VideoSearch(ctx context.Context, req *video.VideoSearchReq, callOptions ...callopt.Option) (r *video.VideoSearchResp, err error)
	VideoTrending(ctx context.Context, req *video.VideoTrendingReq, callOptions ...callopt.Option) (r *video.VideoTrendingResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) VideoSubmission(ctx context.Context, req *video.VideoSubmissionReq, callOptions ...callopt.Option) (r *video.VideoSubmissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoSubmission(ctx, req)
}

func (p *kVideoServiceClient) VideoGet(ctx context.Context, req *video.VideoGetReq, callOptions ...callopt.Option) (r *video.VideoGetResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoGet(ctx, req)
}

func (p *kVideoServiceClient) VideoSearch(ctx context.Context, req *video.VideoSearchReq, callOptions ...callopt.Option) (r *video.VideoSearchResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoSearch(ctx, req)
}

func (p *kVideoServiceClient) VideoTrending(ctx context.Context, req *video.VideoTrendingReq, callOptions ...callopt.Option) (r *video.VideoTrendingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoTrending(ctx, req)
}
