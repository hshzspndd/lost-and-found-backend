package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// 查询所有用户
func GetAllUsers(page int, role string, pageSize int) ([]models.User, int, error) {
	var total int64
	users := make([]models.User, 0)

	query := database.DB.Model(&models.User{})

	//筛选用户角色
	if role != "" && (role == "普通用户" || role == "失物招领管理员" || role == "系统管理员") {
		query = query.Where("role = ?", role)
	}

	//获取用户总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Order("user_id DESC").Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return users, int(total), nil
}

// 修改用户角色
func UpdateRole(operatorID int, userID int, role string) error {
	// 防止修改自己的角色
	if operatorID == userID {
		return errs.ErrNoPermission
	}

	var user models.User
	err := database.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.ErrUserNotFound
		}
		return errs.ErrDatabase
	}

	if user.Role == role {
		return nil // 静默成功
	}
	err = database.DB.Model(&user).Update("role", role).Error
	if err != nil {
		return errs.ErrDatabase
	}

	return nil
}
