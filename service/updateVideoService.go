package service

import (
	"fmt"
	"xxxholic/model"
	"xxxholic/serializer"
)

type UpdateVideoService struct {
	ID uint `json:"id" form:"id"`
	Title string `json:"title" form:"title" binding:"required,min=2,max=50"`
	Info  string `json:"info" form:"info" binding:"max=500"`
	Avatar string `json:"avatar" form:"avatar"`
	VideoType string `json:"videoType" form:"videoType"`
}

func (s *UpdateVideoService) Update() serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, s.ID).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "请求视频不存在",
			Error: err.Error(),
		}
	}
	video.Title = s.Title
	video.Info = s.Info
	video.Avatar = s.Avatar
	video.Status = "audit"
	video.VideoType = s.VideoType

	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "视频信息保存失败",
			Error: err.Error(),
		}
	}
	fmt.Println(s)
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
