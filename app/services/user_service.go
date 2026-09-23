package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/config"
	"lost-and-found-backend/configs/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 检查用户名或手机号是否已存在
func CheckRegisterUserExists(username string, phoneNum string) (bool, error) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("username = ? OR phone_num = ?", username, phoneNum).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// 注册用户，将用户信息存入数据库
func Register(username string, phoneNum string, password string, role string, inviteCode string) (*models.User, error) {
	var user models.User
	//根据邀请码判断是否有权限注册管理员
	if role == "系统管理员" || role == "失物招领管理员" {
		if inviteCode != config.Config.GetString("register.admin_invite_code") {
			return nil, errs.ErrNoPermission
		}
	}

	userExists, err := CheckRegisterUserExists(username, phoneNum)
	if err != nil {
		return nil, errs.ErrUserCheckFail //用户信息校验失败
	}
	if userExists {
		return nil, errs.ErrUserExists //用户已存在
	}

	hashpassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errs.ErrHashPassword //密码加密失败
	}

	user.Username = username
	user.PhoneNum = phoneNum
	user.Password = hashpassword
	user.Role = role
	err = database.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return nil, errs.ErrDatabase //存储失败
	}

	return &user, nil //注册成功
}

// ==============================================================================================================
// 登录时用电话号检查用户是否已存在
func CheckUserExistsByPhoneNum(phoneNum string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("phone_num = ?", phoneNum).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errs.ErrUserNotFound
	} else if err != nil {
		return nil, errs.ErrDatabase
	} else {
		return &user, nil
	}
}

// 校对密码
func CheckPassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

// 登录
func Login(phoneNum string, password string) (*models.User, error) {
	user, err := CheckUserExistsByPhoneNum(phoneNum)
	if err != nil {
		return nil, err
	}

	result := CheckPassword(user.Password, password)
	if !result {
		return nil, errs.ErrWrongPassword
	}

	return user, nil

}

// ==============================================================================================================

// 检查修改后的用户名或手机号是否已存在
func CheckUserExists(userID int, username string, phoneNum string) (bool, error) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("(username = ? OR phone_num = ?) AND user_id != ?", username, phoneNum, userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// 获取用户个人信息
func GetProfile(userID int) (*models.User, error) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &user, nil
}

// 修改用户个人信息
func UpdateProfile(userID int, username string, phoneNum string) (*models.User, error) {
	var user models.User

	//检查用户名或手机号是否已存在
	userExists, err := CheckUserExists(userID, username, phoneNum)
	if err != nil {
		return nil, errs.ErrDatabase
	}
	if userExists {
		return nil, errs.ErrUserExists
	}

	updateUser := make(map[string]interface{})
	if username != "" {
		updateUser["username"] = username
	}
	if phoneNum != "" {
		updateUser["phone_num"] = phoneNum
	}

	//更新用户信息
	err = database.DB.Model(&models.User{}).Where("user_id = ?", userID).Updates(updateUser).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	//查询更新后的用户信息
	err = database.DB.Model(&models.User{}).Where("user_id = ?", userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errs.ErrUserNotFound
	}
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &user, nil
}
