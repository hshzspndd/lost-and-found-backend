package services

import (
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// 注册时检查用户是否已存在
func CheckRegisterUserExists(username string) (bool, error) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

// 注册用户，将用户信息存入数据库
func Register(username string, password string, role string) (*models.User, error) {
	var user models.User
	userExists, err := CheckRegisterUserExists(username)
	if err != nil {
		return &models.User{}, ErrUserCheckFail //用户信息校验失败
	}
	if userExists {
		return &models.User{}, ErrUserExists //用户已存在
	}

	hashpassword, err := utils.HashPassword(password)
	if err != nil {
		return &models.User{}, ErrHashPassword //密码加密失败
	}

	user.Username = username
	user.Password = hashpassword
	user.Role = role
	err = database.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return &models.User{}, ErrDatabase //存储失败
	}

	return &user, nil //注册成功
}
