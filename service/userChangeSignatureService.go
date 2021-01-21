package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xxxholic/model"
	"xxxholic/serializer"
)

type UserChangeSignatureService struct {
	Signature string `json:"signature" form:"signature" binding:"required,min=1,max=50"`
}
func (s *UserChangeSignatureService) UserChangeSignature(c *gin.Context) serializer.Response{
	user,_:=c.Get("user")
	if _, ok := user.(*model.User); ok {
		fmt.Println("断言成功")
		if err:=model.DB.Table("users").Where("email = ?",user.(*model.User).Email).Update(&model.User{
			Signature:     s.Signature,
		}).Error;err!=nil{
			return serializer.Err(serializer.CodeParamErr,"更新失败",err)
		}else{
			return serializer.Response{}
		}
	}else{
		return serializer.Err(serializer.CodeParamErr,"更新失败,请稍后再试",nil)
	}

}