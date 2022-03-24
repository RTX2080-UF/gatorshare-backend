package controllers

import "gatorshare/models"

type UserProfile struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Zipcode	  uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
}

type UpdateUserProfile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	OldPassword  string `json:"oldPassword"`
	Password  string `json:"password" validate:"min=8,max=40,regexp=^(?=.*[0-9])(?=.*[a-z]).{8,32}$"`
	Zipcode	  uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
}

type User struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Zipcode   uint   `json:"zipcode"`
	Avatar    string `json:"avatar"`
	Bookmarks string `json:"bookmarks"`
}

type Post struct {
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description"`
	Participants uint  	 `json:"participants"`
	Expiry       float32 `json:"expiry"`
	ViewCount    int64   `json:"viewCount"`
	UserLimit    uint    `json:"userLimit" binding:"required"`
	Status       int     `json:"status" binding:"required"`
	Categories   string
	Tags         string
}

type Comment struct {
	UserID   uint   `json:"userId" binding:"required"`
	PostID   uint   `json:"postId" binding:"required"`
	Message  string `json:"message" binding:"required"`
	ParentId uint   `json:"parentId"`
	Votes    int    `json:"votes"`
}

type Tag struct {
	Name      string `json:"tagName" binding:"required"`
	Frequency int    `json:"count"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostRequestToDBModel(req Post, UserID uint) models.Post {
	return models.Post {
		UserID:       UserID,  
		Title:        req.Title,
		Description:  req.Description,  		
		UserLimit:    req.UserLimit,
		Participants: req.Participants,
		Expiry:       req.Expiry,
		ViewCount:    req.ViewCount,
		Status:       req.Status,
		Categories:   req.Categories,
	}
}

func CommentRequestToDBModel(req Comment, UserID uint) models.Comment {
	return models.Comment {
		UserID:   UserID,
		PostID:   req.PostID,
		Message:  req.Message,
		ParentId: req.ParentId,
		Votes:    req.Votes,
	}
}

func UserRequestToDBModel(req UserProfile) models.User {
	return models.User {
		Username: req.Username,
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Email: req.Email,
		Password: req.Password,
		Zipcode: req.Zipcode,
	}
}