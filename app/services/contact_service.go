package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
)

// ================================================== 新增联系人 ==================================================

func AddContact(ownerID, contactID int, studentID, name, sex, phoneNum, major, note string) (*models.Contact, error) {
	var contact = models.Contact{
		OwnerID:   ownerID,
		ContactID: contactID,
		StudentID: studentID,
		Name:      name,
		Sex:       sex,
		PhoneNum:  phoneNum,
		Major:     major,
		Note:      note,
	}
	err := database.DB.Model(&models.Contact{}).Create(&contact).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &contact, nil

}

// ================================================== 获取联系人列表 ==================================================

func GetContacts(page, OwnerID int) ([]models.Contact, int, error) {
	var total int64
	var pageSize int = 15

	contacts := make([]models.Contact, 0)

	query := database.DB.Model(&models.Contact{}).Where("OwnerID = ?", OwnerID)

	//获取联系人总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Offset(offset).Limit(pageSize).Find(&contacts).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return contacts, int(total), nil
}
