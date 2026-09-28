package contact_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactListResp struct {
	OwnerID   int    `json:"owner_id"`
	ContactID int    `json:"id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num"`
	Major     string `json:"major"`
	Note      string `json:"note"`
}

type ContactsResp struct {
	List  []ContactListResp `json:"list"`
	Total int               `json:"total"`
	Page  int               `json:"page"`
}

// 获取公开帖子列表
func GetContacts(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	if page < 1 {
		page = 1
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	contactList, total, err := services.GetContacts(page, claims.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	contactListResp := make([]ContactListResp, 0, len(contactList))
	for _, contact := range contactList {
		contactListResp = append(contactListResp, ContactListResp{
			OwnerID:   contact.OwnerID,
			ContactID: contact.ContactID,
			StudentID: contact.StudentID,
			Name:      contact.Name,
			Sex:       contact.Sex,
			PhoneNum:  contact.PhoneNum,
			Major:     contact.Major,
			Note:      contact.Note,
		})
	}

	resp := ContactsResp{
		List:  contactListResp,
		Total: total,
		Page:  page,
	}

	utils.ResponseSuccess(c, resp)
}
