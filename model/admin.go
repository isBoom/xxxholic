package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	UserId       uint   `gorm:"unique;not null"`
}
var AdminList = make(map[uint]uint,0)
func (s *Admin) GetAdminList() error{
	count:=0
	admin:=[]Admin{}
	err:=DB.Find(&admin).Count(&count).Error
	if err!=nil{
		return err
	}
	if count <= 0{
		return err
	}
	for _, value := range admin {
		AdminList[value.UserId] = 1
	}
	return err
}
