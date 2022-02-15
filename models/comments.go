package models

import "gorm.io/gorm"

func AddNewcomment(db *gorm.DB, comments *Comment) (uint, error) {
	err := db.Create(comments).Error
	if err != nil {
		return 0, err
	}
	return comments.ID, nil
}

func GetAllcomment(db *gorm.DB, comments *[]Comment, id int) error {
	res := db.Where("user_id = ?", id).Find(&comments)
	return res.Error
}

func GetOnecomment(db *gorm.DB, comments *Comment, id int) error {
	res := db.First(&comments, id)
	return res.Error
}

func Deletecomment(db *gorm.DB, comments *Comment, id int) error {
	res := db.Delete(&Comment{}, id)
	return res.Error
}

func Updatecomment(db *gorm.DB, comments *Comment) (error) {
	res := db.Model(&comments).Updates(comments)
	return res.Error
}