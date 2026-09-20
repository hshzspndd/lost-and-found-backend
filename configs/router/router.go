package router

import (
	"lost-and-found-backend/app/controllers/userController"
	"lost-and-found-backend/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	pre := "/api"
	api := c.Group(pre)
	api.Use(middlewares.GlobalResponseError)
	{
		api.POST("/register", userController.Register) //注册
		api.POST("/login", userController.Login)
	}
}
