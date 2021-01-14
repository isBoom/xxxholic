package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type ListVideoServics struct {
	Title string `json:"title" form:"title"`
	Info  string `json:"info" form:"info"`
}

func (s *ListVideoServics) List() serializer.Response {
	videos := []model.Video{}
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "查询视频数据失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideos(videos)}
}
