package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type ShowVideoServics struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

func (v *ShowVideoServics) Show(id string,userId uint) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "未找到资源",
			Error: err.Error(),
		}
	}
	if video.Status !="normal"{
		return serializer.Response{
			Code:  50001,
			Msg:   "此视频暂无法播放",
		}
	}
	video.AddView()
	video.SaveHistory(userId)
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
