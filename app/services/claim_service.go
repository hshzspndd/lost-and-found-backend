package services

import (
	"errors"
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// ================================================== 提交认领申请 ==================================================

// 提交认领申请
func CreateClaim(postID int, claimerID int, reason string) (*models.Claim, error) {
	//查帖子是否存在
	var post models.Post
	var count int64
	err := database.DB.First(&post, postID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrPostNotFound
		}
		return nil, errs.ErrDatabase
	}

	//帖子必须“已通过”且“未解决”
	if post.Status != "已通过" {
		return nil, errs.ErrStatusInvalid
	}
	if post.IsResolve == "已解决" {
		return nil, errs.ErrPostAlreadyResolved
	}

	if post.UserID == claimerID {
		return nil, errs.ErrCannotClaimOwnPost
	}

	err = database.DB.Model(&models.Claim{}).Where("post_id = ? AND claimer_id = ?", postID, claimerID).Count(&count).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	if count > 0 {
		return nil, errs.ErrClaimAlreadyExists
	}

	claim := models.Claim{
		PostID:    postID,
		ClaimerID: claimerID,
		Reason:    reason,
		Status:    "待处理",
	}
	if err := database.DB.Create(&claim).Error; err != nil {
		return nil, errs.ErrDatabase
	}

	return &claim, nil
}

// ================================================== 撤回认领申请 ==================================================

// 撤回认领申请
func CancelClaim(claimID int, userID int) error {
	var claim models.Claim
	if err := database.DB.First(&claim, claimID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.ErrClaimNotFound
		}
		return errs.ErrDatabase
	}

	//只能撤回自己的申请
	if claim.ClaimerID != userID {
		return errs.ErrNoPermission
	}

	//只有“待处理”的申请才能撤回
	if claim.Status != "待处理" {
		return errs.ErrStatusInvalid // 已经处理过的申请不能撤回
	}

	if err := database.DB.Delete(&claim).Error; err != nil {
		return errs.ErrDatabase
	}

	return nil
}
