package service

import (
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type ListVideoServics struct {
	Title     string `json:"title" form:"title"`
	Info      string `json:"info" form:"info"`
	VideoType string `json:"videoType" form:"videoType"`
}

func (s *ListVideoServics) List() serializer.Response {
	videos := []model.Video{}
	where:=[]string{}
	var count uint64
	if s.Info!=""{
		where = append(where," info like '%"+s.Info+"%' ")
	}
	if s.Title!=""{
		where = append(where," title like '%"+s.Title+"%' ")
	}
	if err := model.DB.Where(&model.Video{
		VideoType:s.VideoType,
	}).Where(strings.Join(where, "or")).
		Find(&videos).Count(&count).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "查询视频数据失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
		Count:count,
	}
}
