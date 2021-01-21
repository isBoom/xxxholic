package service

import (
	"fmt"
	"math/rand"
	"time"
	"xxxholic/cache"
	"xxxholic/serializer"
)

type GetCaptcha struct {
	Email string `json:"email" form:"email"`
}
func (s *GetCaptcha) GetCaptcha() serializer.Response{
	t,_:=cache.RedisClient.PTTL(cache.GetCaptchaTime(s.Email)).Result()
	if t.Seconds() > 0 {
		return serializer.Response{
			Code:40001,
			Msg:"您点击的太快了,请" + fmt.Sprintf("%v",t.Seconds()) + "后再试",
			Count:uint64(t.Seconds()),
		}
	}
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano()+int64(i))).Intn(len(letterRunes))]
	}
	e:=SendEmail{
		Email: s.Email,
		Msg:   "验证码:"+string(b),
	}
	if err:=e.Send();err!=nil{
		return serializer.Response{
			Code:40001,
			Msg:"验证码发送失败，请稍后再试",
			Error:err.Error(),
		}
	}else{
		//验证码有效期
		cache.RedisClient.Set(cache.GetCaptcha(s.Email),string(b),30 * time.Minute)
		//一分钟内拒绝再发送验证码
		cache.RedisClient.Set(cache.GetCaptchaTime(s.Email),string(b),time.Minute)
		return serializer.Response{}
	}
}