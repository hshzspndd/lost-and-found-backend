package main

import (
	"lost-and-found-backend/configs/config"
	"lost-and-found-backend/configs/database"
	"lost-and-found-backend/configs/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() //加载viper，参数全部读到程序里
	database.InitDB()   //加载数据库，连接MySQL
	r := gin.Default()  //注册主页

	router.Router(r)      //注册路由
	err := r.Run(":8080") //注册端口
	if err != nil {       //报错反馈
		panic("失物招领服务启动失败" + err.Error())
	}
}
