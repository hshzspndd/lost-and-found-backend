package models

type Contact struct {
	ContactID int    `json:"contact_id" gorm:"primarykey;autoIncrement"`
	OwnerID   int    `json:"owner_id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num"`
	Major     string `json:"major"`
	Note      string `json:"note"`
}
