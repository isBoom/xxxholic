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
	Limit int64 `json:"limit" form:"limit"`
	Offset int64 `json:"offset" form:"offset"`
}

func (s *VideoRankService) Get() serializer.Response {
	var videos []model.Video
	rankName := cache.GetRankName(cache.GetType(s.RankType), s.VideoType)
	if s.Limit == 0 {
		s.Limit = 10
	}
	vds, _ := cache.RedisClient.ZRevRange(rankName, s.Offset, s.Limit).Result()
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
