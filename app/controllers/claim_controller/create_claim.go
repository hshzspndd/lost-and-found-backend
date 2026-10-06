package claim_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateClaimData struct {
	Reason string `json:"reason" binding:"required,min=5,max=255"`
}

type ClaimResp struct {
	ClaimID   int    `json:"claim_id"`
	PostID    int    `json:"post_id"`
	ClaimerID int    `json:"claimer_id"`
	Reason    string `json:"reason"`
	Status    string `json:"status"`
}

func CreateClaim(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	var data CreateClaimData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(errs.ErrBindJSON)
		c.Abort()
		return
	}

	val, _ := c.Get("claims")
	claims := val.(*utils.Claims)

	claim, err := services.CreateClaim(postID, claims.UserID, data.Reason)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	resp := ClaimResp{
		ClaimID:   claim.ClaimID,
		PostID:    claim.PostID,
		ClaimerID: claim.ClaimerID,
		Reason:    claim.Reason,
		Status:    claim.Status,
	}
	utils.ResponseSuccess(c, resp)
}
