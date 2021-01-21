package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xxxholic/serializer"
	"xxxholic/service"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Login(c)
		c.JSON(200, res)
	}
}
// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, service.Register())
	}
}
// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
func ChangePassword(c *gin.Context){
	var service service.UserChangePasswordService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.UserChangePassword()
		c.JSON(200, res)
	}
}
func ChangeSignature(c *gin.Context){
	var service service.UserChangeSignatureService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.UserChangeSignature(c)
		c.JSON(200, res)
	}
}


