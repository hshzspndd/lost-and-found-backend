package database

import (
	"fmt"
	"lost-and-found-backend/configs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 数据库初始化
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.GetString("database.user"),
		config.Config.GetString("database.password"),
		config.Config.GetString("database.host"),
		config.Config.GetInt("database.port"),
		config.Config.GetString("database.dbname"),
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
