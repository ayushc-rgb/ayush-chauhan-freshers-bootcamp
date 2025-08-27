package service

import "gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	ShowAll() []entity.Video
}
type videoService struct {
	videos []entity.Video
}

func (v *videoService) Save(newVideo entity.Video) entity.Video {
	// add these video into the slices of videos
	v.videos = append(v.videos, newVideo)
	return newVideo
}

// this function will show all the videos
func (v *videoService) ShowAll() []entity.Video {
	return v.videos
}
func New() VideoService {
	return &videoService{}
}
