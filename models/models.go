package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string 
	Firstname string 
	Lastname  string 
	Email     string 
	Zipcode   uint   
	Avatar    string 
	Password  string 
	Bookmarks string 
}

type Post struct {
	gorm.Model
	UserID       int 	
	User         User
	Title        string  
	Description  string  
	UserLimit    int     
	Participants string  `gorm:"default:1"`
	Expiry       float32 `gorm:"default:24"`
	ViewCount    int64   `gorm:"default:0"`
	Status       int    
	Categories   string
	Tags         string
}

type Comment struct {
	gorm.Model
	UserID   uint   
	User	 User
	PostID   uint   
	Post	 Post
	Message  string 
	ParentId uint   
	Votes    int    `gorm:"default:0"`
}

