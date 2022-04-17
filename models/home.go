package models

import (
	"log"

	"gorm.io/gorm"
)

func GetUserHomePosts(db *gorm.DB, id uint) ([]TagPost, error) {
	var userLikedTags []TagUser
	res := db.Where("user_id = ?", id).Find(&userLikedTags);

	var tagIDs []uint
	for  _, entity := range userLikedTags {
		tagIDs = append(tagIDs, entity.TagID)
	}
	
	var tagPosts []TagPost
	res = db.Preload("Post").Order("created_at desc, updated_at desc").Where("tag_id IN(?)", tagIDs).Find(&tagPosts)
	return tagPosts, res.Error
}