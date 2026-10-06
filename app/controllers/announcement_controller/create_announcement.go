package announcement_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAnnouncementData struct {
	Title   string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required"`
}

type CreateAnnouncementResp struct {
	AnnouncementID int       `json:"announcement_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

// 发布公告
func CreateAnnouncement(c *gin.Context) {
	var announcementData CreateAnnouncementData
	err := c.ShouldBindJSON(&announcementData)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	announcement, err := services.CreateAnnouncement(announcementData.Title, announcementData.Content)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := CreateAnnouncementResp{
		AnnouncementID: announcement.AnnouncementID,
		Title:          announcement.Title,
		Content:        announcement.Content,
		CreatedAt:      announcement.CreatedAt,
	}

	utils.ResponseSuccess(c, resp)
}
