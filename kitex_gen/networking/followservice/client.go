// Code generated by Kitex v0.12.1. DO NOT EDIT.

package followservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	networking "wiliwili/kitex_gen/networking"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, req *networking.FollowOperationReq, callOptions ...callopt.Option) (r *networking.FollowOperationResp, err error)
	GetFollowList(ctx context.Context, req *networking.FollowListReq, callOptions ...callopt.Option) (r *networking.FollowListResp, err error)
	GetFollowerList(ctx context.Context, req *networking.FollowerListReq, callOptions ...callopt.Option) (r *networking.FollowerListResp, err error)
	GetFriendList(ctx context.Context, req *networking.FriendListReq, callOptions ...callopt.Option) (r *networking.FriendListResp, err error)
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
	return &kFollowServiceClient{
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

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) Follow(ctx context.Context, req *networking.FollowOperationReq, callOptions ...callopt.Option) (r *networking.FollowOperationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kFollowServiceClient) GetFollowList(ctx context.Context, req *networking.FollowListReq, callOptions ...callopt.Option) (r *networking.FollowListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kFollowServiceClient) GetFollowerList(ctx context.Context, req *networking.FollowerListReq, callOptions ...callopt.Option) (r *networking.FollowerListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kFollowServiceClient) GetFriendList(ctx context.Context, req *networking.FriendListReq, callOptions ...callopt.Option) (r *networking.FriendListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}
