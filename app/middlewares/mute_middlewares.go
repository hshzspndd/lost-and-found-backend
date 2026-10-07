package middlewares

import (
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/database"
	"time"

	"github.com/gin-gonic/gin"
)

func MuteMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		claimsVal, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{
				"Code":    401,
				"message": "未登录，请重新登录",
				"Data":    nil})
			return
		}
		claims, ok := claimsVal.(*utils.Claims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"Code":    401,
				"message": "token解析异常",
				"Data":    nil})
			return
		}

		var user models.User
		err := database.DB.First(&user, claims.UserID).Error
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"Code":    401,
				"message": "用户不存在",
				"Data":    nil})
			return
		}
		now := time.Now().Unix()
		if user.IsMuted && (user.MutedUntil == 0 || user.MutedUntil > now) {
			c.AbortWithStatusJSON(403, gin.H{
				"Code":    401,
				"message": "账号处于禁言状态，禁止发布内容",
				"Data":    nil})
			return
		}
		c.Next()
	}
}
