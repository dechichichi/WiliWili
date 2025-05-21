package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"wiliwili/app/video/domain/model"
	"wiliwili/config"
	"wiliwili/pkg/constants"

	"github.com/IBM/sarama"
)

func (uc *useCase) SubmitVideo(ctx context.Context, video *model.Video, data []byte) (string, string, error) {
	uid, err := uc.svc.CheckUser(ctx)
	if err != nil {
		return "", "", err
	}
	video.OwnerID = fmt.Sprintf("%d", uid)
	videoid := uc.svc.NewId()

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{config.Kafka.Broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// 将上传任务发送到 Kafka 主题
	task := map[string]interface{}{
		"videoid": videoid,
		"data":    data,
	}
	taskBytes, err := json.Marshal(task)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal upload task: %v", err)
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: config.Kafka.Topic,
		Value: sarama.StringEncoder(taskBytes),
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to produce message: %v", err)
	}
	url := fmt.Sprintf("%s/%s/%d", config.Minio.Addr, constants.VideoBucket, videoid)
	video.VideoID = videoid
	video.VideoURL = url
	err = uc.db.StoreVideo(ctx, video)
	if err != nil {
		return "", "", err
	}
	return videoid, url, nil
}
func (uc *useCase) GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error) {
	model, err := uc.cache.GetVideo(ctx, videoid)
	if err != nil {
		model, err = uc.db.GetVideo(ctx, videoid)
		if err != nil {
			return nil, err
		}
		// 尝试将数据存入缓存，但即使失败也不影响返回结果
		if err := uc.cache.SetVideo(ctx, model); err != nil {
			log.Printf("Failed to set video in cache: %v", err)
		}
	}
	return model, nil
}
func (uc *useCase) SearchVideo(ctx context.Context, keyword string, pageNum, pageSize int64) ([]*model.VideoProfile, error) {
	models, err := uc.db.SearchVideo(ctx, keyword, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return models, nil
}
func (uc *useCase) VideoTrending(ctx context.Context, pageNum, pageSize int64) ([]*model.VideoProfile, error) {
	models, err := uc.db.VideoTrending(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return models, nil
}
