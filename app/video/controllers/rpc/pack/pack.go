package pack

import (
	domainmodel "wiliwili/app/video/domain/model"
	"wiliwili/kitex_gen/model"
)

func ToVideoProfile(video *domainmodel.VideoProfile) *model.Video {
	return &model.Video{
		VideoId:       video.VideoID,
		VideoName:     video.VideoName,
		VideoDuration: video.VideoDuration,
		VideoUrl:      video.VideoURL,
		LikesCount:    video.LikesCount,
	}
}

func ToVideoProfileList(videos []*domainmodel.VideoProfile) []*model.Video {
	var result []*model.Video
	for _, video := range videos {
		result = append(result, ToVideoProfile(video))
	}
	return result
}