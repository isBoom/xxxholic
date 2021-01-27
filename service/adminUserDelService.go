package service

import (
	"fmt"
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminUserDelService struct {
	Ids string `json:"ids" form:"ids" binding:"required`
}
func (s *AdminUserDelService) UserDel() serializer.Response{
	fmt.Println(s)
	id :=strings.Split(s.Ids,",")
	fmt.Println(id)
	if err:=model.DB.Where("id in (?)",id).Delete(model.User{}).Error;err!=nil{
		return serializer.ParamErr("删除失败",err)
	}
	return serializer.Response{}
}
