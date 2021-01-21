package api

import (
	"github.com/gin-gonic/gin"
	"xxxholic/service"
)

func AddComment(c *gin.Context) {
	user := CurrentUser(c)
	s := service.VideoCommentService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Add(user))
	}
}
func DelComment(c *gin.Context) {
	user := CurrentUser(c)
	s := service.VideoCommentService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Del(user))
	}
}
func GetComments(c *gin.Context) {
	s := service.VideoCommentService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Get(c))
	}
}
