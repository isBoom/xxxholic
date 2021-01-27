package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type UserChangeAvatarService struct {
	Avatar string `json:"avatar" form:"avatar" binding:"required"`
}
func (s *UserChangeAvatarService) UserChangeAvatar(user *model.User) serializer.Response {

	if err :=model.DB.Table("users").Where("email = ?",user.Email).Update(&model.User{
		Avatar:         s.Avatar,
	}).Error ;err != nil {
		return serializer.Response{
			Code:  serializer.CodeParamErr,
			Msg:   "上传失败",
			Error: err.Error(),
		}
	}else{
		return serializer.Response{}
	}
}