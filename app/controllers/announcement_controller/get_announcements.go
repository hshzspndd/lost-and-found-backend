package announcement_controller

import (
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AnnouncementListResp struct {
	AnnouncementID int       `json:"announcement_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

type AnnouncementsResp struct {
	List     []AnnouncementListResp `json:"list"`
	Total    int                    `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
}

// 获取公告列表
func GetAnnouncements(c *gin.Context) {
	var pageSize = 15

	pageStr := c.DefaultQuery("page", "1")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	announcements, total, err := services.GetAnnouncements(page, pageSize)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	listResp := make([]AnnouncementListResp, 0, len(announcements))
	for _, a := range announcements {
		listResp = append(listResp, AnnouncementListResp{
			AnnouncementID: a.AnnouncementID,
			Title:          a.Title,
			Content:        a.Content,
			CreatedAt:      a.CreatedAt,
		})
	}

	resp := AnnouncementsResp{
		List:     listResp,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}

	utils.ResponseSuccess(c, resp)
}
