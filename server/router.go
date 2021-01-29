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
		v1.POST("user/changePassword",api.ChangePassword)
		v1.POST("user/changeSignature",api.ChangeSignature)
		v1.GET("videos", api.ListVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("getCaptcha",api.GetCaptcha)
		//	//v1.GET("user/:id", api.OtherUserInfo)
		v1.GET("rank/video", api.VideoRank)
		v1.GET("video/:id/comments", api.GetComments)
	}
	// 需要登录保护的
	auth := v1.Group("")
	auth.Use(middleware.AuthRequired())
	{
		//用户类
		auth.GET("user/me", api.UserMe)
		auth.DELETE("user/logout", api.UserLogout)
		auth.POST("user/changeAvatar", api.ChangeAvatar)
		auth.GET("user/history",api.GetHistory)
		//auth.GET("user/notAuditVideo",api.notAuditVideo)

		//视频类
		auth.POST("upload/tokenAvatar", api.UploadAvatarToken)
		auth.POST("upload/tokenVideo", api.UploadVideoToken)
		auth.POST("video/updateVideo", api.UpdateVideo)
		auth.POST("videos", api.CreateVideo)
		auth.POST("video/comment", api.AddComment)
		auth.DELETE("video/delComment", api.DelComment)

		//管理员专用接口
		admin:=auth.Group("/admin")
		admin.Use(middleware.Admin())
		{
			//用户相关
			admin.GET("user/me", api.UserMe)
			admin.GET("users",api.AdminUserList)
			admin.DELETE("user/delUser",api.AdminDelUser)
			admin.POST("user/update",api.AdminUserUpdate)
			admin.POST("user/create",api.AdminUserCreate)
			//视频相关
			admin.GET("videos",api.AdminVideoList)
			admin.DELETE("video/delVideo",api.AdminDelVideo)
			admin.POST("video/updateVideo",api.AdminVideoUpdate)
		}
	}
	return r
}
