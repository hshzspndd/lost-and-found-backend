package user_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordData struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func UpdatePassword(c *gin.Context) {
	var updatePassword UpdatePasswordData
	err := c.ShouldBindJSON(&updatePassword)
	if err != nil {
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
