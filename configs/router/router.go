package router

import (
	"lost-and-found-backend/app/controllers/user_controller"
	"lost-and-found-backend/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	pre := "/api"
	api := c.Group(pre)
	api.Use(middlewares.GlobalResponseError)
	{
		api.POST("/register", user_controller.Register) //注册
		api.POST("/login", user_controller.Login)       //登录
	}
}
