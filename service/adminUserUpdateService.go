package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)
type AdminUserUpdateService struct {
	ID uint `form:"id" json:"id"`
	Email        string `form:"email" json:"email" binding:"email"`
	UserName        string `form:"userName" json:"userName"`
	Password        string `form:"password" json:"password"`
	Signature *string `json:"signature" form:"signature"`
	Avatar string `json:"avatar" form:"avatar"`
	Status string `json:"status" form:"status"`
	Permissions string `json:"permissions" form:"permissions"`
}
func (s *AdminUserUpdateService) AdminUserUpdate() serializer.Response{
	user := model.User{
		Email:          s.Email,
		UserName:       s.UserName,
		Status:         s.Status,
		Avatar:         s.Avatar,
		Signature:      s.Signature,
	}
	if s.Password!="" {
		if err:=user.SetPassword(s.Password);err!=nil{
			return serializer.ParamErr("设置密码失败",err)
		}
	}
	if err:=model.DB.Table("users").Where("id = ?",s.ID).Update(&user).Error;err!=nil{
		return serializer.Err(serializer.CodeParamErr,"更新失败",err)
	}else{
		//权限相关
		if s.Permissions == "normal"{
			if err:=model.DB.Where("id = ?",s.ID).Delete(&model.Admin{}).Error;err!=nil{
				return serializer.ParamErr("权限删除失败,其他信息更新成功",err)
			}
			delete(model.AdminList, s.ID)
		}else if s.Permissions == "admin"{
			if _,ok := model.AdminList[s.ID] ; !ok{
				if err := model.DB.Create(&model.Admin{UserId: s.ID,}).Error; err != nil {
					return serializer.ParamErr("权限添加失败,其他信息更新成功", err)
				}
			}
			model.AdminList[s.ID] = 1
		}
		return serializer.Response{}
	}
}
