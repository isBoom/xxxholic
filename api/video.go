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
		c.JSON(200, service.List())
	}
}
func ShowVideo(c *gin.Context) {
	s := service.ShowVideoServics{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Show(c.Param("id")))
	}
}
