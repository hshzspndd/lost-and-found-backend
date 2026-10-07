package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PostDetailsResp struct {
	PostID        int       `json:"post_id"`
	UserID        int       `json:"user_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Contact       string    `json:"contact"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	IsResolve     string    `json:"is_resolve"`
	CreatedAt     time.Time `json:"created_at"`
}

// 获取帖子详情
func GetPostDetails(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	// 详情接口已公开：未登录时无 claims，userID 默认为 0，只能查看已通过帖子
	var userID int
	if val, exists := c.Get("claims"); exists {
		claims := val.(*utils.Claims)
		userID = claims.UserID
	}

	post, err := services.GetPostDetails(postID, userID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := PostDetailsResp{
		PostID:        post.PostID,
		UserID:        post.UserID,
		PostType:      post.PostType,
		Title:         post.Title,
		EventLocation: post.EventLocation,
		EventTime:     post.EventTime,
		Contact:       post.Contact,
		Description:   post.Description,
		ImageUrl:      post.ImageUrl,
		IsResolve:     post.IsResolve,
		CreatedAt:     post.CreatedAt,
	}

	utils.ResponseSuccess(c, resp)
}
