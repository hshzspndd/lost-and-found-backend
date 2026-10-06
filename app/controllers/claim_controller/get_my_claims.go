package claim_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMyClaimResp struct {
	List     []services.MyClaimWithPost `json:"list"`
	Total    int                        `json:"total"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
}

// 获取我提交的认领申请
func GetMyClaims(c *gin.Context) {
	var pageSize = 15

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

	list, total, err := services.GetMyClaims(claims.UserID, page, pageSize)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := GetMyClaimResp{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	utils.ResponseSuccess(c, resp)
}
