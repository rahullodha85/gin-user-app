package service

import (
	"github.com/rahullodha85/gin-user-app/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type VideoServiceObj struct{
	videos []entity.Video
}

func (service *VideoServiceObj) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *VideoServiceObj) FindAll() []entity.Video {
	if service.videos == nil {
		service.videos = []entity.Video{}
	}
	return service.videos
}
