package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Email        string `form:"email" json:"email" binding:"required,min=5,max=30,email"`
	UserName        string `form:"userName" json:"userName" binding:"required,min=2,max=12"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=20"`
	PasswordConfirm string `form:"passwordConfirm" json:"passwordConfirm" binding:"required,min=6,max=20"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("email = ?", service.Email).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "邮箱已经注册",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已存在",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName: service.UserName,
		Email: service.Email,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}
	return serializer.BuildUserResponse(user)
}
