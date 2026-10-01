package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id" gorm:"primarykey;autoIncrement"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
