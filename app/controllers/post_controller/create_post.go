package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	PostType      string `json:"post_type" binding:"required,oneof=寻物 招领"`
	Title         string `json:"title" binding:"required"`
	EventLocation string `json:"event_location" binding:"required"`
	EventTime     string `json:"event_time" binding:"required"`
	Contact       string `json:"contact" binding:"required"`
	Description   string `json:"description" binding:"max=150"`
	ImageUrl      string `json:"image_url"`
}

type PostResp struct {
	PostID        int    `json:"post_id"`
	UserID        int    `json:"user_id"`
	PostType      string `json:"post_type"`
	Title         string `json:"title"`
	EventLocation string `json:"event_location"`
	EventTime     string `json:"event_time"`
	Contact       string `json:"contact"`
	Description   string `json:"description"`
	ImageUrl      string `json:"image_url"`
	Status        string `json:"status"`
}

func CreatePost(c *gin.Context) {
	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	post, err := services.CreatePost(claims.UserID, data.PostType, data.Title, data.EventLocation, data.EventTime, data.Contact, data.Description, data.ImageUrl)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := PostResp{
		PostID:        post.PostID,
		UserID:        post.UserID,
		PostType:      post.PostType,
		Title:         post.Title,
		EventLocation: post.EventLocation,
		EventTime:     post.EventTime,
		Contact:       post.Contact,
		Description:   post.Description,
		ImageUrl:      post.ImageUrl,
		Status:        post.Status,
	}

	utils.ResponseSuccess(c, resp)
}
