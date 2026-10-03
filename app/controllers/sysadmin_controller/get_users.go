package sysadmin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResp struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	PhoneNum string `json:"phone_num"`
	Role     string `json:"role"`
}

type UsersResp struct {
	List     []UserListResp `json:"list"`
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

// 获取所有用户
func GetUsers(c *gin.Context) {
	role := c.Query("role")
	pageStr := c.DefaultQuery("page", "1")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	if page < 1 {
		page = 1
	}

	userList, total, err := services.GetAllUsers(page, role)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	userListResp := make([]UserListResp, 0, len(userList))
	for _, user := range userList {
		userListResp = append(userListResp, UserListResp{
			UserID:   user.UserID,
			Username: user.Username,
			PhoneNum: user.PhoneNum,
			Role:     user.Role,
		})
	}

	resp := UsersResp{
		List:     userListResp,
		Total:    total,
		Page:     page,
		PageSize: 15,
	}

	utils.ResponseSuccess(c, resp)

}
