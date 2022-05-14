package services

import (
	"golang-gin-poc/entities"
	"golang-gin-poc/repositories"
)

type VideoService interface {
	Create(entities.Video) entities.Video
	FindAll() []entities.Video
	Update(video entities.Video)
	Delete(video entities.Video)
}

type videoService struct {
	videoRepository repositories.VideoRepository
}

func NewVideoService(repo repositories.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Create(video entities.Video) entities.Video {
	service.videoRepository.Create(video)
	return video
}
func (service *videoService) FindAll() []entities.Video {
	return service.videoRepository.FindAll()
}
func (service *videoService) Update(video entities.Video) {
	service.videoRepository.Update(video)
}
func (service *videoService) Delete(video entities.Video) {
	service.videoRepository.Delete(video)
}
