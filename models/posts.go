package models

func GetAllpost(posts *[]Post, id int) (error) {
	res := DB.Where("user_id = ?", id).Find(&posts)
	return res.Error
}

func AddNewpost(posts *Post) (uint, error) {
	err := DB.Create(posts).Error
	if err != nil {
	  	return 0, err
	}
	return posts.ID, nil
}

func GetOnepost(post *Post, id int) (error) {
	res := DB.Find(&post, id)
	return res.Error
}

func Deletepost(post *Post, id int) (error) {
	res := DB.Delete(&Post{}, id)
	return res.Error
}

func UpdatePost(post *Post, id int) (error) {
	res := DB.Update(&Post{}, id)
	return res.Error
}

