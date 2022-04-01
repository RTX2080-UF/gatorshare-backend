package models

import (
	"encoding/json"
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

func UpdateTag(db *gorm.DB, tag *Tag) error {
	res := db.Model(tag).Clauses(clause.Returning{}).Updates(tag)
	return res.Error
}

func FollowTagsByUser(db *gorm.DB, userId uint, tagId uint) (uint, error) {
	var usersTags = TagUser{
		UserID: userId,
		TagID:  tagId,
	}

	fmt.Print(usersTags)
	err := db.Create(&usersTags).Error
	if err != nil {
		return 0, err
	}

	return usersTags.ID, nil
}

func PopularTags(db *gorm.DB, tags *[]Tag, countTags int) error {
	res := db.Limit(countTags).Order("votes desc").Find(&tags)
	return res.Error
}

func CheckTagsExist(db *gorm.DB, tags []uint) bool {
	for i:=0;i<len(tags);i++{
		res := db.First(&Tag{},tags[i])
		if( res.Error != nil){
			return false
		}
	}
	return true
}

func AddUserTags(db *gorm.DB, uid uint, tags []uint) ([]uint,error){
	// type tagUser []TagUser
	var response []uint
	for i:=0;i<len(tags);i++ { 
		var obj = TagUser{UserID : uid , TagID : tags[i]}
		err := db.Create(&obj).Error
		if err != nil {
			response[i] = 0
		}else{
			response[i] = obj.ID
		}		
		
	}
	return response, nil
	// res := db.CreateInBatches()
}