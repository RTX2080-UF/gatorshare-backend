package models

import (
	"gorm.io/gorm"
)

func AddNewUser(db *gorm.DB, user *User) (uint, error) {
	err := db.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func GetUserProfile(db *gorm.DB, user *User, id int) error {
	res := db.First(&user, id)
	return res.Error
}

func DeleteUser(db *gorm.DB, id int) error {
	res := db.Delete(&User{}, id)
	return res.Error
}

func UpdateUserProfile(db *gorm.DB, user *User) (error) {
	res := db.Model(user).Updates(user)
	return res.Error
}

func AuthenticateUser(db *gorm.DB, username string,  password string) (uint, error){
	var user User
	err := db.Select("ID").Where(&User{Username: username, Password: password}).First(&user).Error
	return user.ID, err
}