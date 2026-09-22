package userController

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 接收参数
type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseLoginData struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	UserId    int       `json:"user_id"`
	Username  string    `json:"username"`
	PhoneNum  string    `json:"phone_num"`
	Role      string    `json:"role"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	user, err := services.Login(data.PhoneNum, data.Password)
	if err != nil {
		c.Error(err) // 将错误存入Error交由异常响应中间件处理
		c.Abort()
		return
	}

	token, expiredAt, err := utils.GenerateJWT(user.UserId, user.Username, user.Role)
	if err != nil {
		c.Error(errs.ErrGenerateToken)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, ResponseLoginData{
		Token:     token,
		ExpiredAt: expiredAt,
		UserId:    user.UserId,
		Username:  user.Username,
		PhoneNum:  user.PhoneNum,
		Role:      user.Role,
	})

}
