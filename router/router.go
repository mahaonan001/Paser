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
	return r
}
