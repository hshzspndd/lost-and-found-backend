package sysadmin_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 用户统计
type UserStats struct {
	TotalUsers  int64 `json:"total_users"`
	NormalUsers int64 `json:"normal_users"`
	AdminUsers  int64 `json:"admin_users"`
}

// 帖子统计
type PostStats struct {
	TotalPosts    int64 `json:"total_posts"`
	PendingPosts  int64 `json:"pending_posts"`
	ApprovedPosts int64 `json:"approved_posts"`
	RejectedPosts int64 `json:"rejected_posts"`
	LostPosts     int64 `json:"lost_posts"`
	FoundPosts    int64 `json:"found_posts"`
}

// 公告统计
type AnnouncementStats struct {
	TotalAnnouncements int64 `json:"total_announcements"`
}

type StatsResp struct {
	UserStats         UserStats         `json:"user_stats"`
	PostStats         PostStats         `json:"post_stats"`
	AnnouncementStats AnnouncementStats `json:"announcement_stats"`
}

func GetStats(c *gin.Context) {
	stats, err := services.GetStats()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := StatsResp{
		UserStats: UserStats{
			TotalUsers:  stats.TotalUsers,
			NormalUsers: stats.NormalUsers,
			AdminUsers:  stats.AdminUsers,
		},
		PostStats: PostStats{
			TotalPosts:    stats.TotalPosts,
			PendingPosts:  stats.PendingPosts,
			ApprovedPosts: stats.ApprovedPosts,
			RejectedPosts: stats.RejectedPosts,
			LostPosts:     stats.LostPosts,
			FoundPosts:    stats.FoundPosts,
		},
		AnnouncementStats: AnnouncementStats{
			TotalAnnouncements: stats.TotalAnnouncements,
		},
	}

	utils.ResponseSuccess(c, resp)
}
