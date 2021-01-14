package service

type UpdateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"required,min=0,max=300"`
}
//
//func (s *UpdateVideoService) Update(id string) serializer.Response {
//	var video model.Video
//	if err := model.DB.First(&video, id).Error; err != nil {
//		return serializer.Response{
//			Code:  5001,
//			Msg:   "请求视频不存在",
//			Error: err.Error(),
//		}
//	}
//	video.Title = s.Title
//	video.Info = s.Info
//	if err := model.DB.Save(&video).Error; err != nil {
//		return serializer.Response{
//			Code:  5001,
//			Msg:   "视频信息保存失败",
//			Error: err.Error(),
//		}
//	}
//	return serializer.Response{Data: serializer.BuildVideo(video)}
//}
