package models

// 用户信息的模型
type User struct {
	UserId   int    `json:"user_id" gorm:"primarykey;autoIncrement"`
	Username string `json:"username" gorm:"uniqueIndex;size:50;not null"`
	PhoneNum string `json:"phone_num" gorm:"uniqueIndex;size:20"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
