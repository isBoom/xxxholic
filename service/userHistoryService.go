package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"xxxholic/cache"
	"xxxholic/model"
	"xxxholic/serializer"
)

type UserHistoryService struct {
	Limit int64 `json:"limit" form:"limit"`
	Offset int64 `json:"offset" form:"offset"`
}
type Range struct {
	Score int64 `json:"score"`
	Member string `json:"member"`
}
func (s *UserHistoryService) UserHistory(user *model.User) serializer.Response{
	hName := cache.GetHistoryName(user.ID)
	if s.Limit == 0 {
		s.Limit = 9 + s.Offset
	}
	vds,_ := cache.RedisClient.ZRevRangeWithScores(hName, s.Offset, s.Limit).Result()
	if len(vds) <=0 {
		return serializer.Response{
			Code:  serializer.CodeParamErr,
			Msg:   "暂无历史记录",
		}
	}
	scores:=[]int64{}
	members:=[]string{}
	for _, value := range vds {
		t,_ := json.Marshal(value)
		r:=Range{}
		json.Unmarshal(t,&r)
		scores = append(scores, int64(r.Score))
		members = append(members,r.Member)
	}
	order := fmt.Sprintf("Field(id,%s)", strings.Join(members, ","))
	videos:=[]model.Video{}
	if err := model.DB.Where("id in (?)", members).Order(order).Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询异常",
			Error: err.Error(),
		}
	}
	for i, _ := range videos {
		videos[i].CreatedAt = time.Unix(scores[i],0)
	}
	return serializer.Response{
		Data:  serializer.BuildVideos(videos),
	}
	//Score
	//Member
	//if len(vds) > 0 {
	//	order := fmt.Sprintf("Field(id,%s)", strings.Join(vds.Member, ","))
	//	if err := model.DB.Where("id in (?)", vds).Order(order).Find(&videos).Error; err != nil {
	//		return serializer.Response{
	//			Code:  50000,
	//			Msg:   "数据库查询异常",
	//			Error: err.Error(),
	//		}
	//	}
	//}
	//a,_:=json.Marshal(vds)
	return serializer.Response{

	}
}
