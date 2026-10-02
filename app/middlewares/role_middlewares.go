package middlewares

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 限制身份
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, _ := c.Get("claims")
		claims := val.(*utils.Claims)

		for _, role := range roles {
			if claims.Role == role {
				c.Next()
				return
			}
		}

		c.Error(errs.ErrNoPermission)
		c.Abort()
	}
}
