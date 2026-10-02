package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
)

// 查询所有用户
func GetAllUsers(page int, role string) ([]models.User, int, error) {
	var total int64
	var pageSize int = 15
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
