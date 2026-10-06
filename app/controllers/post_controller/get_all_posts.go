package post_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PostListResp struct {
	PostID        int       `json:"post_id"`
	UserID        int       `json:"user_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	IsResolve     string    `json:"is_resolve"`
	CreatedAt     time.Time `json:"created_at"`
}

type PostsResp struct {
	List     []PostListResp `json:"list"`
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

// 获取公开帖子列表
func GetAllPosts(c *gin.Context) {
	var pageSize = 15
	postType := c.Query("post_type")

	isResolve := c.DefaultQuery("is_resolve", "未解决")
	if isResolve != "未解决" && isResolve != "已解决" {
		isResolve = "未解决"
	}

	pageStr := c.DefaultQuery("page", "1")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	postList, total, err := services.GetAllPosts(page, postType, isResolve, pageSize)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	postListResp := make([]PostListResp, 0, len(postList))
	for _, post := range postList {
		postListResp = append(postListResp, PostListResp{
			PostID:        post.PostID,
			UserID:        post.UserID,
			PostType:      post.PostType,
			Title:         post.Title,
			EventLocation: post.EventLocation,
			EventTime:     post.EventTime,
			Description:   utils.SubStr(post.Description, 15),
			ImageUrl:      post.ImageUrl,
			IsResolve:     post.IsResolve,
			CreatedAt:     post.CreatedAt,
		})
	}

	resp := PostsResp{
		List:     postListResp,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	utils.ResponseSuccess(c, resp)
}
