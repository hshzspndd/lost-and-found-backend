package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type DeletePostRequest struct {
	PostID int `json:"post_id"`
}

// 删除我的帖子
func DeleteMyPost(c *gin.Context) {
	var dpr DeletePostRequest
	err := c.ShouldBindJSON(&dpr)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	err = services.DeleteMyPost(dpr.PostID, claims.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "帖子删除成功")
}
