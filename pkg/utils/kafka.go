package utils

import (
	"encoding/json"
	"log"
	"wiliwili/config"
	"wiliwili/pkg/constants"

	"github.com/IBM/sarama"
)

type AvatarUploadTask struct {
	Uid    string `json:"imageid"`
	Avatar []byte `json:"avatar"`
}

type VideoUploadTask struct {
	Uid   string `json:"videoid"`
	Video []byte `json:"video"`
}

func InitAvatarUploadTask() {
	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{config.Kafka.Broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()
	// 获取 Kafka 主题的所有分区
	partitionList, err := consumer.Partitions(config.Kafka.Topic)
	if err != nil {
		log.Fatalf("Failed to get partition list for topic %s: %v", config.Kafka.Topic, err)
	}
	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(config.Kafka.Topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to consume partition %d: %v", partition, err)
		}
		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-pc.Messages():
					handleAvatarMessage(msg)
				case err := <-pc.Errors():
					log.Printf("Error while consuming message: %v", err)
				}
			}
		}(pc)
	}

	// 防止主程序退出
	select {}
}
func InitVideoUploadTask() {
	// 创建 Kafka 消费者
	consumer, err := sarama.NewConsumer([]string{config.Kafka.Broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()
	// 获取 Kafka 主题的所有分区
	partitionList, err := consumer.Partitions(config.Kafka.Topic)
	if err != nil {
		log.Fatalf("Failed to get partition list for topic %s: %v", config.Kafka.Topic, err)
	}
	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(config.Kafka.Topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Failed to consume partition %d: %v", partition, err)
		}
		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-pc.Messages():
					handleVideoMessage(msg)
				case err := <-pc.Errors():
					log.Printf("Error while consuming message: %v", err)
				}
			}
		}(pc)
	}

	// 防止主程序退出
	select {}
}

// handleMessage 处理 Kafka 消息
func handleAvatarMessage(msg *sarama.ConsumerMessage) {
	// 解析任务
	var task AvatarUploadTask
	err := json.Unmarshal(msg.Value, &task)
	if err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return
	}

	// 上传头像到 Minio
	err = MinioClientGlobal.UploadFile(constants.ImageBucket, task.Uid, constants.Location, constants.ImageType, task.Avatar)
	if err != nil {
		log.Printf("Failed to upload avatar: %v", err)
		return
	}

	log.Printf("Successfully uploaded avatar for user %d", task.Uid)
}

func handleVideoMessage(msg *sarama.ConsumerMessage) {
	// 解析任务
	var task VideoUploadTask
	err := json.Unmarshal(msg.Value, &task)
	if err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return
	}

	// 上传头像到 Minio
	err = MinioClientGlobal.UploadFile(constants.VideoBucket, task.Uid, constants.Location, constants.VideoType, task.Video)
	if err != nil {
		log.Printf("Failed to upload video: %v", err)
		return
	}

	log.Printf("Successfully uploaded video for user %s", task.Uid)
}
