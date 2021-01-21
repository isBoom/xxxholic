package api

import (
	"github.com/gin-gonic/gin"
	"xxxholic/service"
)

func UploadAvatarToken(c *gin.Context) {
	s := service.UploadTokenService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Post("upload/avatar/"))
	}
}
func UploadVideoToken(c *gin.Context) {
	s := service.UploadTokenService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, s.Post("upload/video/"))
	}
}
