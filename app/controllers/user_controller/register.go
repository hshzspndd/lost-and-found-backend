package user_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username   string `json:"username" binding:"required,min=3,max=15"`
	PhoneNum   string `json:"phone_num" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=系统管理员 失物招领管理员 普通用户"`
	InviteCode string `json:"invite_code"`
}

type RegisterResp struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	PhoneNum string `json:"phone_num"`
	Role     string `json:"role"`
}

// 注册
func Register(c *gin.Context) {
	var registerData RegisterData
	err := c.ShouldBindJSON(&registerData)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	// 检验手机号格式
	if ok := utils.IsValidPhone(registerData.PhoneNum); !ok {
		c.Error(errs.ErrPhoneFormat)
		c.Abort()
		return
	}

	// 尝试将用户信息存入数据库
	user, err := services.Register(registerData.Username, registerData.PhoneNum, registerData.Password, registerData.Role, registerData.InviteCode) // 返回包含用户id的用户信息以及发生的错误
	if err != nil {
		c.Error(err) // 将错误存入Error交由异常响应中间件处理
		c.Abort()
		return
	}

	resp := RegisterResp{
		UserId:   user.UserID,
		Username: user.Username,
		PhoneNum: user.PhoneNum,
		Role:     user.Role,
	}

	// 注册成功
	utils.ResponseSuccess(c, resp)

}
