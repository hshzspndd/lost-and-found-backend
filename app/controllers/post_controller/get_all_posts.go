package post_controller

import (
	"lost-and-found-backend/app/errs"
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
	CreatedAt     time.Time `json:"created_at"`
}

type PostsResp struct {
	List  []PostListResp `json:"list"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
}

// 获取公开帖子列表
func GetAllPosts(c *gin.Context) {

	postType := c.Query("post_type")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	if page < 1 {
		page = 1
	}

	postList, total, err := services.GetAllPosts(page, postType)
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
			Description:   post.Description,
			ImageUrl:      post.ImageUrl,
			CreatedAt:     post.CreatedAt,
		})
	}

	resp := PostsResp{
		List:  postListResp,
		Total: total,
		Page:  page,
	}

	utils.ResponseSuccess(c, resp)
}
