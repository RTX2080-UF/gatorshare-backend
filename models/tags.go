package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func FollowTagsByUser(db *gorm.DB, userId uint, tagId uint) (uint, error) {
	var usersTags = TagUser {
		UserID: userId,
		TagID: tagId,
	}

	fmt.Print(usersTags)
	err := db.Create(&usersTags).Error
	if err != nil {
		return 0, err
	}

	return usersTags.ID, nil
}