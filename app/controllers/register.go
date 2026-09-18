package controllers

import (
	"lost-and-found-backend/app/middlewares"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type ResponseRegisterData struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// 注册
func Register(c *gin.Context) {
	var registerData RegisterData
	err := c.ShouldBindJSON(&registerData)
	if err != nil {
		c.Error(middlewares.GetError(400, "数据获取失败"))
		c.Abort()
		return
	}

	// 尝试将用户信息存入数据库
	user, err := services.Register(registerData.Username, registerData.Password, registerData.Role) // 返回包含用户id的用户信息以及发生的错误
	if err != nil {
		errResponse, ok := err.(*services.ResponseErrorForm)
		if ok {
			c.Error(middlewares.GetError(errResponse.Code, errResponse.Message)) // 将错误存入Error交由异常响应中间件处理
			c.Abort()
			return
		}
	}

	// 注册成功
	utils.ResponseSuccess(c, ResponseRegisterData{
		UserId:   user.UserId,
		Username: user.Username,
		Role:     user.Role,
	})

}
