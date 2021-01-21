package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

// CreateVideoService 上传视频
type CreateVideoService struct {
	Title  string `json:"title" form:"title" binding:"required,min=2,max=50"`
	Info   string `json:"info" form:"info" binding:"min=2,max=500"`
	Url    string `json:"url" form:"url" binding:"required`
	Avatar string `json:"avatar" form:"avatar"`
	UserId uint   `json:"userId" form:"userId"`
}

func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	if user == nil {
		return serializer.Response{
			Code: 5001,
			Msg:  "未登录，请登陆后再上传视频",
		}
	}
	v := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		Url:    service.Url,
		Avatar: service.Avatar,
		UserId: user.ID,
	}

	if err := model.DB.Create(&v).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(v)}
}
