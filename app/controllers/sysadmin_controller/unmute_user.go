package sysadmin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type UnUserService struct {
	UserID int `json:"user_id" binding:"required"`
}

func UnMuteUser(c *gin.Context) {

	var umt UnUserService
	err := c.ShouldBindJSON(&umt)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	err = services.UnMuteUser(umt.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "操作成功")
}
