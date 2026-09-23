package user_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 接收参数
type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(services.ErrBindJSON)
		c.Abort()
		return
	}

	user, err := services.Login(data.PhoneNum, data.Password)
	if err != nil {
		c.Error(err) // 将错误存入Error交由异常响应中间件处理
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, user)

}
