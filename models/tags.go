package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddNewTag(db *gorm.DB, tag *Tag) (uint, error) {
	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(tag).Error
	if err != nil {
		return 0, err
	}

	return tag.ID, nil
}

func InsertTags(db *gorm.DB, tags[] Tag) ([] uint, error) {
	
	var tagsId []uint
	var err error
	for _, tag := range tags {
		err = db.Create(&tag).Error
		if (err != nil) {
			err = nil
			err = db.Clauses(
				clause.OnConflict{
					Columns:   []clause.Column{{Name: "name"}},
					DoUpdates: clause.Assignments(map[string]interface{}{"votes": tag.Votes + 1}),
				}).Create(&tag).Error
		}
		tagsId = append(tagsId, tag.ID)	
	}

	return tagsId, err
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

func GetPopularTags(db *gorm.DB, tags *[]Tag, countTags int) error {
	res := db.Limit(countTags).Order("votes desc").Find(&tags)
	return res.Error
}

func CheckTagsExist(db *gorm.DB, tags []uint) []uint {
	var verifiedTagIds []uint
	for i:=0; i<len(tags); i++ {
		res := db.First(&Tag{},tags[i])
		if (res.Error == nil) {
			verifiedTagIds = append(verifiedTagIds, tags[i])
		} else {
			log.Println("Tag with id %i doesn't exist", tags[i])
		}
	}
	return verifiedTagIds
}

func AddUserTags(db *gorm.DB, uid uint, tags []uint) error{
	var inputObj []TagUser
	for i:=0; i<len(tags); i++ { 
		var obj = TagUser{UserID : uid , TagID : tags[i]}
		inputObj = 	append(inputObj, obj)	
	}
	
	err := db.Create(&inputObj).Error
	return err
} 

func AddPostTags(db *gorm.DB, pid uint, tags []uint) (error){
	var tagsPost []TagPost
	for i:=0; i<len(tags); i++ { 
		var tagPostObj = TagPost{PostID : pid , TagID : tags[i]}
		tagsPost = 	append(tagsPost, tagPostObj)	
	}

	err := db.Create(&tagsPost).Error
	return err
}

func GetUserLikedTags(db *gorm.DB, uid uint)([]Tag, error) {
	var usertags []TagUser
	err := db.Preload("Tag").Omit("users.password").Where("user_id=?", uid).Find(&usertags).Error

	var tagsArr []Tag
	for _, elem := range usertags {
		tagsArr = append(tagsArr, elem.Tag)
	}

	return tagsArr, err
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

func SearchTagIdHelper(db *gorm.DB, tagNames []string) ([]uint){
	var tagIds []uint
	for _, elem := range tagNames {
		var tag Tag
		err := db.Where("name=?",elem).Find(&tag).Error
		if err == nil {
			tagIds = append(tagIds, tag.ID)
		}
	}
	return tagIds
}