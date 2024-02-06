package common

import (
	"PaSer/model"
	"PaSer/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db_User() *gorm.DB {
	host := viper.GetString("datasource.hostname")
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
	var c *gin.Context
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		response.Response(c, http.StatusServiceUnavailable, 500, gin.H{"code": 500}, "数据库连接出错")
	}
	db.AutoMigrate(&model.User{})
	return db
}
func Code_email_DB() *gorm.DB {
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
		response.Response(c, http.StatusServiceUnavailable, 500, gin.H{"code": 500}, "数据库连接出错")
		panic(err)
	}
	db.AutoMigrate(&model.EmailCode{})
	return db
}
func GetDB_User() *gorm.DB {
	return db_User()
}
func GetDB_Email() *gorm.DB {
	return Code_email_DB()
}
