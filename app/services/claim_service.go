package services

import (
	"errors"
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
	"time"

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

	// 只有待处理的申请才算重复，被拒绝后可以重新提交
	err = database.DB.Model(&models.Claim{}).Where("post_id = ? AND claimer_id = ? AND status = ?", postID, claimerID, "待处理").Count(&count).Error
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

// ================================================== 查看自己提交的申请 ==================================================

type MyClaimWithPost struct {
	ClaimID   int       `json:"claim_id"`
	PostID    int       `json:"post_id"`
	PostTitle string    `json:"post_title"`
	PostType  string    `json:"post_type"`
	Reason    string    `json:"reason"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// 查询我提交的认领申请
func GetMyClaims(userID int, page int, pageSize int) ([]MyClaimWithPost, int, error) {
	var total int64
	claims := make([]MyClaimWithPost, 0)

	if err := database.DB.Model(&models.Claim{}).
		Where("claimer_id = ?", userID).
		Count(&total).Error; err != nil {
		return nil, 0, errs.ErrDatabase
	}

	offset := (page - 1) * pageSize
	err := database.DB.Table("claims").
		Select("claims.claim_id, claims.post_id, posts.title AS post_title, posts.post_type, claims.reason, claims.status, claims.created_at").
		Joins("LEFT JOIN posts ON posts.post_id = claims.post_id").
		Where("claims.claimer_id = ?", userID).
		Order("claims.created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&claims).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return claims, int(total), nil
}

// ================================================== 查看自己的帖子收到的认领申请 ==================================================

type PostClaimWithUser struct {
	ClaimID        int       `json:"claim_id"`
	ClaimerID      int       `json:"claimer_id"`
	ClaimerName    string    `json:"claimer_name"`
	ClaimerContact string    `json:"claimer_contact"`
	Reason         string    `json:"reason"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

// 查询自己帖子收到的认领申请
func GetPostClaims(postID int, authorID int, page int, pageSize int) ([]PostClaimWithUser, int, error) {
	var total int64
	claims := make([]PostClaimWithUser, 0)

	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, errs.ErrPostNotFound
		}
		return nil, 0, errs.ErrDatabase
	}
	if post.UserID != authorID {
		return nil, 0, errs.ErrNoPermission // 只有作者能看自己帖子的申请
	}

	query := database.DB.Model(&models.Claim{}).Where("post_id = ?", postID)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errs.ErrDatabase
	}

	offset := (page - 1) * pageSize
	err := database.DB.Table("claims").
		Select("claims.claim_id, claims.claimer_id, users.username AS claimer_name, users.phone_num AS claimer_contact, claims.reason, claims.status, claims.created_at").
		Joins("LEFT JOIN users ON users.user_id = claims.claimer_id").
		Where("claims.post_id = ?", postID).
		Order("claims.created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&claims).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return claims, int(total), nil
}

// ================================================== 处理认领申请 ==================================================

// 作者审批认领申请
func AuditClaim(claimID int, authorID int, approve bool) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var claim models.Claim
		if err := tx.First(&claim, claimID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errs.ErrClaimNotFound
			}
			return errs.ErrDatabase
		}

		var post models.Post
		if err := tx.First(&post, claim.PostID).Error; err != nil {
			return errs.ErrPostNotFound
		}
		if post.UserID != authorID {
			return errs.ErrNoPermission // 只有作者能审批自己帖子的申请
		}

		//只有“待处理”的申请能被审批
		if claim.Status != "待处理" {
			return errs.ErrStatusInvalid
		}

		//更新当前申请的状态
		newStatus := "已拒绝"
		if approve {
			newStatus = "已同意"
		}
		if err := tx.Model(&claim).Update("status", newStatus).Error; err != nil {
			return errs.ErrDatabase
		}

		if approve {
			//把帖子标记为“已解决”（条件更新：帖子还没被解决才允许，防止并发同时同意多个申请）
			result := tx.Model(&models.Post{}).
				Where("post_id = ? AND is_resolve != ?", post.PostID, "已解决").
				Update("is_resolve", "已解决")
			if result.Error != nil {
				return errs.ErrDatabase
			}
			if result.RowsAffected == 0 {
				return errs.ErrPostAlreadyResolved
			}
			//把该帖子的其他待处理申请全部拒绝
			if err := tx.Model(&models.Claim{}).
				Where("post_id = ? AND claim_id != ? AND status = ?", claim.PostID, claimID, "待处理").
				Update("status", "已拒绝").Error; err != nil {
				return errs.ErrDatabase
			}
		}

		return nil
	})
}
