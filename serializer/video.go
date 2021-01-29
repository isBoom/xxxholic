package serializer

import "xxxholic/model"

type Video struct {
	User      User   `json:"user"`
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
	Url       string `json:"url"`
	VideoType string `json:"videoType"`
	Status string `json:"status"`
}

func BuildVideo(item model.Video) Video {
	user, _ := model.GetUser(item.UserId)
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
		Url:       item.VideoUrl(),
		Avatar:    item.AvatarUrl(),
		User:      BuildUser(user),
		View:      item.GetView(),
		VideoType: item.VideoType,
		Status:    item.Status,
	}
}
func BuildVideos(item []model.Video) (videos []Video) {
	for _, value := range item {
		video := BuildVideo(value)
		videos = append(videos, video)
	}
	return videos
}
