package middlewares

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 限制身份
// 每次从数据库实时读取用户最新角色再判断，避免旧 token 中的角色在权限变更后继续生效
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, _ := c.Get("claims")
		claims := val.(*utils.Claims)

		// 实时查库获取最新角色
		var user models.User
		err := database.DB.Model(&models.User{}).Where("user_id = ?", claims.UserID).First(&user).Error
		if err == gorm.ErrRecordNotFound {
			c.Error(errs.ErrUnauthorized)
			c.Abort()
			return
		}
		if err != nil {
			c.Error(errs.ErrDatabase)
			c.Abort()
			return
		}

		for _, role := range roles {
			if user.Role == role {
				c.Next()
				return
			}
		}

		c.Error(errs.ErrNoPermission)
		c.Abort()
	}
}
