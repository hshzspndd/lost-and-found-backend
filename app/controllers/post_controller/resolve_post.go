package post_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResolveStatus struct {
	IsResolve string `json:"is_resolve" binding:"required,oneof=已解决 未解决"`
}

// 帖子的解决与撤销解决
func Resolve(c *gin.Context) {
	var resolveStatus ResolveStatus
	err := c.ShouldBindJSON(&resolveStatus)
	if err != nil {
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

	err = services.Resolve(postID, claims.UserID, resolveStatus.IsResolve)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)
}
