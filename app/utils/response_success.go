package utils

import "github.com/gin-gonic/gin"

type ResponseSuccessForm struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 将成功响应打包成函数
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, ResponseSuccessForm{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}
