package claim_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuditClaimData struct {
	Approve bool `json:"approve"`
}

// 作者审批认领申请
func AuditClaim(c *gin.Context) {
	claimID, err := strconv.Atoi(c.Param("claim_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	var data AuditClaimData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	if err := services.AuditClaim(claimID, claims.UserID, data.Approve); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	utils.ResponseSuccess(c, nil)
}
