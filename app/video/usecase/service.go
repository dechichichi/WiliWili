package usecase

import (
	"context"
	"fmt"
	"wiliwili/app/video/domain/model"
	"wiliwili/config"
	"wiliwili/pkg/constants"
)

func (uc *useCase) SubmitVideo(ctx context.Context, video *model.Video, data []byte) (string, string, error) {
	uid, err := uc.svc.CheckUser(ctx)
	if err != nil {
		return "", "", err
	}
	video.OwnerID = fmt.Sprintf("%d", uid)
	videoid := uc.svc.NewId()
	err = uc.svc.UploadloadAvatar(data, fmt.Sprintf("%d", videoid))
	if err != nil {
		return "", "", err
	}
	url := fmt.Sprintf("%s/%s/%s", config.Minio.Addr, constants.VideoBucket, videoid)
	video.VideoID = videoid
	video.VideoURL = url
	err = uc.db.StoreVideo(ctx, video)
	if err != nil {
		return "", "", err
	}
	return videoid, url, nil
}
func (uc *useCase) GetVideo(ctx context.Context, videoid string) (*model.VideoProfile, error) {
	model, err := uc.db.GetVideo(ctx, videoid)
	if err != nil {
		return nil, err
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
