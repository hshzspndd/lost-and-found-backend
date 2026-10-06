package sysadmin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// MuteUser 禁言用户
// muteSecond：=0 →永久禁言；>0 限时禁言，禁多少秒
type UserService struct {
	UserID     int   `json:"user_id" binding:"required"`
	MuteSecond int64 `json:"mute_second" binding:"required"`
}

func MuteUser(c *gin.Context) {

	var mt UserService
	err := c.ShouldBindJSON(&mt)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	err = services.MuteUser(mt.UserID, mt.MuteSecond)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "操作成功")
}
