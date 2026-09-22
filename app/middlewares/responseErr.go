package middlewares

import (
	"lost-and-found-backend/app/errs"

	"github.com/gin-gonic/gin"
)

// 响应格式
type ResponseForm struct {
	HTTPCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func (e *ResponseForm) Error() string {
	return e.Message
}

// 全局异常统一响应
func GlobalResponseError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		if err, ok := c.Errors.Last().Err.(*errs.ResponseErrorForm); ok {
			c.JSON(err.HTTPCode, ResponseForm{
				Code:    err.Code,
				Message: err.Message,
				Data:    nil,
			})
		} else {
			c.JSON(500, ResponseForm{
				Code:    500,
				Message: "服务器内部错误",
				Data:    nil,
			})
		}
		c.Abort()
	}

}
