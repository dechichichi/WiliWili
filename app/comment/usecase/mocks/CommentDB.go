// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "wiliwili/app/comment/domain/model"

	mock "github.com/stretchr/testify/mock"
)

// CommentDB is an autogenerated mock type for the CommentDB type
type CommentDB struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: ctx, comment
func (_m *CommentDB) CreateComment(ctx context.Context, comment *model.Comment) error {
	ret := _m.Called(ctx, comment)

	if len(ret) == 0 {
		panic("no return value specified for CreateComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Comment) error); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: ctx, commentID, uid
func (_m *CommentDB) DeleteComment(ctx context.Context, commentID string, uid string) error {
	ret := _m.Called(ctx, commentID, uid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, commentID, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommentList provides a mock function with given fields: ctx, videoID, page, pageSize, commenttype
func (_m *CommentDB) GetCommentList(ctx context.Context, videoID string, page int64, pageSize int64, commenttype int64) ([]*model.Comment, error) {
	ret := _m.Called(ctx, videoID, page, pageSize, commenttype)

	if len(ret) == 0 {
		panic("no return value specified for GetCommentList")
	}

	var r0 []*model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int64) ([]*model.Comment, error)); ok {
		return rf(ctx, videoID, page, pageSize, commenttype)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int64) []*model.Comment); ok {
		r0 = rf(ctx, videoID, page, pageSize, commenttype)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64, int64) error); ok {
		r1 = rf(ctx, videoID, page, pageSize, commenttype)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommentDB creates a new instance of CommentDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentDB {
	mock := &CommentDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
