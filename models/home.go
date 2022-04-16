package models

import (
	"gorm.io/gorm"
)

func GetUserHomePosts(db *gorm.DB, id uint) ([]TagPost, error) {
	var userLikedTags []TagUser
	res := db.Joins("tags").Where("user_id = ?", id).Find(&userLikedTags);

	var tagIDs []uint
	for  _, entity := range userLikedTags {
		tagIDs = append(tagIDs, entity.ID)
	}

	var tagPosts []TagPost
	res = db.Joins("posts").Order("created_at desc, updated_at desc").Where("id = ANY(?)", tagIDs).Find(&tagPosts)
	return tagPosts, res.Error
}