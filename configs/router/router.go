package router

import (
	"lost-and-found-backend/app/controllers/admin_controller"
	"lost-and-found-backend/app/controllers/comment_controller"
	"lost-and-found-backend/app/controllers/contact_controller"
	"lost-and-found-backend/app/controllers/post_controller"
	"lost-and-found-backend/app/controllers/user_controller"
	"lost-and-found-backend/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	c.Static("/images", "./images") //开放静态资源目录

	pre := "/api"
	api := c.Group(pre)
	api.Use(middlewares.GlobalResponseError)
	{
		api.POST("/register", user_controller.Register) //注册
		api.POST("/login", user_controller.Login)       //登录
	}

	//jwt鉴权路由组
	auth := api.Group("")
	auth.Use(middlewares.ParseJwt())
	{
		// ================================== 关于用户 ==================================

		auth.GET("/user/profile", user_controller.GetProfile)        //获取用户个人信息
		auth.PATCH("/user/profile", user_controller.UpdateProfile)   //修改用户个人信息
		auth.PATCH("/user/password", user_controller.UpdatePassword) //修改密码

		// ================================== 关于帖子 ==================================

		auth.POST("/upload", post_controller.UploadImage)            //上传图片
		auth.POST("/post", post_controller.CreatePost)               //发布帖子
		auth.GET("/my-posts", post_controller.GetMyPosts)            //查询自己的帖子
		auth.DELETE("/delete-my-post", post_controller.DeleteMyPost) //删除我的帖子
		auth.GET("/post-details", post_controller.GetPostDetails)    //获取帖子详情

		// ================================== 关于联系人 ==================================
		auth.POST("/contact", contact_controller.AddContact)             //添加联系人
		auth.GET("/my-contacts", contact_controller.GetContacts)         //查询联系人列表
		auth.DELETE("/delete-contact", contact_controller.DeleteContact) //删除联系人

		// ================================== 关于评论 ==================================

		auth.POST("/comment", comment_controller.CreateComment)               //发布评论
		auth.DELETE("/delete-my-comment", comment_controller.DeleteMyComment) //删除我的评论
	}
	auth.GET("/all-posts", post_controller.GetAllPosts)       //获取公开帖子列表
	auth.GET("/all-comments", comment_controller.GetComments) //获取评论

	// ================================== 关于管理员 ==================================

	admin := api.Group("/admin")
	admin.Use(middlewares.ParseJwt())
	admin.Use(middlewares.RequireRole("失物招领管理员", "系统管理员"))
	{
		admin.GET("/posts", admin_controller.GetAdminPosts)              // 获取所有帖子
		admin.PATCH("/posts/:post_id/audit", admin_controller.AuditPost) //审核帖子
	}
}
