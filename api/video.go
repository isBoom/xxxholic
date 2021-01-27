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
		res := s.Update()
		c.JSON(200, res)
	}
}

