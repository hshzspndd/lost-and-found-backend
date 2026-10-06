package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
	"time"

	"gorm.io/gorm"
)

// ================================================== 查询所有用户 ==================================================

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

// ================================================== 修改用户角色 ==================================================

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

// ================================================== 获取系统数据 ==================================================

// 统计数据
type StatsData struct {
	TotalUsers         int64 `json:"total_users"`
	NormalUsers        int64 `json:"normal_users"`
	AdminUsers         int64 `json:"admin_users"`
	TotalPosts         int64 `json:"total_posts"`
	PendingPosts       int64 `json:"pending_posts"`
	ApprovedPosts      int64 `json:"approved_posts"`
	RejectedPosts      int64 `json:"rejected_posts"`
	LostPosts          int64 `json:"lost_posts"`
	FoundPosts         int64 `json:"found_posts"`
	TotalAnnouncements int64 `json:"total_announcements"`
}

func GetStats() (*StatsData, error) {
	var userStats struct {
		TotalUsers  int64
		NormalUsers int64
		AdminUsers  int64
	}

	// ========== 用户统计 ==========
	if err := database.DB.Raw(`
		SELECT 
			COUNT(*) AS total_users,
			COUNT(CASE WHEN role = '普通用户' THEN 1 END) AS normal_users,
			COUNT(CASE WHEN role != '普通用户' THEN 1 END) AS admin_users
		FROM users
	`).Scan(&userStats).Error; err != nil {
		return nil, errs.ErrDatabase
	}

	var postStats struct {
		TotalPosts    int64
		PendingPosts  int64
		ApprovedPosts int64
		RejectedPosts int64
		LostPosts     int64
		FoundPosts    int64
	}

	// ========== 帖子统计 ==========
	if err := database.DB.Raw(`
		SELECT 
			COUNT(*) AS total_posts,
			COUNT(CASE WHEN status = '待审核' THEN 1 END) AS pending_posts,
			COUNT(CASE WHEN status = '已通过' THEN 1 END) AS approved_posts,
			COUNT(CASE WHEN status = '已驳回' THEN 1 END) AS rejected_posts,
			COUNT(CASE WHEN post_type = '寻物' THEN 1 END) AS lost_posts,
			COUNT(CASE WHEN post_type = '招领' THEN 1 END) AS found_posts
		FROM posts
	`).Scan(&postStats).Error; err != nil {
		return nil, errs.ErrDatabase
	}

	var announcementStats struct {
		TotalAnnouncements int64
	}

	// ========== 公告统计 ==========
	if err := database.DB.Raw(`
		SELECT COUNT(*) AS total_announcements
		FROM announcements
	`).Scan(&announcementStats).Error; err != nil {
		return nil, errs.ErrDatabase
	}

	return &StatsData{
		TotalUsers:         userStats.TotalUsers,
		NormalUsers:        userStats.NormalUsers,
		AdminUsers:         userStats.AdminUsers,
		TotalPosts:         postStats.TotalPosts,
		PendingPosts:       postStats.PendingPosts,
		ApprovedPosts:      postStats.ApprovedPosts,
		RejectedPosts:      postStats.RejectedPosts,
		LostPosts:          postStats.LostPosts,
		FoundPosts:         postStats.FoundPosts,
		TotalAnnouncements: announcementStats.TotalAnnouncements,
	}, nil
}

// ================================================== 禁言 ==================================================

func MuteUser(userId int, muteSecond int64) error {
	db := database.DB
	var user models.User
	if err := db.First(&user, userId).Error; err != nil {
		return errs.ErrUserNotFound
	}

	if muteSecond == 0 { //永久禁言
		user.IsMuted = true
		user.MutedUntil = 0
	} else { // 限时禁言：当前时间 + 持续秒数
		user.IsMuted = true
		user.MutedUntil = time.Now().Unix() + muteSecond
	}

	return errs.ErrDatabase
}

// UnMuteUser 解除禁言
func UnMuteUser(userId int) error {
	db := database.DB
	var user models.User
	if err := db.First(&user, userId).Error; err != nil {
		return errs.ErrUserNotFound
	}

	user.IsMuted = false
	user.MutedUntil = 0
	return errs.ErrDatabase
}
