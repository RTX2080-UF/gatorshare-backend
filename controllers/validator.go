package controllers

import "gatorshare/models"

type UserRegister struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`

}

type UserProfile struct {
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
	UserID       int     `json:"userId" binding:"required"`
	Title        string  `json:"title" binding:"required"`
	Description  string  `json:"description"`
	Participants string  `json:"participants"`
	Expiry       float32 `json:"expiry"`
	ViewCount    int64   `json:"viewCount"`
	UserLimit    int     `json:"userLimit" binding:"required"`
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

type StdError struct {
	Message string `json:"message"` 
}

func PostRequestToDBModel(req Post) models.Post {
	return models.Post {
		UserID:       req.UserID,  
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

func CommentRequestToDBModel(req Comment) models.Comment {
	return models.Comment {
		UserID:   req.UserID,
		PostID:   req.PostID,
		Message:  req.Message,
		ParentId: req.ParentId,
		Votes:    req.Votes,
	}
}

func UserRequestToDBModel(req UserRegister) models.User {
	return models.User {
		Username: req.Username,
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Email: req.Email,
		Password: req.Password,
	}
}