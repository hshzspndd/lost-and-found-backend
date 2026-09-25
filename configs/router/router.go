package router

import (
	"lost-and-found-backend/app/controllers/post_controller"
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

	// ================================== 关于用户 ==================================

	//jwt鉴权路由组
	auth := api.Group("/")
	auth.Use(middlewares.ParseJwt())
	{
		auth.GET("/user/profile", user_controller.GetProfile)        //获取用户个人信息
		auth.PATCH("/user/profile", user_controller.UpdateProfile)   //修改用户个人信息
		auth.PATCH("/user/password", user_controller.UpdatePassword) //修改密码
	}

	// ================================== 关于帖子 ==================================

	post := api.Group("/")
	{
		post.POST("/upload", post_controller.UploadImage) //上传图片
	}
}
