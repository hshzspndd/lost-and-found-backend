package contact_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type ContactData struct {
	StudentID string `json:"student_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex" binding:"required"`
	PhoneNum  string `json:"phone_num" binding:"required"`
	Major     string `json:"major"`
	Note      string `json:"note"`
}

type ContactResp struct {
	OwnerID   int    `json:"owner_id"`
	ContactID int    `json:"contact_id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num"`
	Major     string `json:"major"`
	Note      string `json:"note"`
}

func AddContact(c *gin.Context) {
	var data ContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	contact, err := services.AddContact(claims.UserID, data.StudentID, data.Name, data.Sex, data.PhoneNum, data.Major, data.Note)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ContactResp{
		OwnerID:   contact.OwnerID,
		ContactID: contact.ContactID,
		StudentID: contact.StudentID,
		Name:      contact.Name,
		Sex:       contact.Sex,
		PhoneNum:  contact.PhoneNum,
		Major:     contact.Major,
		Note:      contact.Note,
	}

	utils.ResponseSuccess(c, resp)
}
