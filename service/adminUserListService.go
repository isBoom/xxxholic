package service

import (
	"strconv"
	"strings"
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminUserListService struct {
	Email string `json:"email" form:"email"`
	UserName string `json:"userName" form:"userName"`
	Limit int64 `json:"limit" form:"limit"`
	Offset int64 `json:"offset" form:"offset"`
}
func (s *AdminUserListService) AdminUserList() serializer.Response{
	var list []string
	var users []model.User
	var where []string
	var count uint64
	if s.Email != "" {
		where = append(where, " email like '%"+s.Email+"%' ")
	}
	if s.UserName != "" {
		where = append(where, " user_name like '%"+s.UserName+"%' ")
	}
	if s.Limit == 0 {
		s.Limit = 10
	}
	for key, _ := range model.AdminList {
		list = append(list, strconv.Itoa(int(key)))
	}
	if err := model.DB.Where("id not in (?)", list).Where(strings.Join(where, "or")).Find(&users).Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr,"查询失败",err)
	}
	if err := model.DB.Where("id not in (?)", list).Where(strings.Join(where, "or")).Offset(s.Offset).Limit(s.Limit).Find(&users).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr,"查询失败",err)
	}
	return serializer.BuildUsersResponse(users,count)
}
