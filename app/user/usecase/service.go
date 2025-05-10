package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"wiliwili/app/user/domain/model"
	"wiliwili/config"
	"wiliwili/pkg/constants"

	"github.com/IBM/sarama"
)

func (uc *useCase) UserRegister(ctx context.Context, user *model.User) (int64, error) {
	err := uc.db.IsUserExist(ctx, user.Username)
	if err != nil {
		return 0, err
	}
	user.Uid = uc.svc.NewId()
	err = uc.db.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return user.Uid, nil
}

func (uc *useCase) UserLogin(ctx context.Context, user *model.User) (*model.UserInfo, error) {
	tuser, err := uc.db.GEtUserById(ctx, user.Uid)
	if err != nil {
		return nil, err
	}
	userInfo, err := uc.svc.CheckPassword(ctx, tuser, user.Password)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (uc *useCase) UserProfile(ctx context.Context, uid int64) (*model.UserProfile, error) {
	//先在缓存中找
	user, err := uc.cache.GetUser(ctx, uid)
	// 如果缓存中没有，则从数据库中获取
	if err != nil {
		user, err = uc.db.GEtUserById(ctx, uid)
		if err != nil {
			return nil, err
		}
		err = uc.cache.StoreUser(ctx, uid, user)
		if err != nil {
			return nil, err
		}
	}
	userProfile := &model.UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Gender:    user.Gender,
		Signature: user.Signature,
	}
	return userProfile, nil
}

func (uc *useCase) UserAvatarUpload(ctx context.Context, uid int64, avatar []byte) (*model.Image, error) {
	// 检查用户是否有权限
	err := uc.svc.IndentifyUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{config.Kafka.Broker}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()
	// 生成唯一的图片 ID
	imageid := uc.svc.NewId()
	// 将上传任务发送到 Kafka 主题
	task := map[string]interface{}{
		"imageid": fmt.Sprintf("%d", imageid),
		"avatar":  avatar,
	}
	taskBytes, err := json.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal upload task: %v", err)
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: config.Kafka.Topic,
		Value: sarama.StringEncoder(taskBytes),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to produce message: %v", err)
	}
	// 返回图片信息（此时图片尚未上传）
	url := fmt.Sprintf("%s/%s/%d", config.Minio.Addr, constants.ImageBucket, imageid)
	image := &model.Image{
		Uid:     uid,
		ImageID: imageid,
		Url:     url,
	}
	return image, nil
}

func (uc *useCase) UserAvatarGet(ctx context.Context, uid int64) (string, error) {
	//先在缓存中找
	url, err := uc.cache.GetImage(ctx, uid)
	if err == nil {
		return url, nil
	}
	//再在数据库中找
	image, err := uc.db.GetImage(ctx, uid)
	if err != nil {
		return "", err
	}
	//更新缓存
	err = uc.cache.StoreImage(ctx, image)
	if err != nil {
		return "", err
	}
	return image.Url, nil
}
