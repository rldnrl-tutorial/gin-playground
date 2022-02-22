package service

import "gin-playground/src/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videService{}
}

func (service *videService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videService) FindAll() []entity.Video {
	return service.videos
}
