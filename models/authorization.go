package models

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

func UpdatePasswordStatus(db *gorm.DB, resetObj ResetPassword) (bool, error) {

	err :=db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"unique_rnd_str": resetObj.UniqueRndStr}),
	  }).Create(&resetObj).Error
	if err != nil {
		return false, err
	}

	return true, nil

}

func VerifyPasswordReset(db *gorm.DB, emailId string, token string) (bool, error) {
	var count int64
	res := db.Model(&ResetPassword{}).Where("Email = ? AND UniqueRndStr = ? AND Status = ?", emailId, token, true).Count(&count)

	if (count <= 0) {
		return false, res.Error
	}

	return true, res.Error
}
