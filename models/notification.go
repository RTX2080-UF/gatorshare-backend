package models

import (
	"gorm.io/gorm"
)

func AddNotification(db *gorm.DB, notification *Notification) (uint, error) {
	err := db.Create(notification).Error
	if err != nil {
		return 0, err
	}

	return notification.ID, nil
}

func GetNewNotifications(db *gorm.DB, notification *[]Notification, id uint) error {
	res := db.Where("user_id = ? AND read_status = ?", id, false).Find(&notification)
	return res.Error
}

func UpdateNotifications(db *gorm.DB, id uint) error {
	res := db.Model(&Notification{}).Where("user_id = ? AND read_status = ?", id, false).Updates( Notification { ReadStatus: true })
	return res.Error
}