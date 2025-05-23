package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"strings"
	"wiliwili/config"
	"wiliwili/pkg/constants"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

func GetMysqlDSN() (string, error) {
	if config.Mysql == nil {
		return "", fmt.Errorf("MySQL 配置未初始化")
	}
	dsn := strings.Join([]string{
		config.Mysql.Username, ":", config.Mysql.Password,
		"@tcp(", config.Mysql.Addr, ")/",
		config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true",
	}, "")

	return dsn, nil
}

// GetAvailablePort 会尝试获取可用的监听地址
func GetAvailablePort() (string, error) {
	if config.Service.AddrList == nil {
		return "", errors.New("服务地址列表为空")
	}
	for _, addr := range config.Service.AddrList {
		if ok := AddrCheck(addr); ok {
			return addr, nil
		}
	}
	return "", errors.New("没有可用的端口")
}

// AddrCheck 会检查当前的监听地址是否已被占用
func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	defer func() {
		if err := l.Close(); err != nil {
			logger.Errorf("utils.AddrCheck: failed to close listener: %v", err.Error())
		}
	}()
	return true
}

// CheckImageFileType 检查文件格式是否合规
func CheckImageFileType(header *multipart.FileHeader) (string, bool) {
	file, err := header.Open()
	if err != nil {
		return "", false
	}
	defer func() {
		// 捕获并处理关闭文件时可能发生的错误
		if err := file.Close(); err != nil {
			logger.Errorf("utils.CheckImageFileType: failed to close file: %v", err.Error())
		}
	}()

	buffer := make([]byte, constants.CheckFileTypeBufferSize)
	_, err = file.Read(buffer)
	if err != nil {
		return "", false
	}

	kind, _ := filetype.Match(buffer)

	// 检查是否为jpg、png
	switch kind {
	case types.Get("jpg"):
		return "jpg", true
	case types.Get("png"):
		return "png", true
	default:
		return "", false
	}
}

func CheckVideoFileType(header *multipart.FileHeader) (string, bool) {
	file, err := header.Open()
	if err != nil {
		return "", false
	}
	defer func() {
		// 捕获并处理关闭文件时可能发生的错误
		if err := file.Close(); err != nil {
			logger.Errorf("utils.CheckImageFileType: failed to close file: %v", err.Error())
		}
	}()

	buffer := make([]byte, constants.CheckFileTypeBufferSize)
	_, err = file.Read(buffer)
	if err != nil {
		return "", false
	}

	kind, _ := filetype.Match(buffer)

	// 检查是否为mp4
	switch kind {
	case types.Get("mp4"):
		return "mp4", true
	default:
		return "", false
	}
}

func FileToBytes(file *multipart.FileHeader) ([]byte, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}

	// 打开文件
	fileOpen, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileOpen.Close()

	// 使用 bytes.Buffer 读取文件内容
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, fileOpen)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
