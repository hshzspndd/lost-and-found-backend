package models

import "time"

type Post struct {
	PostID        int       `json:"post_id" gorm:"primarykey;autoIncrement"`
	UserID        int       `json:"user_id"`
	PostType      string    `json:"post_type"`
	Title         string    `json:"title"`
	EventLocation string    `json:"event_location"`
	EventTime     string    `json:"event_time"`
	Contact       string    `json:"contact"`
	Description   string    `json:"description"`
	ImageUrl      string    `json:"image_url"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
