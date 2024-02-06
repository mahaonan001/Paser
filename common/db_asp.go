package common

import (
	"PaSer/model"
	"PaSer/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Con_Db_asp() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		var c *gin.Context
		response.FalseRe(c, fmt.Sprintf("err:%s", err), gin.H{"msg": "连接数据库TEST——ASP失败"})
		panic(err)
	}
	db.AutoMigrate(&model.PaperNew{})
	return db
}
func GetDB_ASP() *gorm.DB {
	return Con_Db_asp()
}
