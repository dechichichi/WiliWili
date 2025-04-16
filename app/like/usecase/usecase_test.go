package usecase_test

import (
	"testing"
	"wiliwili/app/like/domain/model"
	"wiliwili/app/like/domain/service"
	"wiliwili/app/like/usecase"
	"wiliwili/app/like/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLikeSomeThing(t *testing.T) {
	mockdb := new(mocks.LikeDB)
	mockService := service.NewLikeService(mockdb)
	uc := usecase.NewLikeUseCase(mockdb, mockService)
	comment_like := model.CommentLike{
		CommentID: "123456",
		UserID:    "123456",
	}
	mockdb.On("LikeComment", mock.Anything, &comment_like, true, "123456").Return(nil)
	err := uc.LikeComment(nil, &comment_like, true)
	assert.NoError(t, err)
	video_like := model.VideoLike{
		VideoID: "123456",
		UserID:  "123456",
	}
	mockdb.On("LikeVideo", mock.Anything, &video_like, true, "123456").Return(nil)
	err = uc.LikeVideo(nil, &video_like, true)
	assert.NoError(t, err)
	mockdb.AssertExpectations(t)

}

func TestLikeNum(t *testing.T) {
	mockdb := new(mocks.LikeDB)
	mockService := service.NewLikeService(mockdb)
	uc := usecase.NewLikeUseCase(mockdb, mockService)
	comment_id := "123456"
	mockdb.On("CommentLikeNum", mock.Anything, comment_id).Return(int64(10), nil)
	count, err := uc.CommentLikeNum(nil, comment_id)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)
	video_id := "123456"
	mockdb.On("VideoLikeNum", mock.Anything, video_id).Return(int64(10), nil)
	count, err = uc.VideoLikeNum(nil, video_id)
	assert.NoError(t, err)
	mockdb.AssertExpectations(t)
}
