package claim_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPostClaimResp struct {
	List     []services.PostClaimWithUser `json:"list"`
	Total    int                          `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"page_size"`
}

// 获取自己帖子收到的认领申请
func GetPostClaims(c *gin.Context) {
	var pageSize = 15
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	list, total, err := services.GetPostClaims(postID, claims.UserID, page, pageSize)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := GetPostClaimResp{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	utils.ResponseSuccess(c, resp)
}
