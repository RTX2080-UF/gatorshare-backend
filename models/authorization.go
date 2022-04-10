package models

import (
	"gorm.io/gorm"
)

func UpdatePasswordStatus(db *gorm.DB, resetObj ResetPassword) (bool, error) {

	err := db.Create(resetObj).Error
	if err != nil {
		return false, err
	}

	return true, nil

}
