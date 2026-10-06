package sysadmin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewRole struct {
	Role string `json:"role" binding:"required,oneof=普通用户 失物招领管理员 系统管理员"`
}

// 修改用户角色
func UpdateRole(c *gin.Context) {
	var newRole NewRole
	err := c.ShouldBindJSON(&newRole)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)
	operatorID := claims.UserID

	err = services.UpdateRole(operatorID, userID, newRole.Role)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)
}
