package database

import (
	"fmt"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 数据库初始化
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.GetString("database.user"),
		config.Config.GetString("database.pass"),
		config.Config.GetString("database.host"),
		config.Config.GetInt("database.port"),
		config.Config.GetString("database.name"),
	)

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	//创建用户和联系人数据表
	err = db.AutoMigrate(&models.User{}, &models.Contact{})
	if err != nil {
		panic("数据库创建失败：" + err.Error())
	}
	DB = db
}
