package comment_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteCommentRequest struct {
	CommentID int `json:"comment_id"`
}

// 删除我的评论
func DeleteMyComment(c *gin.Context) {
	var dcr DeleteCommentRequest
	err := c.ShouldBindJSON(&dcr)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	err = services.DeleteMyComment(dcr.CommentID, claims.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "帖子删除成功")
}
