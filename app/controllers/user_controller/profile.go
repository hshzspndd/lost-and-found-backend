package user_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 参数绑定结构体
type UpdateProfileData struct {
	Username string `json:"username"`
	PhoneNum string `json:"phone_num"`
}

// 返回响应结构体
type ProfileResp struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	PhoneNum string `json:"phone_num"`
	Role     string `json:"role"`
}

// 获取用户个人信息
func GetProfile(c *gin.Context) {
	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)
	user, err := services.GetProfile(claims.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ProfileResp{
		UserID:   user.UserID,
		Username: user.Username,
		PhoneNum: user.PhoneNum,
		Role:     user.Role,
	}

	utils.ResponseSuccess(c, resp)
}

//==============================================================================================================

// 修改用户个人信息
func UpdateProfile(c *gin.Context) {
	var updateProfileData UpdateProfileData
	err := c.ShouldBindJSON(&updateProfileData)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)
	user, err := services.UpdateProfile(claims.UserID, updateProfileData.Username, updateProfileData.PhoneNum)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ProfileResp{
		UserID:   user.UserID,
		Username: user.Username,
		PhoneNum: user.PhoneNum,
		Role:     user.Role,
	}

	utils.ResponseSuccess(c, resp)
}
