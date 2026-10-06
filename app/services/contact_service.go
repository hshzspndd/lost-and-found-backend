package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"

	"gorm.io/gorm"
)

// ================================================== 新增联系人 ==================================================

func AddContact(ownerID int, studentID, name, sex, phoneNum, major, note string) (*models.Contact, error) {
	var contact = models.Contact{
		OwnerID:   ownerID,
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

func GetContacts(page, ownerID, pageSize int) ([]models.Contact, int, error) {
	var total int64

	contacts := make([]models.Contact, 0)

	query := database.DB.Model(&models.Contact{}).Where("owner_id = ?", ownerID)

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

// ================================================== 删除联系人 ==================================================

// 删除自己的联系人
func DeleteMyContact(ContactID, OwnerID int) error {
	res := database.DB.Model(&models.Contact{}).Where("contact_id = ? AND owner_id = ?", ContactID, OwnerID).Delete(&models.Contact{})

	if res.Error != nil {
		return errs.ErrDatabase
	}

	if res.RowsAffected == 0 {
		contactExists, err := CheckContactExistByContactID(ContactID)
		if err != nil {
			return errs.ErrDatabase
		}
		if !contactExists {
			return errs.ErrContactNotFound
		}
		return errs.ErrIsNotYourContact
	}
	return nil
}

func CheckContactExistByContactID(contactID int) (bool, error) {
	var contact models.Contact
	err := database.DB.Model(&models.Contact{}).Where("contact_id = ?", contactID).First(&contact).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
