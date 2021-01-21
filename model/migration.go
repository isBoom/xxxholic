package model

import (
	"xxxholic/util"
)
//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Video{}).
		AutoMigrate(&VideoType{}).
		AutoMigrate(&Comment{}).
		AutoMigrate(&Admin{})
	//创建视频类型
	for _, value := range VideoTypes {
		if err := DB.Create(&VideoType{VideoType:value}).Error; err != nil {
			util.Log().Println(err.Error())
		}
	}
	//创建管理员
	if err := DB.Create(&Admin{UserId:1}).Error; err != nil {
		util.Log().Println(err.Error())
	}
	//获取全局可用的管理员列表
	if err:=(&Admin{}).GetAdminList();err!=nil {
		util.Log().Println(err.Error())
	}
	DB.Model(&Comment{}).AddForeignKey("video_id", "users(id)", "RESTRICT", "RESTRICT")
	//DB.Model(&Comment{}).AddForeignKey("parent_user_id", "users(id)", "no action", "no action")
	DB.Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Admin{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
