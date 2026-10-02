package models

type Contact struct {
	OwnerID   int    `json:"owner_id" gorm:"primarykey;autoIncrement"`
	ContactID int    `json:"id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num"`
	Major     string `json:"major"`
	Note      string `json:"note"`
}
