package usecase_test

import (
	"testing"
	"wiliwili/app/user/domain/model"
	"wiliwili/app/user/domain/service"
	"wiliwili/app/user/usecase"
	"wiliwili/app/user/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRegister(t *testing.T) {
	mockdb := new(mocks.UserDB)
	mockService := service.NewUserService(mockdb, nil)
	uc := usecase.NewUserUseCase(mockdb, mockService, nil)
	user := model.User{
		Username: "testuser",
		Password: "password",
		Email:    "test@example.com",
	}
	mockdb.On("IsUserExist", mock.Anything, user.Username).Return(false, nil)
	mockdb.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	uid, err := uc.UserRegister(nil, &user)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), uid)
	mockdb.AssertExpectations(t)
}
