package server

import (
	"os"
	"xxxholic/api"
	"xxxholic/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/register", api.UserRegister)
		v1.POST("videos", api.ListVideo)
		v1.GET("video/:id", api.ShowVideo)
		//	//v1.GET("user/:id", api.OtherUserInfo)
		//
		v1.POST("rank/video", api.VideoRank)
		//	v1.GET("video/:id/comments", api.GetComments)
	}
	// 需要登录保护的
	auth := v1.Group("")
	auth.Use(middleware.AuthRequired())
	{
		//用户类
		auth.GET("user/me", api.UserMe)
		auth.DELETE("user/logout", api.UserLogout)
		//视频类
		//auth.POST("upload/tokenAvatar", api.UploadAvatarToken)
		//auth.POST("upload/tokenVideo", api.UploadVideoToken)
		//auth.PUT("videos/:id", api.UpdateVideo)
		//auth.POST("videos", api.CreateVideo)
		//auth.POST("video/comment", api.AddComment)
		//auth.DELETE("video/delComment", api.DelComment)
	}
	return r
}
