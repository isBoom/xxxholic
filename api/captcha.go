package api

import (
	"github.com/gin-gonic/gin"
	"xxxholic/service"
)

func GetCaptcha (c *gin.Context) {
	var service service.GetCaptcha
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		c.JSON(200, service.GetCaptcha())
	}
}