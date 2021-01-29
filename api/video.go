package api

import (
	"github.com/gin-gonic/gin"
	"xxxholic/service"
)

func ListVideo(c *gin.Context) {
	var service service.ListVideoServics
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		var userId uint
		if user:=CurrentUser(c); user!=nil{
			userId = user.ID
		}
		c.JSON(200, service.List(userId))
	}
}
func ShowVideo(c *gin.Context) {
	s := service.ShowVideoServics{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		var userId uint
		if user:=CurrentUser(c); user!=nil{
			userId = user.ID
		}
		c.JSON(200, s.Show(c.Param("id"),userId))
	}
}
func CreateVideo(c *gin.Context) {
	user := CurrentUser(c)
	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Create(user))
	}
}
func UpdateVideo(c *gin.Context) {
	s := service.UpdateVideoService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(5001, ErrorResponse(err))
	} else {
		var userId uint
		if user:=CurrentUser(c); user!=nil{
			userId = user.ID
		}
		res := s.Update(userId)
		c.JSON(200, res)
	}
}
func AdminVideoUpdate(c *gin.Context) {
	s := service.UpdateVideoService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(5001, ErrorResponse(err))
	} else {
		res := s.AdminUpdate()
		c.JSON(200, res)
	}
}
func AdminVideoList(c *gin.Context) {
	s := service.AdminVideoListService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(5001, ErrorResponse(err))
	} else {
		res := s.AdminVideoList()
		c.JSON(200, res)
	}
}
func AdminDelVideo(c *gin.Context) {
	s := service.AdminVideoDelService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(5001, ErrorResponse(err))
	} else {
		res := s.VideoDel()
		c.JSON(200, res)
	}
}