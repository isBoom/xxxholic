package service

import (
	"fmt"
	"strings"
	"xxxholic/cache"
	"xxxholic/model"
	"xxxholic/serializer"
)

type VideoRankService struct {
	VideoType string `json:"videoType" form:"videoType"`
	RankType  string `json:"rankType" form:"rankType"`
}

func (s *VideoRankService) Get() serializer.Response {
	var videos []model.Video
	rankName := cache.GetRankName(cache.GetType(s.RankType), s.VideoType)
	vds, _ := cache.RedisClient.ZRevRange(rankName, 0, 10).Result()
	if len(vds) > 0 {
		order := fmt.Sprintf("Field(id,%s)", strings.Join(vds, ","))
		if err := model.DB.Where("id in (?)", vds).Order(order).Find(&videos).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库查询异常",
				Error: err.Error(),
			}
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
