package comment_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentData struct {
	Contact string `json:"contact" binding:"required"`
}

type CommentResp struct {
	CommentID int    `json:"comment_id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
}

func CreateComment(c *gin.Context) {
	var data CommentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	// 从URL获取参数
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	comment, err := services.CreateComment(postID, claims.UserID, data.Contact)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := CommentResp{
		CommentID: comment.CommentID,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Content:   comment.Content,
	}

	utils.ResponseSuccess(c, resp)
}
