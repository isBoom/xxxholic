package service

import (
	"time"
	"xxxholic/model"
	"xxxholic/serializer"
)

type AdminCommentListService struct {
	Content string `json:"content" form:"content"`
	Limit int64 `json:"limit" form:"limit"`
	Offset int64 `json:"offset" form:"offset"`
}
type Comment struct {
	ID        uint `json:"id" form:"id"`
	UserId       uint `json:"userId" form:"userId"`
	VideoId      uint `json:"videoId" form:"videoId"`
	ParentId     uint `json:"parentId" form:"parentId"`
	FirstId      uint `json:"firstId" form:"firstId"`
	Content      string `json:"content" form:"content"`
	ParentUserId uint `json:"parentUserId" form:"parentUserId"`
	CreatedAt time.Time `json:"-" form:"-"`
	CreatedAtInt64 int64 `json:"createdAt" form:"createdAt"`
}

func (s *AdminCommentListService) CommentList() serializer.Response {
	if s.Limit == 0 {
		s.Limit = 10
	}
	comments:=[]Comment{}
	var count uint64
	if err := model.DB.Where("content like ? and deleted_at IS NULL","%"+s.Content+"%").Find(&comments).Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeDBError,err.Error(),err)
	}

	if err := model.DB.Where("content like ? and deleted_at IS NULL","%"+s.Content+"%").Offset(s.Offset).Limit(s.Limit).Find(&comments).Error; err != nil {
		return serializer.Err(serializer.CodeDBError,err.Error(),err)
	}
	for key, value := range comments {
		comments[key].CreatedAtInt64 = value.CreatedAt.Unix()
	}
	return serializer.Response{
		Data:  comments,
		Count: count,
	}
}