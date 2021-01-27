package service

import (
	"xxxholic/model"
	"xxxholic/serializer"
)

// CreateVideoService 上传视频
type CreateVideoService struct {
	Title  string `json:"title" form:"title" binding:"required,min=2,max=50"`
	Info   string `json:"info" form:"info" binding:"max=500"`
	Url    string `json:"url" form:"url" binding:"required`
	Avatar string `json:"avatar" form:"avatar"`
	UserId uint   `json:"userId" form:"userId"`
	VideoType string `json:"videoType" form:"videoType"`
	Status string
}

func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	if user == nil {
		return serializer.Response{
			Code: 5001,
			Msg:  "未登录，请登陆后再上传视频",
		}
	}
	for _, value := range model.VideoTypes {
		if value == service.VideoType{
			v := model.Video{
				Title:  service.Title,
				Info:   service.Info,
				Url:    service.Url,
				Avatar: service.Avatar,
				UserId: user.ID,
				VideoType:service.VideoType,
				Status: "notAudit",
			}
			if err := model.DB.Create(&v).Error; err != nil {
				return serializer.Response{
					Code:  serializer.CodeParamErr,
					Msg:   "视频保存失败",
					Error: err.Error(),
				}
			}else{
				return serializer.Response{Data: serializer.BuildVideo(v)}
			}

		}
	}
	return serializer.Response{Code:serializer.CodeParamErr,Msg:"请选择正确的视频类型",Error:service.VideoType}
}
