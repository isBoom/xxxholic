package api

import (
	"github.com/gin-gonic/gin"
	"xxxholic/service"
)

func VideoRank(c *gin.Context) {
	s := service.VideoRankService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := s.Get()
		c.JSON(200, res)
	}
}
