package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type UserChangeSignatureService struct {
	Signature *string `json:"signature" form:"signature" binding:"max=50"`
}
func (s *UserChangeSignatureService) UserChangeSignature(user *model.User) serializer.Response{
	if err:=model.DB.Table("users").Where("email = ?",user.Email).Update(&model.User{
		Signature:     s.Signature,
	}).Error;err!=nil{
		return serializer.Err(serializer.CodeParamErr,"更新失败",err)
	}else{
		return serializer.Response{}
	}
}