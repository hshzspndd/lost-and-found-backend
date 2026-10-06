package user_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordData struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=15"`
}

func UpdatePassword(c *gin.Context) {
	var updatePassword UpdatePasswordData
	err := c.ShouldBindJSON(&updatePassword)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	// 去除首尾空格后重新校验，避免空白密码
	updatePassword.OldPassword = strings.TrimSpace(updatePassword.OldPassword)
	updatePassword.NewPassword = strings.TrimSpace(updatePassword.NewPassword)
	if updatePassword.OldPassword == "" || updatePassword.NewPassword == "" {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}
	if len([]rune(updatePassword.NewPassword)) < 6 || len([]rune(updatePassword.NewPassword)) > 15 {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	if updatePassword.OldPassword == updatePassword.NewPassword {
		c.Error(errs.ErrSamePassword)
		c.Abort()
		return
	}

	err = services.UpdatePassword(claims.UserID, updatePassword.OldPassword, updatePassword.NewPassword)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)

}
