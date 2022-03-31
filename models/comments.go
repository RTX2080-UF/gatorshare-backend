package models

import "gorm.io/gorm"

func AddNewComment(db *gorm.DB, comments *Comment) (uint, error) {
	err := db.Create(comments).Error
	if err != nil {
		return 0, err
	}
	return comments.ID, nil
}

func GetAllComment(db *gorm.DB, comments *[]Comment, id uint) error {
	res := db.Where("post_id = ?", id).Find(&comments)
	return res.Error
}

func GetOneComment(db *gorm.DB, comments *Comment, id uint) error {
	res := db.First(&comments, id)
	return res.Error
}

func DeleteComment(db *gorm.DB, comments *Comment, id uint) error {
	res := db.Delete(&Comment{}, id)
	return res.Error
}

func UpdateComment(db *gorm.DB, comments *Comment) (error) {
	res := db.Model(&comments).Where("user_id = ? AND post_id = ?", comments.UserID, comments.PostID).Update("message", comments.Message)
	return res.Error
}

// func UpdatecommentParentId(db *gorm.DB, comments *Comment, parentId uint) error {
// 	res := db.Model(&comments).Where("user_id = ? AND post_id = ?", comments.UserID, comments.PostID).Update("parent_id", parentId)
// 	return res.Error
// }

// func replyComment(db *gorm.DB, comments *Comment) (error) {
// 	res, err := AddNewcomment(db, comments)	
// 	if res != 0 {
// 	res1 := UpdatecommentParentId(db, comments, res)
// 	return res1
// 	}
// 	return err
// }