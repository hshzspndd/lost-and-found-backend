package models

import "time"

type Claim struct {
	ClaimID   int       `json:"claim_id" gorm:"primarykey;autoIncrement"`
	PostID    int       `json:"post_id" gorm:"index;not null"`
	ClaimerID int       `json:"claimer_id" gorm:"index;not null"`
	Reason    string    `json:"reason" gorm:"size:255;not null"`
	Status    string    `json:"status" gorm:"default:'待处理'"` // 待处理 / 已同意 / 已拒绝
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
