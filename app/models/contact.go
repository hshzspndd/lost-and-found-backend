package models

type Contact struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num"`
	Major     string `json:"major"`
	Blacklist bool   `json:"blacklist"`
}
