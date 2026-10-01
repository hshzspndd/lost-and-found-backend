package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
)

// ================================================== 发布评论 ==================================================

func CreateComment(commentID, postID, userID int, content string) (*models.Comment, error) {
	var comment = models.Comment{
		CommentID: commentID,
		PostID:    postID,
		UserID:    userID,
		Content:   content,
	}
	err := database.DB.Model(&models.Contact{}).Create(&comment).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &comment, nil

}

// ================================================== 获取评论 ==================================================

func GetComments(page, postID int) ([]models.Comment, int, error) {
	var total int64
	var pageSize int = 30

	comments := make([]models.Comment, 0)

	query := database.DB.Model(&models.Comment{}).Where("PostID = ?", postID)

	//获取评论总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Offset(offset).Limit(pageSize).Find(&comments).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return comments, int(total), nil
}
