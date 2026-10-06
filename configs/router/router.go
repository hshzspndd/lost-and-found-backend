package router

import (
	"lost-and-found-backend/app/controllers/admin_controller"
	"lost-and-found-backend/app/controllers/announcement_controller"
	"lost-and-found-backend/app/controllers/claim_controller"
	"lost-and-found-backend/app/controllers/comment_controller"
	"lost-and-found-backend/app/controllers/contact_controller"
	"lost-and-found-backend/app/controllers/post_controller"
	"lost-and-found-backend/app/controllers/sysadmin_controller"
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
		api.POST("/register", user_controller.Register)                     //注册
		api.POST("/login", user_controller.Login)                           //登录
		api.GET("/posts", post_controller.GetAllPosts)                      //获取公开帖子列表
		api.GET("/posts/:post_id/comments", comment_controller.GetComments) //获取评论
		api.GET("/announcements", announcement_controller.GetAnnouncements)
	}

	// jwt鉴权路由组
	auth := api.Group("")
	auth.Use(middlewares.ParseJwt())
	{
		// ================================== 关于用户 ==================================
		auth.GET("/user/profile", user_controller.GetProfile)        //获取用户个人信息
		auth.PATCH("/user/profile", user_controller.UpdateProfile)   //修改用户个人信息
		auth.PATCH("/user/password", user_controller.UpdatePassword) //修改密码

		// ================================== 关于帖子 ==================================
		auth.POST("/upload", post_controller.UploadImage)               //上传图片
		auth.POST("/posts", post_controller.CreatePost)                 //发布帖子
		auth.GET("/my/posts", post_controller.GetMyPosts)               //查询自己的帖子
		auth.DELETE("/my/posts/:post_id", post_controller.DeleteMyPost) //删除我的帖子
		auth.GET("/posts/:post_id", post_controller.GetPostDetails)     //获取帖子详情
		auth.PUT("/my/posts/:post_id", post_controller.UpdateMyPost)    //编辑并重新提交帖子

		// ================================== 关于认领 ==================================
		auth.POST("/posts/:post_id/claims", claim_controller.CreateClaim)  //提交认领申请
		auth.DELETE("/claims/:claim_id", claim_controller.CancelClaim)     //撤回认领申请
		auth.GET("/my/claims", claim_controller.GetMyClaims)               // 我提交的申请
		auth.GET("/posts/:post_id/claims", claim_controller.GetPostClaims) // 我帖子收到的申请
		auth.PUT("/claims/:claim_id/audit", claim_controller.AuditClaim)   //处理认领申请

		// ================================== 关于联系人 ==================================
		auth.POST("/contact", contact_controller.AddContact)                      //添加联系人
		auth.GET("/my/contacts", contact_controller.GetContacts)                  //查询联系人列表
		auth.DELETE("/my/contacts/:contact_id", contact_controller.DeleteContact) //删除联系人

		// ================================== 关于评论 ==================================
		auth.POST("/posts/:post_id/comments", comment_controller.CreateComment)     //发布评论
		auth.DELETE("/posts/:post_id/comments", comment_controller.DeleteMyComment) //删除我的评论
	}

	// ================================== 关于失物招领管理员 ==================================
	admin := api.Group("/admin/posts")
	admin.Use(middlewares.ParseJwt())
	admin.Use(middlewares.RequireRole("失物招领管理员", "系统管理员"))
	{
		admin.GET("", admin_controller.GetAdminPosts)               // 获取所有帖子
		admin.PATCH("/:post_id/audit", admin_controller.AuditPost)  //审核帖子
		admin.DELETE("/:post_id", admin_controller.AdminDeletePost) //删除帖子
	}

	// ================================== 关于系统管理员 ==================================
	sysadmin := api.Group("/sys")
	sysadmin.Use(middlewares.ParseJwt())
	sysadmin.Use(middlewares.RequireRole("系统管理员"))
	{
		sysadmin.GET("/users", sysadmin_controller.GetUsers)                                           //查询所有用户
		sysadmin.PATCH("/users/:user_id/role", sysadmin_controller.UpdateRole)                         //修改用户角色
		sysadmin.POST("/announcements", announcement_controller.CreateAnnouncement)                    //发布公告
		sysadmin.DELETE("/announcements/:announcement_id", announcement_controller.DeleteAnnouncement) //删除公告
		sysadmin.GET("/stats", sysadmin_controller.GetStats)                                           //获取系统数据
	}
}
