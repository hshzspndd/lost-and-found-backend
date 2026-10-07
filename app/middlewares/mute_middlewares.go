package middlewares

import (
	"lost-and-found-backend/app/errs"
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
			c.Error(errs.ErrUnauthorized)
			c.Abort()
			return
		}
		claims, ok := claimsVal.(*utils.Claims)
		if !ok {
			c.Error(errs.ErrUnauthorized)
			c.Abort()
			return
		}

		var user models.User
		err := database.DB.First(&user, claims.UserID).Error
		if err != nil {
			c.Error(errs.ErrUserNotFound)
			c.Abort()
			return
		}
		now := time.Now().Unix()
		if user.IsMuted && (user.MutedUntil == 0 || user.MutedUntil > now) {
			c.Error(errs.ErrMuted)
			c.Abort()
			return
		}
		if user.IsMuted && user.MutedUntil != 0 && user.MutedUntil <= now {
			database.DB.Model(&models.User{}).
				Where("user_id = ? AND muted_until != 0 AND muted_until <= ?", user.UserID, now).
				Updates(map[string]interface{}{"is_muted": false, "muted_until": 0})
		}
		c.Next()
	}
}
