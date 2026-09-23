package errs

type ResponseErrorForm struct {
	HTTPCode int    `json:"httpcode"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func (e *ResponseErrorForm) Error() string {
	return e.Message
}

// 打包错误
var (
	ErrBindJSON      = &ResponseErrorForm{400, 1001, "数据获取失败"}
	ErrUserExists    = &ResponseErrorForm{409, 1002, "用户名或手机号已存在"}
	ErrUserCheckFail = &ResponseErrorForm{500, 1003, "用户信息校验失败"}
	ErrHashPassword  = &ResponseErrorForm{500, 1004, "密码加密失败"}
	ErrDatabase      = &ResponseErrorForm{500, 1005, "数据库出错"}
	ErrNoPermission  = &ResponseErrorForm{403, 1006, "没有权限"}
	ErrUserNotFound  = &ResponseErrorForm{404, 1007, "用户不存在"}
	ErrWrongPassword = &ResponseErrorForm{403, 1008, "密码错误"}
	ErrGenerateToken = &ResponseErrorForm{500, 1009, "登录令牌生成失败"}
	ErrUnauthorized  = &ResponseErrorForm{401, 1010, "未登录或无效的token"}
	ErrInvalidToken  = &ResponseErrorForm{401, 1011, "token解析失败"}
)
