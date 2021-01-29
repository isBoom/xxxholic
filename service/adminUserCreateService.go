package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminUserCreateService struct {
	Email        string `form:"email" json:"email" binding:"required,min=5,max=30,email"`
	UserName        string `form:"userName" json:"userName" binding:"required,min=2,max=12"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Signature *string `json:"signature" form:"signature"`
	Permissions string `json:"permissions" form:"permissions"`
	Avatar string `json:"avatar" form:"avatar"`
}
func (s *AdminUserCreateService) UserCreate() serializer.Response{
	user:=model.User{
		Email:          s.Email,
		UserName:       s.UserName,
		Status:         model.Active,
		Avatar:         s.Avatar,
		Signature:      s.Signature,
	}

	if err := s.valid(); err != nil {
		return *err
	}

	if err := user.SetPassword(s.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	if err:=model.DB.Create(&user).Error;err!=nil{
		return serializer.ParamErr("添加失败",err)
	}

	if s.Permissions == "admin"{
		if err := model.DB.Create(&model.Admin{UserId: user.ID,}).Error; err != nil {
			return serializer.ParamErr("权限添加失败,其他信息更新成功", err)
		}
		model.AdminList[user.ID] = 1
	}

	return serializer.Response{}
}

// valid 验证表单
func (s *AdminUserCreateService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("email = ?", s.Email).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "邮箱已经注册",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", s.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已存在",
		}
	}
	return nil
}
