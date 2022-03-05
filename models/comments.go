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
	res := db.Where("post_id = ?", id).Find(&comments)
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
	res := db.Model(&comments).Where("user_id = ? AND post_id = ?", comments.UserID, comments.PostID).Update("message", comments.Message)
	return res.Error
}