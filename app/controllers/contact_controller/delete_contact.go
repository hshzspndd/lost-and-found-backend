package contact_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteContactRequest struct {
	ContactID int `json:"contact_id"`
}

// 删除联系人
func DeleteContact(c *gin.Context) {
	var dcr DeleteContactRequest
	err := c.ShouldBindJSON(&dcr)
	if err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	err = services.DeleteMyPost(dcr.ContactID, claims.UserID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, "帖子删除成功")
}
