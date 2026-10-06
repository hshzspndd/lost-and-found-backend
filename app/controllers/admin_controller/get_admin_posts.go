package admin_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminPostListResp struct {
	PostID        int       `json:"post_id"`
	UserID        int       `json:"user_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	IsResolve     string    `json:"is_resolve"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AdminPostsResp struct {
	List     []AdminPostListResp `json:"list"`
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

// 管理员获取所有帖子
func GetAdminPosts(c *gin.Context) {
	var pageSize = 15

	postType := c.Query("post_type")
	status := c.Query("status")
	isResolve := c.Query("is_resolve")

	pageStr := c.DefaultQuery("page", "1")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	postList, total, err := services.AdminGetAllPosts(page, postType, status, isResolve, pageSize)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	postListResp := make([]AdminPostListResp, 0, len(postList))
	for _, post := range postList {
		postListResp = append(postListResp, AdminPostListResp{
			PostID:        post.PostID,
			UserID:        post.UserID,
			PostType:      post.PostType,
			Title:         post.Title,
			EventLocation: post.EventLocation,
			EventTime:     post.EventTime,
			Description:   utils.SubStr(post.Description, 15),
			ImageUrl:      post.ImageUrl,
			IsResolve:     post.IsResolve,
			Status:        post.Status,
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		})
	}

	resp := AdminPostsResp{
		List:     postListResp,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	utils.ResponseSuccess(c, resp)
}
