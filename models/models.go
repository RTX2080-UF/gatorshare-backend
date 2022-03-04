package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"userName" gorm:"uniqueIndex"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `gorm:"uniqueIndex"`
	Zipcode   uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
	Bookmarks string `json:"bookmark"`
}

type Post struct {
	gorm.Model
	UserID       uint 	 `json:"userId"`
	User         User
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	UserLimit    uint    `json:"userLimit"`
	Participants uint  	 `json:"participantNum" gorm:"default:1"`
	Expiry       float32 `gorm:"default:24"`
	ViewCount    int64   `json:"viewCount" gorm:"default:0"`
	Status       int     `json:"status"`
	Categories   string  `json:"categories"`
	Tags         string  `json:"tags"`
}

type Comment struct {
	gorm.Model
	UserID   uint    `json:"userId"`
	User	 User
	PostID   uint    `json:"postId"`
	Post	 Post
	Message  string  `json:"message"`
	ParentId uint    `json:"parentId"`
	Votes    int     `json:"votes" gorm:"default:0"`
}

