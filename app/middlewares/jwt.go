package middlewares

import (
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		authorizationList := strings.SplitN(authorization, " ", 2)
		if len(authorizationList) != 2 || authorizationList[0] != "Bearer" {
			c.Error(GetError(401, "未登录或无效的token"))
			c.Abort()
			return
		}
		tokenString := authorizationList[1]
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(config.Config.GetString("jwt.key")), nil
		})
		if err != nil {
			c.Error(GetError(401, "token解析失败"))
			c.Abort()
			return
		}

		c.Set("claims", token.Claims)
		c.Next()
	}
}
