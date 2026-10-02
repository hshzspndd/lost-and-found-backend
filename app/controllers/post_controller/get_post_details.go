package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type PostDetailsRequest struct {
	PostID int `json:"post_id"`
}

type PostDetailsResp struct {
	PostID        int       `json:"post_id" gorm:"primarykey;autoIncrement"`
	UserID        int       `json:"user_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Contact       string    `json:"contact"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
}

// 获取帖子详情
func GetPostDetails(c *gin.Context) {
	var gpd PostDetailsRequest
	err := c.ShouldBindJSON(&gpd)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	post, err := services.GetPostDetails(gpd.PostID, claims.UserID, claims.Role)
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
		CreatedAt:     post.CreatedAt,
	}

	utils.ResponseSuccess(c, resp)
}
