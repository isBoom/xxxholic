package middleware

import (
	"github.com/gin-gonic/gin"
	"xxxholic/model"
	"xxxholic/serializer"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				if _, ook := model.AdminList[user.(*model.User).ID];ook{
					c.Next()
					return
				}
			}
		}
		c.JSON(200, serializer.CheckAdmin())
		c.Abort()
	}
}