package post_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MyPostListResp struct {
	PostID        int       `json:"post_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Contact       string    `json:"contact"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	IsResolve     string    `json:"is_resolve"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type MyPostsResp struct {
	List     []MyPostListResp `json:"list"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}

// 查询自己的帖子
func GetMyPosts(c *gin.Context) {
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

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	postList, total, err := services.GetMyPosts(claims.UserID, page, status, postType, isResolve)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	postListResp := make([]MyPostListResp, 0, len(postList))
	for _, post := range postList {
		postListResp = append(postListResp, MyPostListResp{
			PostID:        post.PostID,
			PostType:      post.PostType,
			Title:         post.Title,
			EventLocation: post.EventLocation,
			EventTime:     post.EventTime,
			Contact:       post.Contact,
			Description:   post.Description,
			ImageUrl:      post.ImageUrl,
			IsResolve:     post.IsResolve,
			Status:        post.Status,
			CreatedAt:     post.CreatedAt,
			UpdatedAt:     post.UpdatedAt,
		})
	}

	resp := MyPostsResp{
		List:     postListResp,
		Total:    total,
		Page:     page,
		PageSize: 15,
	}

	utils.ResponseSuccess(c, resp)
}
