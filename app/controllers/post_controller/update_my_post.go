package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMyPostData struct {
	Title         string `json:"title" binding:"required,max=20"`
	EventLocation string `json:"event_location" binding:"required,max=20"`
	EventTime     string `json:"event_time" binding:"required,max=20"`
	Contact       string `json:"contact" binding:"required,max=20"`
	Description   string `json:"description" binding:"max=150"`
	ImageUrl      string `json:"image_url"`
}

// 编辑并重新提交帖子
func UpdateMyPost(c *gin.Context) {
	var data UpdateMyPostData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	err = services.UpdateMyPost(postID, claims.UserID, data.Title, data.EventLocation, data.EventTime, data.Contact, data.Description, data.ImageUrl)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)
}
