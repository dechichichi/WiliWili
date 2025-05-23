// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "wiliwili/app/like/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// LikeUseCase is an autogenerated mock type for the LikeUseCase type
type LikeUseCase struct {
	mock.Mock
}

// CommentLikeNum provides a mock function with given fields: ctx, comment_id
func (_m *LikeUseCase) CommentLikeNum(ctx context.Context, comment_id string) (int64, error) {
	ret := _m.Called(ctx, comment_id)

	if len(ret) == 0 {
		panic("no return value specified for CommentLikeNum")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, comment_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, comment_id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, comment_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LikeComment provides a mock function with given fields: ctx, comment_like, islike
func (_m *LikeUseCase) LikeComment(ctx context.Context, comment_like *model.CommentLike, islike bool) error {
	ret := _m.Called(ctx, comment_like, islike)

	if len(ret) == 0 {
		panic("no return value specified for LikeComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CommentLike, bool) error); ok {
		r0 = rf(ctx, comment_like, islike)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LikeVideo provides a mock function with given fields: ctx, video_like, islike
func (_m *LikeUseCase) LikeVideo(ctx context.Context, video_like *model.VideoLike, islike bool) error {
	ret := _m.Called(ctx, video_like, islike)

	if len(ret) == 0 {
		panic("no return value specified for LikeVideo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.VideoLike, bool) error); ok {
		r0 = rf(ctx, video_like, islike)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VideoLikeNum provides a mock function with given fields: ctx, video_id
func (_m *LikeUseCase) VideoLikeNum(ctx context.Context, video_id string) (int64, error) {
	ret := _m.Called(ctx, video_id)

	if len(ret) == 0 {
		panic("no return value specified for VideoLikeNum")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, video_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, video_id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, video_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLikeUseCase creates a new instance of LikeUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLikeUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *LikeUseCase {
	mock := &LikeUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
