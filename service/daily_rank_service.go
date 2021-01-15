package service

import (
	"fmt"
	"strings"
	"xxxholic/cache"
	"xxxholic/model"
	"xxxholic/serializer"
)

type DailyRankService struct {
}

func (s *DailyRankService) Get() serializer.Response {
	var videos []model.Video
	vds, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 10).Result()
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
