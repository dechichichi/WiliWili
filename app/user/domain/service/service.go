package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"wiliwili/app/user/domain/model"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/errno"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func (svc *UserService) CheckPassword(ctx context.Context, user *model.User, password string) (*model.UserInfo, error) {
	if user.Password != password {
		return nil, errno.Errorf(errno.ErrCodePasswordIncorrect, "password incorrect")
	}
	userInfo := &model.UserInfo{
		Username: user.Username,
		Id:       user.Uid,
	}
	return userInfo, nil
}

// NewId 生成新id
func (svc *UserService) NewId() int64 {
	node, _ := snowflake.NewNode(1)
	return node.Generate().Int64()
}

// DetectImageFormat 检测图片格式
func (svc *UserService) DetectImageFormat(data []byte) (string, error) {
	if len(data) < 4 {
		return "", fmt.Errorf("data too short to detect format")
	}

	switch {
	case data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return "JPEG", nil
	case data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47:
		return "PNG", nil
	case data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x38:
		return "GIF", nil
	case data[0] == 0x42 && data[1] == 0x4D:
		return "BMP", nil
	case data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46:
		return "WebP", nil
	default:
		return "", fmt.Errorf("unknown image format")
	}
}

func (svc *UserService) DownloadAvatar(avatar []byte, imageID string, imageFormat string) (string, error) {
	// 确保存储目录存在
	if err := os.MkdirAll(constants.ImageDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	var fileExtension string
	switch strings.ToUpper(imageFormat) {
	case "JPEG":
		fileExtension = ".jpg"
	case "PNG":
		fileExtension = ".png"
	default:
		return "", fmt.Errorf("unsupported image format: %s", imageFormat)
	}
	// 生成唯一的文件名
	fileName := uuid.New().String() + fileExtension
	filePath := filepath.Join(constants.ImageDir, fileName)
	// 将图片数据写入文件
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, bytes.NewReader(avatar)); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	// 生成访问 URL
	imageURL := fmt.Sprintf("%s/%s", constants.ImageBaseURL, fileName)
	return imageURL, nil
}
