package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdatResetPassword(db *gorm.DB, resetObj ResetPassword) (bool, error) {

	err :=db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"unique_rnd_str": resetObj.UniqueRndStr}),
	  }).Create(&resetObj).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdatPassword(db *gorm.DB, userId uint, password string) (error) {
	res := db.Model(User{}).Where("id = ?", userId).Update("password", password)
	if res.Error == nil {
		res = db.Model(ResetPassword{}).Where("user_id = ?", userId).Update("status", false)
	}
	return res.Error
}

func VerifyPasswordReset(db *gorm.DB, userId uint, token string) (bool, error) {
	var count int64
	res := db.Model(&ResetPassword{}).Where("user_id = ? AND unique_rnd_str = ? AND Status = ?", userId, token, true).Count(&count)

	if (count <= 0) {
		return false, res.Error
	}

	return true, res.Error
}
