package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xxxholic/model"
	"xxxholic/serializer"
	"strconv"
	"time"
)

type VideoCommentService struct {
	ID           uint      `json:"id" form:"id"`
	UserId       uint      `json:"-" form:"-"`
	VideoId      uint      `json:"videoId" form:"videoId"`
	FirstId      uint      `json:"firstId" form:"firstId"` //最终挂载哪条评论下
	Content      string    `json:"content" form:"content"`
	ParentId     uint      `json:"parentId" form:"parentId"` //直接父级评论
	Count        uint      `json:"count" json:"count"`
	ParentUserId uint      `json:"-" form:"-"`
	CreatedAt    time.Time `json:"-" form:"-"`
	limitStart   uint
	limitEnd     uint
}
type Comments struct {
	User       serializer.User `json:"user" form:"user"`
	ParentUser serializer.User `json:"parentUser" form:"parentUser"`
	VideoCommentService
	CreatedAtInt64 int64      `json:"createdAt" form:"createdAt"`
	Child          []Comments `json:"child" form:"child"`
}
type ResData struct {
	Count    uint        `json:"count" form:"count"`
	Comments interface{} `json:"comments" form:"comments"`
}

func (s *VideoCommentService) Del(user *model.User) serializer.Response {
	if user == nil {
		return serializer.Response{
			Code: 5001,
			Msg:  "未登录",
		}
	}
	mod := model.Comment{}
	if err := model.DB.Model(model.Comment{}).Where("id = ? and user_id = ?", s.ID, user.ID).First(&mod).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Error: err.Error(),
			Msg:   "这不是您的评论",
		}
	}
	if err := model.DB.Delete(&mod).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Error: err.Error(),
			Msg:   "删除失败",
		}
	} else if mod.FirstId == 0 {
		if err := model.DB.Delete(model.Comment{}, "first_id = ?", s.ID).Error; err != nil {
			return serializer.Response{
				Code:  5001,
				Error: err.Error(),
				Msg:   "删除子评论失败",
			}
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "删除成功",
	}
}
func (s *VideoCommentService) Add(user *model.User) serializer.Response {
	fmt.Printf("%+v", s)
	if user == nil {
		return serializer.Response{
			Code: 5001,
			Msg:  "未登录，请登陆后再进行评论",
		}
	}
	if s.ParentId != 0 {
		com := model.Comment{}
		if err := model.DB.Where("id = ? and video_id = ?", s.ParentId, s.VideoId).First(&com).Error; err != nil {
			return serializer.Response{
				Msg: "回复的评论不存在",
			}
		}
		if com.ParentId == 0 {
			s.FirstId = com.ID
		} else {
			s.FirstId = com.FirstId
		}
		s.ParentUserId = com.UserId
	}
	c := model.Comment{
		UserId:       user.ID,
		VideoId:      s.VideoId,
		ParentId:     s.ParentId,
		ParentUserId: s.ParentUserId,
		Content:      s.Content,
		FirstId:      s.FirstId,
	}
	if err := model.DB.Create(&c).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "评论失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildComment(c)}
}

func (s *VideoCommentService) Get(c *gin.Context) serializer.Response {
	//子评论
	mods := make([]Comments, 0)
	//父评论
	res := make([]Comments, 0)
	//用户表
	user := make([]model.User, 0)
	//构建评论所需的map
	mapUser := make(map[uint]model.User, 0)
	var count uint
	if err := model.DB.Model(Comments{}).Where("video_id = ? and deleted_at IS NULL", c.Param("id")).Count(&count).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取评论数失败",
			Error: err.Error(),
		}
	}
	//所有父评论
	if err := model.DB.Where("video_id = ? and parent_id = 0 and deleted_at IS NULL", c.Param("id")).Order("id desc").Find(&res).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取父评论失败",
			Error: err.Error(),
		}
	}
	if err := model.DB.Where("video_id = ? and parent_id != 0 and deleted_at IS NULL", c.Param("id")).Order("id desc").Find(&mods).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取子评论失败",
			Error: err.Error(),
		}
	}
	//构建用户表
	for _, value := range res {
		mapUser[value.UserId] = model.User{}
	}
	for _, value := range mods {
		mapUser[value.UserId] = model.User{}
		mapUser[value.ParentId] = model.User{}
	}
	//拼接userId
	userIdString := make([]string, 0)
	for key, _ := range mapUser {
		userIdString = append(userIdString, strconv.Itoa(int(key)))
	}
	if err := model.DB.Where("id in (?)", userIdString).Find(&user).Error; err != nil {
		return serializer.Response{
			Code:  5001,
			Msg:   "获取用户信息失败",
			Error: err.Error(),
		}
	}
	//构建map
	for key, value := range user {
		mapUser[value.ID] = user[key]
	}
	mapIndex := make(map[uint]uint)
	for i, mod := range res {
		//添加时间戳
		res[i].CreatedAtInt64 = res[i].CreatedAt.Unix()
		//添加用户信息
		res[i].User = serializer.BuildUser(mapUser[mod.UserId])
		//构建下标信息
		mapIndex[mod.ID] = uint(i)
	}
	//实现二维化评论
	for _, mod := range mods {
		mod.User = serializer.BuildUser(mapUser[mod.UserId])
		mod.ParentUser = serializer.BuildUser(mapUser[mod.ParentUserId])
		mod.CreatedAtInt64 = mod.CreatedAt.Unix()
		if _, ok := mapIndex[mod.FirstId]; ok {
			if res[mapIndex[mod.FirstId]].Child == nil {
				res[mapIndex[mod.FirstId]].Child = make([]Comments, 0)
			}
			res[mapIndex[mod.FirstId]].Child = append(res[mapIndex[mod.FirstId]].Child, mod)
		}

	}
	return serializer.Response{Data: ResData{
		Count:    count,
		Comments: res,
	}}
}
