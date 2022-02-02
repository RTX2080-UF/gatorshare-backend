package models

import (
	"github.com/jinzhu/gorm"
	// "encoding/json"
	// "github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	gorm.Model
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Zipcode   uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password" binding:"required"`
	Bookmarks string `json:"bookmarks"`
}

type Post struct {
	gorm.Model	
	UserId       uint      `json:"userId" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	UserLimit    int    `json:"userLimit" binding:"required"`
	Participants string    `json:"participants" gorm:"default:1"`
	Expiry       float32   `json:"expiry" gorm:"default:24"`
	ViewCount    int64     `json:"viewCount" gorm:"default:0"`
	Status       int       `json:"status" binding:"required"`
	Categories   string
	Tags         string
}

type Comment struct {
	gorm.Model
	UserId     User      `json:"userId" binding:"required"`
	PostId     Post      `json:"postId" binding:"required"`
	Message    string    `json:"message" binding:"required"`
	ParentId   uint      `json:"parentId"`
	Votes      int       `json:"votes" gorm:"default:0"`
}