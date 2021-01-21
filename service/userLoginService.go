package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xxxholic/model"
	"xxxholic/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Email string `form:"email" json:"email" binding:"required,min=2,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("userId", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("email = ? or user_name = ?", service.Email,service.Email).First(&user).Error; err != nil {
		return serializer.ParamErr("账号不存在", err)
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
