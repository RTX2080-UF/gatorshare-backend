package models

import (
	"gorm.io/gorm"
)

func GetAllPost(db *gorm.DB, posts *[]Post, id uint) error {
	res := db.Where("user_id = ?", id).Find(&posts)
	return res.Error
}

func AddNewPost(db *gorm.DB, posts *Post) (uint, error) {
	err := db.Create(&posts).Error
	if err != nil {
		return 0, err
	}

	return posts.ID, nil
}

func GetOnePost(db *gorm.DB, post *Post, id int) error {
	res := db.Omit("User.Password").Joins("User").First(&post, id)
	return res.Error
}

func DeletePost(db *gorm.DB, post *Post, id int) error {
	res := db.Delete(&Post{}, id)
	return res.Error
}

func UpdatePost(db *gorm.DB, post *Post) error {
	res := db.Model(post).Updates(post)
	return res.Error
}

func ReactToPost(db *gorm.DB, postReaction *UserPost) (uint, error) {
	err := db.Create(postReaction).Error
	if err != nil {
		return 0, err
	}

	return postReaction.ID, nil
}

func GetReactions(db *gorm.DB, postId uint, postReaction *[]UserPost) (error) {
	err := db.Preload("User").Omit("users.password").Where("post_id=?", postId).Find(&postReaction).Error
	return err
}

func SearchPost(db *gorm.DB, posts *[]Post, tagIds []uint) error {
	var err error
	for _ , elem := range tagIds {
		var tp TagPost
		err = db.Preload("Post").Where("tag_id=?",elem).Find(&tp).Error
		if(err == nil){
			*posts = append(*posts, tp.Post)
		}
	}
	return err
}