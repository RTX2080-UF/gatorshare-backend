package models

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

func AddNewUser(db *gorm.DB, user *User) (uint, error) {
	err := db.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func GetUserProfile(db *gorm.DB, user *User, id uint) error {
	res := db.First(&user, id)
	return res.Error
}

func DeleteUser(db *gorm.DB, id int) error {
	res := db.Delete(&User{}, id)
	return res.Error
}

func UpdateUserProfile(db *gorm.DB, user *User) (error) {
	res := db.Model(user).Clauses(clause.Returning{}).Updates(user)
	return res.Error
}

func GetUserDetailByUsername(db *gorm.DB, username string) (User, error){
	var user User
	res := db.Omit("User.Password").Where(&User{Username: username}).First(&user)
	return user, res.Error
}

func GetUserDetailByEmail(db *gorm.DB, email string) (User, error){
	var user User
	res := db.Omit("User.Password").Where(&User{Email: email}).First(&user)
	return user, res.Error
}

func CheckUserExists(db *gorm.DB, emailId string) (bool, error) {
	var count int64
	res := db.Model(&User{}).Where("Email = ?", emailId).Count(&count)

	if (count <= 0) {
		return false, res.Error
	}

	return true, res.Error
}

func FollowUserByUser(db *gorm.DB, userId uint, followerId uint) (uint, error) {
	var usersFollower = Follower {
		UserID: followerId,
		FollowerID: userId,
	}

	err := db.Create(&usersFollower).Error
	if err != nil {
		return 0, err
	}

	return usersFollower.ID, nil
}

func GetFollowers(db *gorm.DB, follower *[]Follower, id uint) error {
	res := db.Preload("User").Omit("User.password").Where("user_id=?", id).Find(&follower)
	return res.Error
}

func AddFeedback(db *gorm.DB, feedback *FeedBack) (uint,error){
	var user User
	err := db.Create(feedback).Error
	if err != nil {
		return 0, err
	}
	uid := feedback.UserID
	var count int64
	var sum int64
	db.Model(feedback).Where("user_id=?",uid).Count(&count)
	db.Model(feedback).Select("sum(rating)").Row().Scan(&sum)
	avg := sum / count
	db.Model(user).Find("user_id=?",uid).Updates(User{Rating: uint(avg)})
	return feedback.ID, nil
}

func GetFeedback(db *gorm.DB, feedback *FeedBack, id int) error {
	res := db.Where("user_id =?", id).Find(&feedback)
	return res.Error
}