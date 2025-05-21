package usecase_test

import (
	"testing"
	"wiliwili/app/comment/domain/model"
	"wiliwili/app/comment/domain/service"
	"wiliwili/app/comment/usecase"
	"wiliwili/app/comment/usecase/mocks"
	"wiliwili/pkg/constants"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCommentSomthing(t *testing.T) {
	mockdb := new(mocks.CommentDB)
	mockService := service.NewCommentService(mockdb, nil)
	uc := usecase.NewcommentUseCase(mockdb, mockService, nil)
	video_comment := model.Comment{
		UserID:      "123456",
		CommentType: constants.CommentTypeVideo,
		BeCommentID: "123456",
	}
	mockdb.On("CreateComment", mock.Anything, &video_comment).Return(nil)
	comment_id, err := uc.CommentVideo(nil, &video_comment)
	assert.NoError(t, err)
	assert.Equal(t, "123456", comment_id)
	reply_comment := model.Comment{
		UserID:      "123456",
		CommentType: constants.CommentTypeReply,
		BeCommentID: "123456",
	}
	mockdb.On("CreateComment", mock.Anything, &reply_comment).Return(nil)
	comment_id, err = uc.ReplyComment(nil, &reply_comment)
	assert.NoError(t, err)
	assert.Equal(t, "123456", comment_id)
	mockdb.AssertExpectations(t)
}
