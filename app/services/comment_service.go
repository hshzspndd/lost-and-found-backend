package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// ================================================== 发布评论 ==================================================

func CreateComment(postID, userID int, content string) (*models.Comment, error) {
	var comment = models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
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

// ================================================== 删除评论 ==================================================

// 删除自己的评论
func DeleteMyComment(CommentID, UserID int) error {
	res := database.DB.Model(&models.Comment{}).Where("comment_id = ? AND user_id = ?", CommentID, UserID).Delete(&models.Comment{})

	if res.Error != nil {
		return errs.ErrDatabase
	}

	if res.RowsAffected == 0 {
		commentExists, err := CheckCommentExistByCommentID(CommentID)
		if err != nil {
			return errs.ErrDatabase
		}
		if !commentExists {
			return errs.ErrCommentNotFound
		}
		return errs.ErrIsNotYourComment
	}
	return nil
}

func CheckCommentExistByCommentID(commentID int) (bool, error) {
	var comment models.Comment
	err := database.DB.Model(&models.Comment{}).Where("comment_id = ?", commentID).First(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
