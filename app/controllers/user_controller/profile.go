package user_controller

import (
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 获取用户个人信息
func GetProfile(c *gin.Context) {
	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	resp := gin.H{
		"user_id": claims.UserID,
		"role":    claims.Role,
	}

	utils.ResponseSuccess(c, resp)
}
