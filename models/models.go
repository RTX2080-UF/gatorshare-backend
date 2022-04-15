package models

import (
	"gorm.io/gorm"
)
type ReactionType string

const (
    INTERESTED  	ReactionType = "INTERESTED"
    MAYBE 			ReactionType = "MAYBE"
    NOTINTERESTED 	ReactionType = "NOTINTERESTED"
)

type User struct {
	gorm.Model
	Username  string `json:"userName" gorm:"uniqueIndex"`
	Firstname string `json:"firstName" gorm:"not null"`
	Lastname  string `json:"lastName" gorm:"not null"`
	Email     string `gorm:"uniqueIndex" gorm:"not null"`
	Zipcode   uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password" gorm:"not null"`
	Bookmarks string `json:"bookmark"`
}

type ResetPassword struct {
	UserID       uint 	 `json:"userId" gorm:"uniqueIndex"`
	User         User
	Status       bool	 `gorm:"default:false"`
	UniqueRndStr string  `json:"uniqueRndStr" gorm:"not null"`
}

type Post struct {
	gorm.Model
	UserID       uint 	 `json:"userId" gorm:"not null"`
	User         User
	Title        string  `json:"title" gorm:"not null"`
	Description  string  `json:"description"`
	UserLimit    uint    `json:"userLimit" gorm:"default:2"`
	Participants uint  	 `json:"participantNum" gorm:"default:1"`
	Expiry       float32 `gorm:"default:24"`
	ViewCount    int64   `json:"viewCount" gorm:"default:0"`
	Status       int     `json:"status" gorm:"not null"`
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

type Tag struct {
	gorm.Model
	Name        string  `json:"name" gorm:"uniqueIndex"`
	CreatorId   uint    `json:"creatorId"`
	Creator		User
	Votes       int     `json:"votes" gorm:"default:0"`
	Description string  `json:"description"`
}

type Notification struct {
	gorm.Model
	UserID      uint    `json:"userId" gorm:"primaryKey"`
	User	    User
	Link	    string  `json:"link"`
	ReadStatus  bool	`json:"seen" gorm:"default:false"`
	Description string  `json:"description"`
}

type TagUser struct {
	gorm.Model
	UserID   uint    `json:"userId" gorm:"primaryKey"`
	User	 User
	TagID    uint    `json:"tagId" gorm:"primaryKey"`
	Tag	 	 Tag
}

type TagPost struct {
	gorm.Model
	PostID   uint    `json:"postId" gorm:"primaryKey"`
	Post	 Post
	TagID    uint    `json:"tagId" gorm:"primaryKey"`
	Tag	 	 Tag
}

type UserPost struct {
	gorm.Model
	PostID   uint    `json:"postId"`
	Post	 Post
	UserID   uint    `json:"tagId"`
	User	 User
	Reaction ReactionType
}

type Follower struct {
	gorm.Model
	UserID       uint    `json:"userId" gorm:"primaryKey"`
	User	     User
	FollowerID   uint    `json:"followerId" gorm:"primaryKey"`
	Follower	 User
}
