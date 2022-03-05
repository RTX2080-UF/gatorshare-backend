package test

import (
	"fmt"
	"gatorshare/models"
	"math/big"
	"crypto/rand"
	"testing"
)

func TestCreateNewComment(t *testing.T) {
	rnum, _ := rand.Int(rand.Reader, big.NewInt(1000))

	user := &models.User{
		Username: "TestUser1" + fmt.Sprint(rnum),
		Firstname: "Test User",
		Email: "TestUser1" + fmt.Sprint(rnum) + "@gatorshare.com",
		Lastname: "1",
		Password: "Test",
	} 

	models.AddNewUser(testobj.DB, user)

	post := &models.Post{
		Title: "Test post 1",
		UserID: user.ID,
		Description: "Test Message",
		UserLimit: 4,
		Status: 2,
	}

	models.AddNewpost(testobj.DB, post)
	comment := &models.Comment {
		PostID: post.ID,
		UserID: user.ID,
		Message: "Test comment" + fmt.Sprint(rnum),
  	} 

  res, _ := models.AddNewcomment(testobj.DB, comment)
  if res == 0 {
	  t.Error("Unable to create post!")
	} else {
	  t.Log("Succesfully able to create post")	  
  }
 
  testobj.DB.Delete(user)
  testobj.DB.Delete(post)
  testobj.DB.Delete(comment)
  print(res)
}
