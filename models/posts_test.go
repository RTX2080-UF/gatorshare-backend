package models

import (
	"gatorshare/middleware"
	"testing"

	"gorm.io/gorm"
)

type TestSuit struct {
	DB *gorm.DB
} 

var testobj = TestSuit{}

func TestMain (t *testing.T) {
	envsrc := middleware.LoadEnv(".env")
  	Init(envsrc)
  	testobj.DB = GetDB()
	// os.Exit(t.Run())
}

// func TestAddNewpost(t *testing.T) {
  	
// 	var post *Post;
		
// 	post.Title = "Test post 1"
// 	post.Description = "Test Message"
// 	post.UserLimit = 4
// 	post.Status = 2
	
// 	res, _ := AddNewpost(testobj.DB, post)
// 	// if res == {
// 	// t.Error("Test for Create Post has failed!")
//   	// }  
// 	print(res)  

// }
