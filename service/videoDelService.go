package service

import (
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type VideoDelService struct {
	Ids string `json:"ids" form:"ids" binding:"required`
}
func (s *VideoDelService) VideoDel(user *model.User) serializer.Response{
	id :=strings.Split(s.Ids,",")
	if err:=model.DB.Where("user_id = ?",user.ID).Where("id in (?)",id).Delete(model.Video{}).Error;err!=nil{
		return serializer.ParamErr("删除失败",err)
	}
	return serializer.Response{}
}


