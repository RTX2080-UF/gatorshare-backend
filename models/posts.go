package models

import (
	"gorm.io/gorm"
)

func GetAllpost(db *gorm.DB, posts *[]Post, id int) error {
	res := db.Where("user_id = ?", id).Find(&posts)
	return res.Error
}

func AddNewpost(db *gorm.DB, posts *Post) (uint, error) {
	err := db.Create(posts).Error
	if err != nil {
		return 0, err
	}

	return posts.ID, nil
}

func GetOnepost(db *gorm.DB, post *Post, id int) error {
	res := db.Omit("User.Password").Joins("User").First(&post, id)
	return res.Error
}

func Deletepost(db *gorm.DB, post *Post, id int) error {
	res := db.Delete(&Post{}, id)
	return res.Error
}

func UpdatePost(db *gorm.DB, post *Post) error {
	res := db.Model(post).Updates(post)
	return res.Error
}
