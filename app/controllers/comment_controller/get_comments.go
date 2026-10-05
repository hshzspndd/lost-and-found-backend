package comment_controller

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListResp struct {
	CommentID int       `json:"comment_id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentsResp struct {
	List  []CommentListResp `json:"list"`
	Total int               `json:"total"`
	Page  int               `json:"page"`
}

func GetComments(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// 从URL获取参数
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		c.Error(errs.ErrInvalidQuery)
		c.Abort()
		return
	}

	commentList, total, err := services.GetComments(page, postID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	commentListResp := make([]CommentListResp, 0, len(commentList))
	for _, comment := range commentList {
		commentListResp = append(commentListResp, CommentListResp{
			CommentID: comment.CommentID,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
		})
	}

	resp := CommentsResp{
		List:  commentListResp,
		Total: total,
		Page:  page,
	}

	utils.ResponseSuccess(c, resp)
}
