package sysadmin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MuteUser 禁言用户
// muteSecond：=0 →永久禁言；>0 限时禁言，禁多少秒

func MuteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	muteSecomd, err := strconv.Atoi(c.Param("mute_secomd"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}
	err = services.MuteUser(userID, int64(muteSecomd))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "操作成功")
}
