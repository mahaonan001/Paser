package main

import (
	"os"

	"PaSer/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// 允许的域名或IP地址
		AllowOrigins: []string{"http://localhost:8080"},
		// 允许的请求方法（GET, POST等）
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// 允许的请求头
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		// 是否允许认证信息跟随请求
		AllowCredentials: true,
	}))
	r = router.CollectRouter(r)
	r.Run(":" + viper.GetString("server.port"))
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
