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
	//================================== 关于用户 ==================================

	ErrBindJSON         = &ResponseErrorForm{400, 1001, "数据获取失败"}
	ErrUserExists       = &ResponseErrorForm{409, 1002, "用户名或手机号已存在"}
	ErrUserCheckFail    = &ResponseErrorForm{500, 1003, "用户信息查询失败"}
	ErrHashPassword     = &ResponseErrorForm{500, 1004, "密码加密失败"}
	ErrDatabase         = &ResponseErrorForm{500, 1005, "数据库出错"}
	ErrNoPermission     = &ResponseErrorForm{403, 1006, "没有权限"}
	ErrUserNotFound     = &ResponseErrorForm{404, 1007, "用户不存在"}
	ErrWrongPassword    = &ResponseErrorForm{401, 1008, "密码错误"}
	ErrGenerateToken    = &ResponseErrorForm{500, 1009, "登录令牌生成失败"}
	ErrUnauthorized     = &ResponseErrorForm{401, 1010, "未登录或无效的token"}
	ErrInvalidToken     = &ResponseErrorForm{401, 1011, "登录过期"}
	ErrPhoneFormat      = &ResponseErrorForm{400, 1012, "手机号格式不正确"}
	ErrSamePassword     = &ResponseErrorForm{400, 1013, "新旧密码相同"}
	ErrWrongOldPassword = &ResponseErrorForm{400, 1014, "旧密码错误"}
	ErrNameTooLong      = &ResponseErrorForm{400, 1015, "用户名限制长度为50字符"}
	ErrLoginLocked      = &ResponseErrorForm{429, 1016, "登录失败次数过多，账号已锁定，请15分钟后再试"}

	//================================== 关于帖子 ==================================

	ErrFileNotUploaded = &ResponseErrorForm{400, 2001, "未接收到上传文件"}
	ErrUploadFailed    = &ResponseErrorForm{400, 2002, "文件上传失败"}
	ErrInvalidFileType = &ResponseErrorForm{400, 2003, "文件格式不正确，仅支持 jpg、jpeg、png"}
	ErrFileTooLarge    = &ResponseErrorForm{400, 2004, "图片大小不能超过5MB"}
	ErrInvalidQuery    = &ResponseErrorForm{400, 2005, "路径参数格式错误"}
	ErrPostNotFound    = &ResponseErrorForm{404, 2006, "帖子不存在"}
	ErrIsNotYourPost   = &ResponseErrorForm{403, 2007, "这不是你的帖子"}
	ErrStatusInvalid   = &ResponseErrorForm{409, 2008, "当前状态不允许此操作"}
	ErrPostNotEditable = &ResponseErrorForm{409, 2009, "当前状态不允许编辑"}

	//================================== 关于公告 ==================================

	ErrAnnouncementNotFound = &ResponseErrorForm{404, 3001, "公告不存在"}

	//================================== 关于认领 ==================================

	ErrPostAlreadyResolved = &ResponseErrorForm{409, 4001, "该物品已被认领/解决"}
	ErrCannotClaimOwnPost  = &ResponseErrorForm{403, 4002, "不能认领自己发布的帖子"}
	ErrClaimAlreadyExists  = &ResponseErrorForm{409, 4003, "您已提交过认领申请"}
	ErrClaimNotFound       = &ResponseErrorForm{404, 4004, "认领申请不存在"}
)
