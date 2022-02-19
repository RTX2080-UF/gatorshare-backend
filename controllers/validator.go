package controllers

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
	UserID       int 		`json:"userId"`
	Title        string    	`json:"title" binding:"required"`
	Description  string    	`json:"description"`
	Participants string    	`json:"participants"`
	Expiry       float32   	`json:"expiry"`
	ViewCount    int64     	`json:"viewCount"`
	UserLimit    int       	`json:"userLimit" binding:"required"`
	Status       int       	`json:"status" binding:"required"`
	Categories   string
}

type Comment struct {
	UserId     uint      `json:"userId" binding:"required"`
	PostId     uint      `json:"postId" binding:"required"`
	Message    string    `json:"message" binding:"required"`
	ParentId   uint      `json:"parentId"`
	Votes      int       `json:"votes"`
}

type Tag struct {
	Name		string	 `json:"tagName" binding:"required"`
	Frequency	int		 `json:"count"`
}