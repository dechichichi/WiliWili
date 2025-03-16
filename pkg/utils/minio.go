package utils

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/minio/minio-go"
)

var MinioClientGlobal *MinioClient

type MinioClient struct {
	Client *minio.Client
}

func NewMinioClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinioClient, error) {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}
	client := &MinioClient{
		Client: minioClient,
	}
	MinioClientGlobal = client
	return client, nil
}

// UploadFile 上传文件
func (m *MinioClient) UploadFile(bucketName, objectName string, file []byte) error {
	// 将 []byte 转换为 io.Reader
	reader := bytes.NewReader(file)

	// 获取文件大小
	fileSize := int64(len(file))

	// 设置 PutObjectOptions
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream", // 根据需要设置 MIME 类型
	}

	// 调用 PutObject 方法
	_, err := m.Client.PutObject(bucketName, objectName, reader, fileSize, options)
	if err != nil {
		return err
	}
	return nil
}

// DownloadFile 下载文件
func (m *MinioClient) DownloadFile(bucketName, objectName, filePath string) error {
	// 创建本地文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 下载存储桶中的文件到本地
	err = m.Client.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Successfully downloaded", objectName)
	return nil
}

// DeleteFile 删除文件
func (m *MinioClient) DeleteFile(bucketName, objectName string) (bool, error) {
	// 删除存储桶中的文件
	err := m.Client.RemoveObject(bucketName, objectName)
	if err != nil {
		return false, err
	}
	return true, err
}

// ListObjects 列出文件
func (m *MinioClient) ListObjects(bucketName, prefix string) ([]string, error) {
	var objectNames []string

	for object := range m.Client.ListObjects(bucketName, prefix, true, nil) {
		if object.Err != nil {
			return nil, object.Err
		}

		objectNames = append(objectNames, object.Key)
	}

	return objectNames, nil
}

// GetPresignedGetObject 返回对象的url地址，有效期时间为expires
func (m *MinioClient) GetPresignedGetObject(bucketName string, objectName string, expires time.Duration) (string, error) {
	object, err := m.Client.PresignedGetObject(bucketName, objectName, expires, nil)
	if err != nil {
		return "", err
	}

	return object.String(), nil
}
