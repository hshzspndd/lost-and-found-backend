package services

type ResponseErrorForm struct {
	Code    int
	Message string
}

func (e *ResponseErrorForm) Error() string {
	return e.Message
}

// 打包错误
var (
	ErrHashPassword  = &ResponseErrorForm{500, "密码加密失败"}
	ErrUserCheckFail = &ResponseErrorForm{500, "用户信息校验失败"}
	ErrWrongPassword = &ResponseErrorForm{403, "密码错误"}
	ErrUserExists    = &ResponseErrorForm{409, "用户已存在"}
	ErrDatabase      = &ResponseErrorForm{500, "数据库出错"}
)
