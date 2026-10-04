package admin_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminResolvePostData struct {
	IsResolve string `json:"is_resolve" binding:"required,oneof=已解决 未解决"`
}

// 管理员改变帖子的解决状态
func AdminResolvePost(c *gin.Context) {
	var data AdminResolvePostData
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

	err = services.AdminResolvePost(postID, data.IsResolve)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)
}
