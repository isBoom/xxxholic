package service

import (
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminVideoListService struct {
	Title string `json:"title" form:"title"`
	Info string `json:"info" form:"info"`
	//VideoType string `json:"videoType" form:"videoType"`
	Status string `json:"status" form:"status"`
	UserName string `json:"userName" form:"userName"`
	Offset int64 `json:"offset" form:"offset"`
	Limit int64 `json:"limit" form:"limit"`
}
func (s *AdminVideoListService) AdminVideoList() serializer.Response{
	var videos []model.Video
	var where []string
	var count uint64
	if s.Title != "" {
		where = append(where, " videos.title like '%"+s.Title+"%' ")
	}
	if s.Info != "" {
		where = append(where, " videos.info like '%"+s.Info+"%' ")
	}
	if s.UserName != "" {
		where = append(where, " users.user_name like '%"+s.UserName+"%' ")
	}
	if s.Status == ""{
		s.Status = "normal"
	}

	if s.Limit == 0 {
		s.Limit = 10
	}
	//联合查询导致查询数量时 被软删除的也算
	if err := model.DB.Table("videos").Where("videos.status = ?",s.Status).Where(strings.Join(where, "or")).Where("videos.deleted_at is null").Joins("left join users on videos.user_id = users.id").Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr,"查询失败",err)
	}
	if err := model.DB.Where(strings.Join(where, "or")).Where("videos.status = ?",s.Status).Joins("left join users on videos.user_id = users.id").Offset(s.Offset).Limit(s.Limit).Find(&videos).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr,"查询失败",err)
	}
	return serializer.Response{Data:serializer.BuildVideos(videos),Count:count}
}
