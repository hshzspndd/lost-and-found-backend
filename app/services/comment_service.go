package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// ================================================== 发布评论 ==================================================

func CreateComment(postID, userID int, content string) (*models.Comment, error) {
	// 校验帖子存在且已通过审核，防止对不存在的帖子产生孤儿评论
	var post models.Post
	err := database.DB.Model(&models.Post{}).Where("post_id = ?", postID).First(&post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrPostNotFound
		}
		return nil, errs.ErrDatabase
	}
	if post.Status != "已通过" {
		return nil, errs.ErrStatusInvalid
	}

	var comment = models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	err = database.DB.Model(&models.Comment{}).Create(&comment).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &comment, nil

}

// ================================================== 获取评论 ==================================================

func GetComments(page, postID int, pageSize int) ([]models.Comment, int, error) {
	var total int64

	comments := make([]models.Comment, 0)

	query := database.DB.Model(&models.Comment{}).Where("post_id = ?", postID)

	//获取评论总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return comments, int(total), nil
}
