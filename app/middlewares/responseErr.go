package middlewares

import "github.com/gin-gonic/gin"

// 响应格式
type ResponseForm struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ResponseForm) Error() string {
	return e.Message
}

func GetError(code int, message string) *ResponseForm {
	return &ResponseForm{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 全局异常统一响应
func GlobalResponseError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		if err, ok := c.Errors.Last().Err.(*ResponseForm); ok {
			c.JSON(err.Code, ResponseForm{
				Code:    err.Code,
				Message: err.Message,
				Data:    err.Data,
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
