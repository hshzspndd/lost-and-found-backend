package userController

import (
	"errors"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 接收参数
type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 200501,
			"msg":  "参数错误",
			"data": nil,
		})
	}

	//用电话号判断用户是否存在
	err = services.CheckUserExistsByPhoneNum(data.PhoneNum)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(401, gin.H{
				"code": 200502,
				"msg":  "用户不存在",
				"data": nil,
			})
		} else {
			c.JSON(401, gin.H{
				"code": 200500,
				"msg":  "其他错误",
				"data": nil,
			})
		}
		return
	}

	//获取用户信息
	var user *models.User
	user, err = services.GetUserByPhoneNum(data.PhoneNum)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 200500,
			"msg":  "数据丢失",
			"data": nil,
		})
		return
	}

	//判断密码是否正确
	flag := services.ComparePassword(data.Password, user.Password)
	if !flag {
		c.JSON(500, gin.H{
			"code": 200503,
			"msg":  "密码错误",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "OK",
		"data": user,
	})

}
