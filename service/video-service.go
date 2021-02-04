package service

import "golang-gin-poc/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	findAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) findAll() []entity.Video {
	return service.videos
}

func New() VideoService {
	return &videoService{}
}
