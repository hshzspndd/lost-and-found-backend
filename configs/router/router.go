package router

import (
	"lost-and-found-backend/app/controllers"
	"lost-and-found-backend/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	pre := "/api"
	api := c.Group(pre)
	api.Use(middlewares.GlobalResponseError)
	{
		api.POST("/register", controllers.Register) //注册
	}
}
