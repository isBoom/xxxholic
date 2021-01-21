package service

import (
	"xxxholic/cache"
	"xxxholic/model"
	"xxxholic/serializer"
)

type UserChangePasswordService struct {
	Email        string `form:"email" json:"email" binding:"required,min=5,max=30,email"`
	Captcha string `form:"captcha" json:"captcha"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=20"`
	PasswordConfirm string `form:"passwordConfirm" json:"passwordConfirm" binding:"required,min=6,max=20"`
}
func (s *UserChangePasswordService) UserChangePassword() serializer.Response {
	if t,err:=cache.RedisClient.Get(cache.GetCaptcha(s.Email)).Result();err!=nil{
		return serializer.Response{
			Code:40001,
			Msg:"验证码无效",
			Error:err.Error(),
		}
	}else if t == "" {
		return serializer.Response{
			Code:40001,
			Msg:"验证码无效",
		}
	}else if t != s.Captcha{
		return serializer.Response{
			Code:40001,
			Msg:"验证码错误",
		}
	}else{
		if err:=s.ChangePassword();err!=nil{
			return serializer.Response{
				Code:40001,
				Msg:"更改密码失败",
				Error:err.Error(),
			}
		}else{
			cache.RedisClient.Del(cache.GetCaptcha(s.Email))
			return serializer.Response{
				Code:0,
				Msg:"更改密码成功",
			}
		}
	}
}
//无验证改密码
func (s *UserChangePasswordService) ChangePassword() error{
	user:=&model.User{Email:s.Email,}
	if err := user.SetPassword(s.Password); err != nil {
		return err
	}
	if err :=model.DB.Table("users").Where("email = ?",s.Email).Update(&user).Error ;err != nil {
		return err
	}
	return nil
}
