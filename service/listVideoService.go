package service

import (
	"fmt"
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type ListVideoServics struct {
	Title     string `json:"title" form:"title"`
	Info      string `json:"info" form:"info"`
	VideoType string `json:"videoType" form:"videoType"`
	Limit     uint64 `json:"limit" form:"limit"`
	Offset    uint64 `json:"offset" form:"offset"`
	UserId    uint   `json:"userId" form:"userId"`
	Status string `json:"status" form:"status"`
}

func (s *ListVideoServics) List(args ... interface{}) serializer.Response {
	videos := []model.Video{}
	where := []string{}
	var count uint64
	if s.Info != "" {
		where = append(where, " info like '%"+s.Info+"%' ")
	}
	if s.Title != "" {
		where = append(where, " title like '%"+s.Title+"%' ")
	}
	if s.Limit == 0 {
		s.Limit = 12
	}
	if s.UserId ==0 || len(args)==0{
		s.Status = "normal"
	}else{
		s.UserId = args[0].(uint)
		fmt.Println(s)
	}
	if err := model.DB.Where(&model.Video{
		UserId:s.UserId,
		VideoType: s.VideoType,
		Status:s.Status,
	}).Where(strings.Join(where, "or")).
		Find(&videos).Count(&count).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "查询视频数据失败",
			Error: err.Error(),
		}
	}

	if err := model.DB.Where(&model.Video{
		UserId:s.UserId,
		VideoType: s.VideoType,
		Status:s.Status,
	}).Where(strings.Join(where, "or")).Offset(s.Offset).Limit(s.Limit).
		Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "查询视频数据失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:  serializer.BuildVideos(videos),
		Count: count,
	}
}
