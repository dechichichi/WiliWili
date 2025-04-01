package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go"
)

var MinioClientGlobal *MinioClient

type MinioClient struct {
	Client *minio.Client
}

func InitMinioClient(endpoint, accessKeyID, secretAccessKey string) error {
	// 初始化 MinIO 客户端
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}
	// 设置全局客户端实例
	MinioClientGlobal = &MinioClient{Client: client}
	return nil
}

func (m *MinioClient) UploadFile(bucketName, objectName, Location, ContentType string, file []byte) error {
	reader := bytes.NewReader(file)
	exist, _ := m.Client.BucketExists(bucketName)
	if !exist {
		err := m.Client.MakeBucket(bucketName, Location)
		if err != nil {
			return  fmt.Errorf("Failed to create bucket %s: %v", bucketName, err)
		}
	}
	// 调用 PutObject 方法
	options := minio.PutObjectOptions{ContentType: ContentType}
	n, err := m.Client.PutObject(bucketName, objectName, reader, -1, options)
	if err != nil {
		log.Printf("Failed to upload %s: %v", objectName, err)
		return  fmt.Errorf("Failed to upload %s: %v", objectName, err)
	}
	log.Printf("Successfully uploaded %s of size %d", objectName, n)
	return  nil
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
