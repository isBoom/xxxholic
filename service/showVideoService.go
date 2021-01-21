package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type ShowVideoServics struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

func (v *ShowVideoServics) Show(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "未找到资源",
			Error: err.Error(),
		}
	}
	video.AddView()
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
