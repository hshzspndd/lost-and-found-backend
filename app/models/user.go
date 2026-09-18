package models

// 用户信息的模型
type User struct {
	UserId   int    `json:"user_id" gorm:"primarykey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
