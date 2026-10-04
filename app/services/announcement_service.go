package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
)

// 查询公告列表
func GetAnnouncements(page int, pageSize int) ([]models.Announcement, int, error) {
	var total int64
	announcements := make([]models.Announcement, 0)

	query := database.DB.Model(&models.Announcement{})

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	// 2. 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&announcements).Error; err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return announcements, int(total), nil
}
