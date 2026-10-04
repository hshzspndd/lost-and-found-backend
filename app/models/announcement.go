package models

import "time"

type Announcement struct {
	AnnouncementID int       `json:"announcement_id" gorm:"primarykey;autoIncrement"`
	Title          string    `json:"title" gorm:"size:100;not null"`
	Content        string    `json:"content" gorm:"type:text;not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
