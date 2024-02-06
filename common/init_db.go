package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init_db() {
	host := viper.GetString("datasource.hostname")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database := viper.GetString("datasource.database")
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database)).Error
	if err != nil {
		panic(err)
	}
}
