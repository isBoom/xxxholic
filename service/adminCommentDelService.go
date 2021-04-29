package service

import (
	"fmt"
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminCommentDelService struct {
	Ids string `json:"ids" form:"ids" binding:"required`
}
func (s *AdminCommentDelService) CommentDel() serializer.Response{
	id :=strings.Split(s.Ids,",")
	fmt.Println(id)
	if err:=model.DB.Where("id in (?)",id).Delete(model.Comment{}).Error;err!=nil{
		return serializer.ParamErr("删除失败",err)
	}
	return serializer.Response{}
}