package router

import (
	"PaSer/common"
	"PaSer/controller"
	"PaSer/mail"
	"PaSer/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	common.Init_db()
	auth := r.Group("/api/admin")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.GET("/info", middleware.AuthMiddleware(), controller.Info)
		auth.POST("/get_code", mail.Code_email)
		auth.POST("/addPaper", controller.AddPaper)
	}
	user := r.Group("/api/user")
	{
		user.POST("/register", controller.Register_User)
		user.POST("/login", controller.Login_User)
		user.GET("/info", middleware.AuthMiddleware(), controller.Info_User)
		user.POST("/get_code", mail.Code_email)
		user.GET("/getpaper", controller.GetPaper)
	}
	return r
}
