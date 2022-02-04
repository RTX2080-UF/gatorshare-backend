package models

func AddNewcomment(comments *Comment)(uint,error){
	err:= DB.Create(comments).Error
	if err != nil {
		return 0, err
	}
	return comments.ID, nil
}
func GetAllcomment(comments *[]Comment, id int) (error) {
	res := DB.Where("user_id = ?", id).Find(&comments)
	return res.Error
}

func GetOnecomment(comments *Comment, id int) (error) {
	res := DB.Find(&comments, id)
	return res.Error
}
func Deletecomment(comments *Comment, id int) (error) {
	res := DB.Delete(&Comment{}, id)
	return res.Error
}
func Updatecomment(comments *Comment, id int) (error) {
	res := DB.Update(&Comment{}, id)
	return res.Error
}