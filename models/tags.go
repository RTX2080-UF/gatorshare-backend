package models

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

func AddNewTag(db *gorm.DB, tag *Tag) (uint, error) {
	err := db.Create(tag).Error
	if err != nil {
		return 0, err
	}

	return tag.ID, nil
}

func GetTag(db *gorm.DB, tag *Tag, id uint) error {
	res := db.First(&tag, id)
	return res.Error
}

func DeleteTag(db *gorm.DB, id uint) error {
	res := db.Delete(&Tag{}, id)
	return res.Error
}

func UpdateTag(db *gorm.DB, tag *Tag) (error) {
	res := db.Model(tag).Clauses(clause.Returning{}).Updates(tag)
	return res.Error
}
